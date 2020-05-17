// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"bytes"
	"github.com/koykov/fastconv"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestObjectInspector struct {
	inspector.BaseInspector
}

func (i3 *TestObjectInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i3.GetTo(src, &buf, path...)
	return buf, err
}

func (i3 *TestObjectInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "Id" {
			*buf = &x.Id
			return
		}
		if path[0] == "Name" {
			*buf = &x.Name
			return
		}
		if path[0] == "Status" {
			*buf = &x.Status
			return
		}
		if path[0] == "Cost" {
			*buf = &x.Cost
			return
		}
		if path[0] == "Permission" {
			x0 := x.Permission
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				var k int32
				t13, err13 := strconv.ParseInt(path[1], 0, 0)
				if err13 != nil {
					return err13
				}
				k = int32(t13)
				x1 := (*x0)[k]
				_ = x1
				*buf = &x1
				return
			}
			*buf = x0
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					if len(path) > 2 {
						if x1 == nil {
							return
						}
						if path[2] == "DateUnix" {
							*buf = &x1.DateUnix
							return
						}
						if path[2] == "Cost" {
							*buf = &x1.Cost
							return
						}
						if path[2] == "Comment" {
							*buf = &x1.Comment
							return
						}
					}
				}
			}
			*buf = x0
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					*buf = &x1
					return
				}
			}
			*buf = x0
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "MoneyIn" {
					*buf = &x0.MoneyIn
					return
				}
				if path[1] == "MoneyOut" {
					*buf = &x0.MoneyOut
					return
				}
				if path[1] == "Balance" {
					*buf = &x0.Balance
					return
				}
				if path[1] == "AllowBuy" {
					*buf = &x0.AllowBuy
					return
				}
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					if len(path) > 2 {
						var i int
						t14, err14 := strconv.ParseInt(path[2], 0, 0)
						if err14 != nil {
							return err14
						}
						i = int(t14)
						if len(x1) > i {
							x2 := &x1[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									*buf = &x2.DateUnix
									return
								}
								if path[3] == "Cost" {
									*buf = &x2.Cost
									return
								}
								if path[3] == "Comment" {
									*buf = &x2.Comment
									return
								}
							}
						}
					}
					*buf = x1
				}
			}
		}
	}
	return
}

func (i3 *TestObjectInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "Id" {
			var rightExact string
			rightExact = right

			switch cond {
			case inspector.OpEq:
				*result = x.Id == rightExact
			case inspector.OpNq:
				*result = x.Id != rightExact
			case inspector.OpGt:
				*result = x.Id > rightExact
			case inspector.OpGtq:
				*result = x.Id >= rightExact
			case inspector.OpLt:
				*result = x.Id < rightExact
			case inspector.OpLtq:
				*result = x.Id <= rightExact
			}
			return
		}
		if path[0] == "Name" {
			var rightExact []byte
			rightExact = fastconv.S2B(right)

			if cond == inspector.OpEq {
				*result = bytes.Equal(x.Name, rightExact)
			} else {
				*result = !bytes.Equal(x.Name, rightExact)
			}
			return
		}
		if path[0] == "Status" {
			var rightExact int32
			t17, err17 := strconv.ParseInt(right, 0, 0)
			if err17 != nil {
				return err17
			}
			rightExact = int32(t17)
			switch cond {
			case inspector.OpEq:
				*result = x.Status == rightExact
			case inspector.OpNq:
				*result = x.Status != rightExact
			case inspector.OpGt:
				*result = x.Status > rightExact
			case inspector.OpGtq:
				*result = x.Status >= rightExact
			case inspector.OpLt:
				*result = x.Status < rightExact
			case inspector.OpLtq:
				*result = x.Status <= rightExact
			}
			return
		}
		if path[0] == "Cost" {
			var rightExact float64
			t18, err18 := strconv.ParseFloat(right, 0)
			if err18 != nil {
				return err18
			}
			rightExact = float64(t18)
			switch cond {
			case inspector.OpEq:
				*result = x.Cost == rightExact
			case inspector.OpNq:
				*result = x.Cost != rightExact
			case inspector.OpGt:
				*result = x.Cost > rightExact
			case inspector.OpGtq:
				*result = x.Cost >= rightExact
			case inspector.OpLt:
				*result = x.Cost < rightExact
			case inspector.OpLtq:
				*result = x.Cost <= rightExact
			}
			return
		}
		if path[0] == "Permission" {
			x0 := x.Permission
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				var k int32
				t19, err19 := strconv.ParseInt(path[1], 0, 0)
				if err19 != nil {
					return err19
				}
				k = int32(t19)
				x1 := (*x0)[k]
				_ = x1
				var rightExact bool
				t20, err20 := strconv.ParseBool(right)
				if err20 != nil {
					return err20
				}
				rightExact = bool(t20)
				if cond == inspector.OpEq {
					*result = x1 == rightExact
				} else {
					*result = x1 != rightExact
				}
				return
			}
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					if len(path) > 2 {
						if x1 == nil {
							return
						}
						if path[2] == "DateUnix" {
							var rightExact int64
							t21, err21 := strconv.ParseInt(right, 0, 0)
							if err21 != nil {
								return err21
							}
							rightExact = int64(t21)
							switch cond {
							case inspector.OpEq:
								*result = x1.DateUnix == rightExact
							case inspector.OpNq:
								*result = x1.DateUnix != rightExact
							case inspector.OpGt:
								*result = x1.DateUnix > rightExact
							case inspector.OpGtq:
								*result = x1.DateUnix >= rightExact
							case inspector.OpLt:
								*result = x1.DateUnix < rightExact
							case inspector.OpLtq:
								*result = x1.DateUnix <= rightExact
							}
							return
						}
						if path[2] == "Cost" {
							var rightExact float64
							t22, err22 := strconv.ParseFloat(right, 0)
							if err22 != nil {
								return err22
							}
							rightExact = float64(t22)
							switch cond {
							case inspector.OpEq:
								*result = x1.Cost == rightExact
							case inspector.OpNq:
								*result = x1.Cost != rightExact
							case inspector.OpGt:
								*result = x1.Cost > rightExact
							case inspector.OpGtq:
								*result = x1.Cost >= rightExact
							case inspector.OpLt:
								*result = x1.Cost < rightExact
							case inspector.OpLtq:
								*result = x1.Cost <= rightExact
							}
							return
						}
						if path[2] == "Comment" {
							var rightExact []byte
							rightExact = fastconv.S2B(right)

							if cond == inspector.OpEq {
								*result = bytes.Equal(x1.Comment, rightExact)
							} else {
								*result = !bytes.Equal(x1.Comment, rightExact)
							}
							return
						}
					}
				}
			}
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					var rightExact int32
					t24, err24 := strconv.ParseInt(right, 0, 0)
					if err24 != nil {
						return err24
					}
					rightExact = int32(t24)
					switch cond {
					case inspector.OpEq:
						*result = x1 == rightExact
					case inspector.OpNq:
						*result = x1 != rightExact
					case inspector.OpGt:
						*result = x1 > rightExact
					case inspector.OpGtq:
						*result = x1 >= rightExact
					case inspector.OpLt:
						*result = x1 < rightExact
					case inspector.OpLtq:
						*result = x1 <= rightExact
					}
					return
				}
			}
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "MoneyIn" {
					var rightExact float64
					t25, err25 := strconv.ParseFloat(right, 0)
					if err25 != nil {
						return err25
					}
					rightExact = float64(t25)
					switch cond {
					case inspector.OpEq:
						*result = x0.MoneyIn == rightExact
					case inspector.OpNq:
						*result = x0.MoneyIn != rightExact
					case inspector.OpGt:
						*result = x0.MoneyIn > rightExact
					case inspector.OpGtq:
						*result = x0.MoneyIn >= rightExact
					case inspector.OpLt:
						*result = x0.MoneyIn < rightExact
					case inspector.OpLtq:
						*result = x0.MoneyIn <= rightExact
					}
					return
				}
				if path[1] == "MoneyOut" {
					var rightExact float64
					t26, err26 := strconv.ParseFloat(right, 0)
					if err26 != nil {
						return err26
					}
					rightExact = float64(t26)
					switch cond {
					case inspector.OpEq:
						*result = x0.MoneyOut == rightExact
					case inspector.OpNq:
						*result = x0.MoneyOut != rightExact
					case inspector.OpGt:
						*result = x0.MoneyOut > rightExact
					case inspector.OpGtq:
						*result = x0.MoneyOut >= rightExact
					case inspector.OpLt:
						*result = x0.MoneyOut < rightExact
					case inspector.OpLtq:
						*result = x0.MoneyOut <= rightExact
					}
					return
				}
				if path[1] == "Balance" {
					var rightExact float64
					t27, err27 := strconv.ParseFloat(right, 0)
					if err27 != nil {
						return err27
					}
					rightExact = float64(t27)
					switch cond {
					case inspector.OpEq:
						*result = x0.Balance == rightExact
					case inspector.OpNq:
						*result = x0.Balance != rightExact
					case inspector.OpGt:
						*result = x0.Balance > rightExact
					case inspector.OpGtq:
						*result = x0.Balance >= rightExact
					case inspector.OpLt:
						*result = x0.Balance < rightExact
					case inspector.OpLtq:
						*result = x0.Balance <= rightExact
					}
					return
				}
				if path[1] == "AllowBuy" {
					var rightExact bool
					t28, err28 := strconv.ParseBool(right)
					if err28 != nil {
						return err28
					}
					rightExact = bool(t28)
					if cond == inspector.OpEq {
						*result = x0.AllowBuy == rightExact
					} else {
						*result = x0.AllowBuy != rightExact
					}
					return
				}
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					if len(path) > 2 {
						var i int
						t29, err29 := strconv.ParseInt(path[2], 0, 0)
						if err29 != nil {
							return err29
						}
						i = int(t29)
						if len(x1) > i {
							x2 := &x1[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									var rightExact int64
									t30, err30 := strconv.ParseInt(right, 0, 0)
									if err30 != nil {
										return err30
									}
									rightExact = int64(t30)
									switch cond {
									case inspector.OpEq:
										*result = x2.DateUnix == rightExact
									case inspector.OpNq:
										*result = x2.DateUnix != rightExact
									case inspector.OpGt:
										*result = x2.DateUnix > rightExact
									case inspector.OpGtq:
										*result = x2.DateUnix >= rightExact
									case inspector.OpLt:
										*result = x2.DateUnix < rightExact
									case inspector.OpLtq:
										*result = x2.DateUnix <= rightExact
									}
									return
								}
								if path[3] == "Cost" {
									var rightExact float64
									t31, err31 := strconv.ParseFloat(right, 0)
									if err31 != nil {
										return err31
									}
									rightExact = float64(t31)
									switch cond {
									case inspector.OpEq:
										*result = x2.Cost == rightExact
									case inspector.OpNq:
										*result = x2.Cost != rightExact
									case inspector.OpGt:
										*result = x2.Cost > rightExact
									case inspector.OpGtq:
										*result = x2.Cost >= rightExact
									case inspector.OpLt:
										*result = x2.Cost < rightExact
									case inspector.OpLtq:
										*result = x2.Cost <= rightExact
									}
									return
								}
								if path[3] == "Comment" {
									var rightExact []byte
									rightExact = fastconv.S2B(right)

									if cond == inspector.OpEq {
										*result = bytes.Equal(x2.Comment, rightExact)
									} else {
										*result = !bytes.Equal(x2.Comment, rightExact)
									}
									return
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

func (i3 *TestObjectInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "Permission" {
			x0 := x.Permission
			_ = x0
			if x0 == nil {
				return
			}
			for k := range *x0 {
				if l.RequireKey() {
					*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal((*x0)[k], &inspector.StaticInspector{})
				ctl := l.Iterate()
				if ctl == inspector.LoopCtlBrk {
					break
				}
				if ctl == inspector.LoopCtlCnt {
					continue
				}
			}
			return
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			for k := range x0 {
				if l.RequireKey() {
					*buf = append((*buf)[:0], k...)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal((x0)[k], &TestHistoryInspector{})
				ctl := l.Iterate()
				if ctl == inspector.LoopCtlBrk {
					break
				}
				if ctl == inspector.LoopCtlCnt {
					continue
				}
			}
			return
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			for k := range x0 {
				if l.RequireKey() {
					*buf = append((*buf)[:0], k...)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal((x0)[k], &inspector.StaticInspector{})
				ctl := l.Iterate()
				if ctl == inspector.LoopCtlBrk {
					break
				}
				if ctl == inspector.LoopCtlCnt {
					continue
				}
			}
			return
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					for k := range x1 {
						if l.RequireKey() {
							*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
							l.SetKey(buf, &inspector.StaticInspector{})
						}
						l.SetVal(&(x1)[k], &TestHistoryInspector{})
						ctl := l.Iterate()
						if ctl == inspector.LoopCtlBrk {
							break
						}
						if ctl == inspector.LoopCtlCnt {
							continue
						}
					}
					return
				}
			}
		}
	}
	return
}

func (i3 *TestObjectInspector) Set(dst, value interface{}, path ...string) {
}
