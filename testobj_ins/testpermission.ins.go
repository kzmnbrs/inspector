// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"encoding/json"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestPermissionInspector struct {
	inspector.BaseInspector
}

func (i4 *TestPermissionInspector) TypeName() string {
	return "TestPermission"
}

func (i4 *TestPermissionInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i4.GetTo(src, &buf, path...)
	return buf, err
}

func (i4 *TestPermissionInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := src.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		var k int32
		t37, err37 := strconv.ParseInt(path[0], 0, 0)
		if err37 != nil {
			return err37
		}
		k = int32(t37)
		x0 := (*x)[k]
		_ = x0
		*buf = &x0
		return
	}
	return
}

func (i4 *TestPermissionInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := src.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var k int32
		t38, err38 := strconv.ParseInt(path[0], 0, 0)
		if err38 != nil {
			return err38
		}
		k = int32(t38)
		x0 := (*x)[k]
		_ = x0
		var rightExact bool
		t39, err39 := strconv.ParseBool(right)
		if err39 != nil {
			return err39
		}
		rightExact = bool(t39)
		if cond == inspector.OpEq {
			*result = x0 == rightExact
		} else {
			*result = x0 != rightExact
		}
		return
	}
	return
}

func (i4 *TestPermissionInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := src.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}

	for k := range *x {
		if l.RequireKey() {
			*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
			l.SetKey(buf, &inspector.StaticInspector{})
		}
		l.SetVal((*x)[k], &inspector.StaticInspector{})
		ctl := l.Iterate()
		if ctl == inspector.LoopCtlBrk {
			break
		}
		if ctl == inspector.LoopCtlCnt {
			continue
		}
	}
	return
	return
}

func (i4 *TestPermissionInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := dst.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := dst.(testobj.TestPermission); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		var k int32
		t40, err40 := strconv.ParseInt(path[0], 0, 0)
		if err40 != nil {
			return err40
		}
		k = int32(t40)
		x0 := (*x)[k]
		_ = x0
		inspector.AssignBuf(&x0, value, buf)
		(*x)[k] = x0
		return nil
	}
	return nil
}

func (i4 *TestPermissionInspector) Set(dst, value interface{}, path ...string) error {
	return i4.SetWB(dst, value, nil, path...)
}

func (i4 *TestPermissionInspector) DeepEqual(l, r interface{}) bool {
	return i4.DeepEqualWithOptions(l, r, nil)
}

func (i4 *TestPermissionInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestPermission
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestPermission); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestPermission); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestPermission); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestPermission); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestPermission); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestPermission); ok {
		rx, req = &rp, true
	}
	if !leq || !req {
		return false
	}
	if lx == nil && rx == nil {
		return true
	}
	if (lx == nil && rx != nil) || (lx != nil && rx == nil) {
		return false
	}

	if len(*lx) != len(*rx) {
		return false
	}
	for k := range *lx {
		lx1 := (*lx)[k]
		rx1, ok1 := (*rx)[k]
		_, _, _ = lx1, rx1, ok1
		if !ok1 {
			return false
		}
		if lx1 != rx1 && inspector.DEQMustCheck("", opts) {
			return false
		}
	}
	return true
}

func (i4 *TestPermissionInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestPermission
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i4 *TestPermissionInspector) Copy(x interface{}) (interface{}, error) {
	var cpy testobj.TestPermission
	switch x.(type) {
	case testobj.TestPermission:
		cpy = x.(testobj.TestPermission)
	case *testobj.TestPermission:
		cpy = *x.(*testobj.TestPermission)
	case **testobj.TestPermission:
		cpy = **x.(**testobj.TestPermission)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	return cpy, nil
}
