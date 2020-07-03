package inspector

import (
	"errors"
	"go/format"
	"go/types"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/tools/go/loader"
)

// Internal node type.
type typ int

// Node compile mode.
type mode int

const (
	// Possible node's types.
	typeStruct typ = iota
	typeMap
	typeSlice
	typeBasic

	// Possible compile modes.
	modeGet mode = iota
	modeSet
	modeCmp
	modeLoop
)

// Extended writer interface.
type ByteStringWriter interface {
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteString(s string) (n int, err error)
	Bytes() []byte
	String() string
	Reset()
}

// Logger interface.
type Logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
}

// The compiler.
type Compiler struct {
	// Path to the package relative to $GOPATH/src.
	pkg string
	// THe same as pkg plus dot at the end, needs for internal logic.
	pkgDot string
	// Only package name.
	pkgName string
	// Destination dir relative to $GOPATH/src.
	dst string
	// Absolute path to the destination dir.
	dstAbs string
	// List of blacklisted types (key is a type name, value is ignored).
	bl map[string]bool
	// Cache of processed types, needs to avoid duplications.
	uniq map[string]bool
	// Tree of parsed types.
	nodes []*node
	// List of imports that should be in final output.
	imp []string
	// Internal types counter.
	cntr int
	// Logger object
	l Logger
	// Writer object
	wr ByteStringWriter

	err error
}

// Node type.
// It's an internal representation of Go type.
type node struct {
	// Type flag, see constants above.
	typ typ
	// Type name as string.
	typn string
	// Underlying type name.
	typu string
	// If node has a parent, eg struct or map, this field contains the name or key.
	name string
	// Package name, requires for import suffixes.
	pkg string
	// Package path, requires for imports.
	pkgi string
	// Is node passed as pointer in parent or not.
	ptr bool
	// List of child nodes.
	chld []*node
	// If is a map, that node contains information about key's type.
	mapk *node
	// If is a map, that node contains information about value's type.
	mapv *node
	// If is a slice, that node contains information about value's type.
	slct *node
}

var (
	// Check map.
	reMap = regexp.MustCompile(`map\[[^]]+].*`)
	// Check slice.
	reSlc = regexp.MustCompile(`\[].*`)
	// Check if first symbol is in upper case.
	reUp = regexp.MustCompile(`^[A-Z]+.*`)
	// Check vendor directory.
	reVnd = regexp.MustCompile(`.*/vendor/(.*)`)

	// Std errors.
	ErrNoGOPATH     = errors.New("no GOPATH variable found")
	ErrDstNotExists = errors.New("destination directory doesn't exists")
)

// Make new instance of the compiler.
func NewCompiler(pkg, dst string, bl map[string]bool, w ByteStringWriter, l Logger) *Compiler {
	c := Compiler{
		pkg:    pkg,
		pkgDot: pkg + ".",
		dst:    dst,
		bl:     bl,
		uniq:   make(map[string]bool),
		wr:     w,
		l:      l,
		imp:    make([]string, 0),
	}
	return &c
}

func (c *Compiler) String() string {
	return ""
}

// Main compile method.
func (c *Compiler) Compile() error {
	if c.l != nil {
		c.l.Print("Parse package " + c.pkg)
	}
	// Try import the package.
	var conf loader.Config
	conf.Import(c.pkg)
	prog, err := conf.Load()
	if err != nil {
		return err
	}

	// Parse the package to nodes list.
	pkg := prog.Package(c.pkg)
	c.pkgName = pkg.Pkg.Name()
	err = c.parsePkg(pkg)
	if err != nil {
		return err
	}

	// Prepare destination.
	if c.l != nil {
		c.l.Print("Prepare destination dir " + c.dst)
	}
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		return ErrNoGOPATH
	}
	ps := string(os.PathSeparator)
	c.dstAbs = os.Getenv("GOPATH") + ps + "src" + ps + c.dst

	dstExists := true
	if _, err := os.Stat(c.dstAbs); os.IsNotExist(err) {
		dstExists = false
	}
	if dstExists {
		if err := syscall.Access(c.dstAbs, syscall.O_RDWR); err != nil {
			return err
		}
		if err := os.RemoveAll(c.dstAbs); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(c.dstAbs, 0755); err != nil {
		return ErrDstNotExists
	}

	// Write init file of destination package.
	file := "init.go"
	if c.l != nil {
		c.l.Print("Compiling init to " + file)
	}
	err = c.writeInit()
	if err != nil {
		return err
	}
	err = c.writeFile(c.dstAbs + ps + "init.go")
	if err != nil {
		return err
	}

	// Walk over nodes and compile each of them to separate file.
	for _, node := range c.nodes {
		file = strings.ToLower(node.name) + ".ins.go"
		if c.l != nil {
			c.l.Print("Compiling ", node.name, "Inspector to "+file)
		}
		c.wr.Reset()
		err = c.writeType(node)
		if err != nil {
			return err
		}
		err = c.writeFile(c.dstAbs + ps + file)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get the total number of compiled types.
func (c *Compiler) GetTotal() int {
	return c.cntr
}

// Parse package.
func (c *Compiler) parsePkg(pkg *loader.PackageInfo) error {
	for _, scope := range pkg.Info.Scopes {
		// Only scopes without parents available.
		if parent := scope.Parent(); parent != nil {
			for _, name := range parent.Names() {
				// Get the object representation of scope.
				o := parent.Lookup(name)
				if !c.isExported(o) {
					continue
				}
				// Get type of object and parse it.
				t := o.Type()
				node, err := c.parseType(t)
				if err != nil {
					return err
				}
				if node == nil || node.typ == typeBasic || reMap.MatchString(node.typn) || reSlc.MatchString(node.typn) {
					continue
				}
				if node.typ == typeStruct {
					// Make type name of struct node actual.
					node.typn = o.Name()
				}
				node.pkg = o.Pkg().Name()
				node.name = o.Name()
				// Check and skip type if it is already parsed or blacklisted.
				if _, ok := c.uniq[node.name]; ok {
					continue
				}
				if len(c.bl) > 0 {
					if _, ok := c.bl[node.name]; ok {
						continue
					}
				}
				c.uniq[node.name] = true

				c.nodes = append(c.nodes, node)
			}
		}
	}

	return nil
}

// Main parse method.
func (c *Compiler) parseType(t types.Type) (*node, error) {
	// Prepare and fill up default node.
	node := &node{
		typ:  typeBasic,
		typn: strings.Replace(t.String(), c.pkgDot, "", 1),
		ptr:  false,
	}
	if n, ok := t.(*types.Named); ok {
		node.typn = n.Obj().Name()
		if n.Obj().Pkg() == nil {
			return nil, nil
		}
		node.pkg = n.Obj().Pkg().Name()
		node.pkgi = c.clearPkg(n.Obj().Pkg().Path())
	}

	// Get the underlying type.
	u := t.Underlying()
	// Common skips considering by underlying type.
	if _, ok := u.(*types.Interface); ok {
		return nil, nil
	}
	if _, ok := u.(*types.Signature); ok {
		return nil, nil
	}

	if p, ok := u.(*types.Pointer); ok {
		// Dereference pointer type and go inside to parse.
		e := p.Elem()
		u = e.Underlying()
		unode, err := c.parseType(u)
		if err != nil {
			return node, err
		}
		if unode == nil {
			return nil, nil
		}
		unode.ptr = true
		if n, ok := e.(*types.Named); ok {
			unode.typn = n.Obj().Name()
			unode.pkg = n.Obj().Pkg().Name()
			unode.pkgi = c.clearPkg(n.Obj().Pkg().Path())
		}
		return unode, err
	}

	if s, ok := u.(*types.Struct); ok {
		// Walk over fields in struct and parse each of them as separate node.
		node.typ = typeStruct
		for i := 0; i < s.NumFields(); i++ {
			f := s.Field(i)
			if !f.Exported() {
				continue
			}
			ch, err := c.parseType(f.Type())
			if err != nil {
				return node, err
			}
			if ch == nil {
				continue
			}
			ch.name = f.Name()
			if ch.ptr {
				ch.typn = strings.Replace(f.Type().String(), c.pkgDot, "", 1)
			}
			node.chld = append(node.chld, ch)
		}
		return node, nil
	}

	if m, ok := u.(*types.Map); ok {
		// Just parse key and value types of map.
		var err error
		node.typ = typeMap
		node.mapk, err = c.parseType(m.Key())
		node.mapv, err = c.parseType(m.Elem())
		return node, err
	}

	if s, ok := u.(*types.Slice); ok {
		// Just parse value type of slice.
		var err error
		node.typ = typeSlice
		node.slct, err = c.parseType(s.Elem())
		return node, err
	}

	if b, ok := u.(*types.Basic); ok {
		// Fill up the underlying type of basic node.
		node.typu = b.Name()
	}

	return node, nil
}

// Format compiled inspector code using format package and write it to the given file.
func (c *Compiler) writeFile(filename string) error {
	source := c.wr.Bytes()
	fmtSource, err := format.Source(source)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, fmtSource, 0644)
}

// Build and write init() func of the generated package.
func (c *Compiler) writeInit() error {
	c.wl("// Code generated by inspc. DO NOT EDIT.")
	c.wdl("// source: ", c.pkg)

	c.wdl("package ", c.pkgName, "_ins")
	c.wdl("import (\n\"github.com/koykov/inspector\"\n)")

	c.wl("func init() {")
	for _, node := range c.nodes {
		c.wl(`inspector.RegisterInspector("`, node.name, `", &`, node.name, `Inspector{})`)
	}
	c.wdl("}")

	return nil
}

// Write given type including headers and imports.
func (c *Compiler) writeType(node *node) error {
	c.wl("// Code generated by inspc. DO NOT EDIT.")
	c.wdl("// source: ", c.pkg)

	c.wdl("package ", c.pkgName, "_ins")

	c.imp = append(c.imp[:0], `"github.com/koykov/inspector"`, `"`+c.pkg+`"`)
	c.wdl("import (\n!{import}\n)")

	err := c.writeRootNode(node, c.cntr)
	if err != nil {
		return err
	}
	c.cntr++
	c.r("!{import}", strings.Join(c.imp, "\n"))

	return nil
}

// Register imports for output file.
func (c *Compiler) regImport(imports []string) {
	for _, i := range imports {
		found := false
		for _, j := range c.imp {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			c.imp = append(c.imp, i)
		}
	}
}

// Write methods of the inspector required by Inspector interface type.
func (c *Compiler) writeRootNode(node *node, idx int) (err error) {
	inst := node.name + "Inspector"
	recv := "i" + strconv.Itoa(idx)
	pname := c.pkgName + "." + node.typn

	c.wdl("type ", inst, " struct {\ninspector.BaseInspector\n}")

	// Common checks and type cast.
	funcHeader := `if len(path) == 0 { return }
if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`
	// Custom header for Set() method.
	funcHeaderSet := `if len(path) == 0 { return nil }
if dst == nil { return nil }
var x *` + pname + `
_ = x
if p, ok := dst.(*` + pname + `); ok { x = p } else if v, ok := dst.(` + pname + `); ok { x = &v } else { return nil }`
	// Custom header for GetTo() method.
	funcHeaderGetTo := `if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }
if len(path) == 0 { *buf = &(*x)
return}`
	// Header for Loop() method.
	funcHeaderLoop := ""
	if node.typ != typeSlice {
		funcHeaderLoop += "if len(path) == 0 { return }\n"
	}
	funcHeaderLoop += `if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`

	// Getter methods.
	c.wl("func (", recv, " *", inst, ") Get(src interface{}, path ...string) (interface{}, error) {")
	c.wl("var buf interface{}\nerr := " + recv + ".GetTo(src, &buf, path...)\nreturn buf, err")
	c.wdl("}")

	c.wl("func (", recv, " *", inst, ") GetTo(src interface{}, buf *interface{}, path ...string) (err error) {")
	c.wdl(funcHeaderGetTo)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeGet)
	if err != nil {
		return err
	}
	c.wl("*buf = &(*x)")
	c.wdl("return }")

	// Compare method.
	c.wl("func (", recv, " *", inst, ") Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {")
	c.wdl(funcHeader)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeCmp)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Loop method.
	c.wl("func (", recv, " *", inst, ") Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {")
	c.wdl(funcHeaderLoop)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeLoop)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Setter method.
	c.wl("func (", recv, " *", inst, ") Set(dst, value interface{}, path ...string) error {")
	c.wdl(funcHeaderSet)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeSet)
	if err != nil {
		return err
	}
	c.wdl("return nil }")

	return c.err
}

// Write body of the methods according given mode.
func (c *Compiler) writeNode(node, parent *node, recv, v, vsrc string, depth int, mode mode) error {
	depths := strconv.Itoa(depth)

	// Flag to check length of the path array.
	requireLenCheck := node.typ != typeBasic && !(mode == modeLoop && (node.typ == typeMap || (node.typ == typeSlice && node.typn != "[]byte")))

	if requireLenCheck {
		c.wl("if len(path) > ", depths, " {")
	}
	if node.ptr {
		// Value may be nil on pointer types.
		c.wl("if ", v, " == nil { ", c.fmtR(mode, "nil"), " }")
	}

	switch node.typ {
	case typeStruct:
		// Walk over fields.
		for _, ch := range node.chld {
			isBasic := ch.typ == typeBasic || (ch.typ == typeSlice && ch.typn == "[]byte")
			if isBasic && mode == modeLoop {
				continue
			}
			c.wl("if path[", depths, "] == ", `"`, ch.name, `" {`)
			if isBasic {
				switch mode {
				case modeGet:
					c.wl("*buf = &", v, ".", ch.name)
					c.wl("return")
				case modeCmp:
					c.writeCmp(ch, v+"."+ch.name)
					c.wl("return")
				case modeSet:
					pfx := ""
					if !c.isBuiltin(ch.typn) {
						pfx = ch.pkg + "."
					}
					c.wl("if exact, ok := value.(*", pfx, ch.typn, "); ok {", v, ".", ch.name, " = *exact}")
					c.wl("if exact, ok := value.(", pfx, ch.typn, "); ok {", v, ".", ch.name, " = exact}")
					c.wl("return nil")
				}
			} else {
				nv := "x" + strconv.Itoa(depth)
				vsrc := v + "." + ch.name
				pfx := ""
				if !ch.ptr && ch.typ != typeMap && ch.typ != typeSlice {
					pfx = "&"
				}
				c.wl(nv, " := ", pfx, vsrc)
				c.wl("_ = ", nv)
				if mode == modeCmp && ch.ptr {
					c.writeCmp(ch, nv)
				}
				err := c.writeNode(ch, node, recv, nv, vsrc, depth+1, mode)
				if err != nil {
					return err
				}
			}
			c.wl("}")
		}
	case typeMap:
		// todo find a way to get element from the map without allocation (see runtime.mapaccess1 and runtime.mapaccess2).
		origPtr := node.ptr
		if depth == 0 {
			node.ptr = true
		}
		switch mode {
		case modeLoop:
			// Loop magic.
			c.wl("for k := range ", c.fmtV(node, v), " {")
			c.wl("if l.RequireKey() {")
			switch node.mapk.typn {
			case "string", "[]byte":
				c.wl("*buf = append((*buf)[:0], k...)")
			case "bool":
				c.wl(`if k { *buf = append((*buf)[:0], "true"...) } else { *buf = append(*buf[:0], "false"...) }`)
			case "int", "int8", "int16", "int32", "int64":
				c.wl("*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)")
			case "uint", "uint8", "uint16", "uint32", "uint64":
				c.wl("*buf = strconv.AppendUint((*buf)[:0], uint64(k), 10)")
			case "float32", "float64":
				c.wl("*buf = strconv.AppendFloat((*buf)[:0], float6464(k), 'f', -1, 64)")
			default:
				c.regImport([]string{`"github.com/koykov/any2bytes"`})
				c.wl("*buf, err = any2bytes.AnyToBytes(*buf[:0], k)")
				c.wl("if err != nil { return }")
			}
			c.wl("l.SetKey(buf, &inspector.StaticInspector{})")
			c.wl("}")
			insName := "inspector.StaticInspector"
			if node.mapv.typ == typeStruct {
				insName = ""
				if node.mapv.pkg != c.pkgName {
					insName = node.mapv.pkg + "_ins."
					c.regImport([]string{`"` + node.mapv.pkgi + `_ins"`})
				}
				insName += node.mapv.typn + "Inspector"
			}
			c.wl("l.SetVal(", c.fmtV(node, v), "[k], &", insName, "{})")
			c.wl("ctl := l.Iterate()")
			c.wl("if ctl == inspector.LoopCtlBrk { break }")
			c.wl("if ctl == inspector.LoopCtlCnt { continue }")
			c.wl("}")
			c.wl("return")
		default:
			nv := "x" + strconv.Itoa(depth)
			if node.mapk.typn == "string" {
				// Key is string, simple case.
				c.wl("if ", nv, ", ok := ", c.fmtV(node, v), "[path[", depths, "]]; ok {")
				c.wl("_ = ", nv)
				err := c.writeNode(node.mapv, node, recv, nv, "", depth+1, mode)
				if err != nil {
					return err
				}
				if mode == modeSet {
					c.wl(c.fmtV(node, v), "[path[", depths, "]] = ", nv)
					c.wl("return nil")
				}
				c.wl("}")
			} else {
				// Convert path value to the key type and try to find it in the map.
				c.wl("var k ", node.mapk.typn)
				snippet, imports, err := StrConvSnippet("path["+depths+"]", node.mapk.typn, node.mapk.typu, "k")
				c.regImport(imports)
				if err != nil {
					return err
				}
				c.wl(snippet)
				c.wl(nv, " := ", c.fmtV(node, v), "[k]")
				c.wl("_ = ", nv)
				err = c.writeNode(node.mapv, node, recv, nv, "", depth+1, mode)
				if mode == modeSet {
					c.wl(c.fmtV(node, v), "[k] = ", nv)
					c.wl("return nil")
				}
				if err != nil {
					return err
				}
			}
		}
		node.ptr = origPtr
	case typeSlice:
		// Quick check of []byte type.
		if node.typn == "[]byte" {
			switch mode {
			case modeGet:
				c.wl("*buf = &", v)
				c.wl("return")
			case modeCmp:
				c.writeCmp(node, v)
				c.wl("return")
			case modeSet:
				c.wl("if exact, ok := value.(*", node.typn, "); ok {", v, " = *exact}")
				c.wl("if exact, ok := value.(", node.typn, "); ok {", v, " = exact}")
				c.wl("return nil")
			}
		}
		switch mode {
		case modeLoop:
			// Loop magic.
			c.wl("for k := range ", c.fmtVd(node, v, depth), " {")
			c.wl("if l.RequireKey() {")
			c.wl("*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)")
			c.wl("l.SetKey(buf, &inspector.StaticInspector{})")
			c.wl("}")
			insName := "inspector.StaticInspector"
			if node.slct.typ == typeStruct {
				insName = ""
				if node.slct.pkg != c.pkgName {
					insName = node.slct.pkg + "_ins."
					c.regImport([]string{`"` + node.slct.pkgi + `_ins"`})
				}
				insName += node.slct.typn + "Inspector"
			}
			c.wl("l.SetVal(&", c.fmtVd(node, v, depth), "[k], &", insName, "{})")
			c.wl("ctl := l.Iterate()")
			c.wl("if ctl == inspector.LoopCtlBrk { break }")
			c.wl("if ctl == inspector.LoopCtlCnt { continue }")
			c.wl("}")
			c.wl("return")
		default:
			// Convert path value to the int index and try to find value in the slice using it.
			nv := "x" + strconv.Itoa(depth)
			c.wl("var i int")
			snippet, imports, err := StrConvSnippet("path["+depths+"]", "int", "", "i")
			c.regImport(imports)
			if err != nil {
				return err
			}
			c.wl(snippet)
			c.wl("if len(", c.fmtVnb(node, v, depth), ") > i {")
			if node.slct.ptr || c.isBuiltin(node.slct.typn) {
				c.wl(nv, " := ", c.fmtVd(node, v, depth), "[i]")
			} else {
				c.wl(nv, " := &", c.fmtVd(node, v, depth), "[i]")
			}
			c.wl("_ = ", nv)
			err = c.writeNode(node.slct, node, recv, nv, "", depth+1, mode)
			if err != nil {
				return err
			}
			if mode == modeSet {
				pfx := ""
				if !node.slct.ptr && !c.isBuiltin(node.slct.typn) {
					pfx = "*"
				}
				c.wl(c.fmtVd(node, v, depth), "[i] = ", pfx, nv)
				c.wl("return nil")
			}
			c.wl("}")
		}
	case typeBasic:
		switch mode {
		case modeGet:
			c.wl("*buf = &", v)
			c.wl("return")
		case modeCmp:
			c.writeCmp(node, v)
			c.wl("return")
		case modeSet:
			pfx := ""
			if !c.isBuiltin(node.typn) {
				pfx = node.pkg + "."
			}
			c.wl("if exact, ok := value.(*", pfx, node.typn, "); ok {", v, " = *exact}")
			c.wl("if exact, ok := value.(", pfx, node.typn, "); ok {", v, " = exact}")
			if parent.typ != typeMap {
				c.wl("return nil")
			}
		}
	}
	if requireLenCheck {
		c.wl("}")
	}
	// Special case to take value by pointer to avoid allocation.
	if (node.typ == typeStruct || node.typ == typeMap || node.typ == typeSlice) && mode == modeGet && depth > 1 {
		pfx := ""
		if parent != nil && parent.typ != typeSlice {
			pfx = "&"
		}
		if len(vsrc) > 0 {
			c.wl("*buf = ", pfx, vsrc)
		} else {
			c.wl("*buf = ", pfx, v)
		}
	}

	return c.err
}

// Write comparison code of the node.
func (c *Compiler) writeCmp(left *node, leftVar string) {
	// Node is pointer, check nil case.
	if left.ptr {
		c.wl("if right==inspector.Nil {")
		c.wl("if cond == inspector.OpEq {")
		c.wl("*result = ", leftVar, " == nil")
		c.wl("} else {")
		c.wl("*result = ", leftVar, " != nil")
		c.wl("}")
		c.wl("return\n}")
		return
	}

	// Get type name as a string.
	bi := c.isBuiltin(left.typn)
	pname := left.typn
	if !bi {
		pname = c.pkgName + "." + pname
	}
	// Get the exact value of the right variable.
	c.wl("var rightExact ", pname)
	snippet, imports, err := StrConvSnippet("right", left.typn, left.typu, "rightExact")
	c.regImport(imports)
	if err != nil {
		c.err = err
		return
	}
	if !bi {
		snippet = strings.Replace(snippet, left.typu, pname, -1)
	}
	c.wl(snippet)

	switch left.typn {
	// []byte and bool types allows only equal and non-equal comparison operations.
	case "[]byte":
		c.wl("if cond == inspector.OpEq {")
		c.wl("*result = bytes.Equal(", leftVar, ", rightExact)")
		c.wl("} else {")
		c.wl("*result = !bytes.Equal(", leftVar, ", rightExact)")
		c.wl("}")
	case "bool":
		c.wl("if cond == inspector.OpEq {")
		c.wl("*result = ", leftVar, " == rightExact")
		c.wl("} else {")
		c.wl("*result = ", leftVar, " != rightExact")
		c.wl("}")
	default:
		// All other cases allows all type sof the comparison.
		c.wl("switch cond {")
		c.wl("case inspector.OpEq:")
		c.wl("*result = ", leftVar, " == rightExact")
		c.wl("case inspector.OpNq:")
		c.wl("*result = ", leftVar, " != rightExact")
		c.wl("case inspector.OpGt:")
		c.wl("*result = ", leftVar, " > rightExact")
		c.wl("case inspector.OpGtq:")
		c.wl("*result = ", leftVar, " >= rightExact")
		c.wl("case inspector.OpLt:")
		c.wl("*result = ", leftVar, " < rightExact")
		c.wl("case inspector.OpLtq:")
		c.wl("*result = ", leftVar, " <= rightExact")
		c.wl("}")
	}
}

// Check if given type is a builtin type.
func (c *Compiler) isBuiltin(typ string) bool {
	switch typ {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"string", "[]byte", "bool":
		return true
	default:
		return false
	}
}

// Check if given object has an exportable type.
func (c *Compiler) isExported(o types.Object) bool {
	if !o.Exported() {
		return false
	}
	if _, ok := o.(*types.Var); ok {
		return false
	}
	t := o.Type()
	if n, ok := t.(*types.Named); ok {
		if !reUp.MatchString(n.Obj().Name()) {
			return false
		}
	}
	u := t.Underlying()
	if p, ok := u.(*types.Pointer); ok {
		_ = p
		return true
	}
	return true
}

// Remove vendor path from the package.
func (c *Compiler) clearPkg(pkgi string) string {
	if m := reVnd.FindStringSubmatch(pkgi); m != nil {
		return m[1]
	}
	return pkgi
}

// Bounds variable according node options.
func (c *Compiler) fmtV(node *node, v string) string {
	if node.ptr {
		return "(*" + v + ")"
	}
	return "(" + v + ")"
}

// Bounds variable according node options and depth.
func (c *Compiler) fmtVd(node *node, v string, depth int) string {
	if node.ptr || depth == 0 {
		return "(*" + v + ")"
	}
	return "(" + v + ")"
}

// No brackets version of fmtV().
func (c *Compiler) fmtVnb(node *node, v string, depth int) string {
	if node.ptr || (node.typ == typeSlice && depth == 0) {
		return "*" + v
	}
	return v
}

func (c *Compiler) fmtR(mode mode, err string) string {
	if mode == modeSet {
		return "return " + err
	} else {
		return "return"
	}
}

// Shorthand write method.
func (c *Compiler) w(s ...string) {
	_, err := c.wr.WriteString(strings.Join(s, ""))
	if err != nil && c.err == nil {
		c.err = err
	}
}

// Shorthand write new line helper.
func (c *Compiler) wl(s ...string) {
	s = append(s, "\n")
	c.w(s...)
}

// Shorthand write new double lines helper.
func (c *Compiler) wdl(s ...string) {
	s = append(s, "\n\n")
	c.w(s...)
}

// Shorthand replace helper.
func (c *Compiler) r(old, new string) {
	s := c.wr.String()
	s = strings.Replace(s, old, new, -1)
	c.wr.Reset()
	_, _ = c.wr.WriteString(s)
}
