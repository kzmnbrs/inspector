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
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
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
		if path[0] == "Ustate" {
			*buf = &x.Ustate
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
				t14, err14 := strconv.ParseInt(path[1], 0, 0)
				if err14 != nil {
					return err14
				}
				k = int32(t14)
				x1 := (*x0)[k]
				_ = x1
				*buf = &x1
				return
			}
			*buf = &x.Permission
			return
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
					*buf = &x1
				}
			}
			*buf = &x.HistoryTree
			return
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
			*buf = &x.Flags
			return
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
						t15, err15 := strconv.ParseInt(path[2], 0, 0)
						if err15 != nil {
							return err15
						}
						i = int(t15)
						if len(x1) > i {
							x2 := &(x1)[i]
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
							*buf = x2
						}
					}
					*buf = &x0.History
					return
				}
			}
			*buf = &x.Finance
			return
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
	if p, ok := src.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestObject); ok {
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
			t18, err18 := strconv.ParseInt(right, 0, 0)
			if err18 != nil {
				return err18
			}
			rightExact = int32(t18)
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
		if path[0] == "Ustate" {
			var rightExact uint64
			t19, err19 := strconv.ParseUint(right, 0, 0)
			if err19 != nil {
				return err19
			}
			rightExact = uint64(t19)
			switch cond {
			case inspector.OpEq:
				*result = x.Ustate == rightExact
			case inspector.OpNq:
				*result = x.Ustate != rightExact
			case inspector.OpGt:
				*result = x.Ustate > rightExact
			case inspector.OpGtq:
				*result = x.Ustate >= rightExact
			case inspector.OpLt:
				*result = x.Ustate < rightExact
			case inspector.OpLtq:
				*result = x.Ustate <= rightExact
			}
			return
		}
		if path[0] == "Cost" {
			var rightExact float64
			t20, err20 := strconv.ParseFloat(right, 0)
			if err20 != nil {
				return err20
			}
			rightExact = float64(t20)
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
			if right == inspector.Nil {
				if cond == inspector.OpEq {
					*result = x0 == nil
				} else {
					*result = x0 != nil
				}
				return
			}
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				var k int32
				t21, err21 := strconv.ParseInt(path[1], 0, 0)
				if err21 != nil {
					return err21
				}
				k = int32(t21)
				x1 := (*x0)[k]
				_ = x1
				var rightExact bool
				t22, err22 := strconv.ParseBool(right)
				if err22 != nil {
					return err22
				}
				rightExact = bool(t22)
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
							t23, err23 := strconv.ParseInt(right, 0, 0)
							if err23 != nil {
								return err23
							}
							rightExact = int64(t23)
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
							t24, err24 := strconv.ParseFloat(right, 0)
							if err24 != nil {
								return err24
							}
							rightExact = float64(t24)
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
					t26, err26 := strconv.ParseInt(right, 0, 0)
					if err26 != nil {
						return err26
					}
					rightExact = int32(t26)
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
			if right == inspector.Nil {
				if cond == inspector.OpEq {
					*result = x0 == nil
				} else {
					*result = x0 != nil
				}
				return
			}
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "MoneyIn" {
					var rightExact float64
					t27, err27 := strconv.ParseFloat(right, 0)
					if err27 != nil {
						return err27
					}
					rightExact = float64(t27)
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
					t28, err28 := strconv.ParseFloat(right, 0)
					if err28 != nil {
						return err28
					}
					rightExact = float64(t28)
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
					t29, err29 := strconv.ParseFloat(right, 0)
					if err29 != nil {
						return err29
					}
					rightExact = float64(t29)
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
					t30, err30 := strconv.ParseBool(right)
					if err30 != nil {
						return err30
					}
					rightExact = bool(t30)
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
						t31, err31 := strconv.ParseInt(path[2], 0, 0)
						if err31 != nil {
							return err31
						}
						i = int(t31)
						if len(x1) > i {
							x2 := &(x1)[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									var rightExact int64
									t32, err32 := strconv.ParseInt(right, 0, 0)
									if err32 != nil {
										return err32
									}
									rightExact = int64(t32)
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
									t33, err33 := strconv.ParseFloat(right, 0)
									if err33 != nil {
										return err33
									}
									rightExact = float64(t33)
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
	if p, ok := src.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestObject); ok {
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

func (i3 *TestObjectInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := dst.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := dst.(testobj.TestObject); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "Id" {
			inspector.AssignBuf(&x.Id, value, buf)
			return nil
		}
		if path[0] == "Name" {
			inspector.AssignBuf(&x.Name, value, buf)
			return nil
		}
		if path[0] == "Status" {
			inspector.AssignBuf(&x.Status, value, buf)
			return nil
		}
		if path[0] == "Ustate" {
			inspector.AssignBuf(&x.Ustate, value, buf)
			return nil
		}
		if path[0] == "Cost" {
			inspector.AssignBuf(&x.Cost, value, buf)
			return nil
		}
		if path[0] == "Permission" {
			x0 := x.Permission
			if uvalue, ok := value.(*testobj.TestPermission); ok {
				x0 = uvalue
			}
			if x0 == nil {
				z := make(testobj.TestPermission)
				x0 = &z
				x.Permission = x0
			}
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return nil
				}
				var k int32
				t35, err35 := strconv.ParseInt(path[1], 0, 0)
				if err35 != nil {
					return err35
				}
				k = int32(t35)
				x1 := (*x0)[k]
				_ = x1
				inspector.AssignBuf(&x1, value, buf)
				(*x0)[k] = x1
				return nil
			}
			x.Permission = x0
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			if uvalue, ok := value.(*map[string]*testobj.TestHistory); ok {
				x0 = *uvalue
			}
			if x0 == nil {
				z := make(map[string]*testobj.TestHistory)
				x0 = z
				x.HistoryTree = x0
			}
			_ = x0
			if len(path) > 1 {
				x1 := (x0)[path[1]]
				_ = x1
				if len(path) > 2 {
					if x1 == nil {
						return nil
					}
					if path[2] == "DateUnix" {
						inspector.AssignBuf(&x1.DateUnix, value, buf)
						return nil
					}
					if path[2] == "Cost" {
						inspector.AssignBuf(&x1.Cost, value, buf)
						return nil
					}
					if path[2] == "Comment" {
						inspector.AssignBuf(&x1.Comment, value, buf)
						return nil
					}
				}
				(x0)[path[1]] = x1
				return nil
			}
			x.HistoryTree = x0
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			if uvalue, ok := value.(*testobj.TestFlag); ok {
				x0 = *uvalue
			}
			if x0 == nil {
				z := make(testobj.TestFlag)
				x0 = z
				x.Flags = x0
			}
			_ = x0
			if len(path) > 1 {
				x1 := (x0)[path[1]]
				_ = x1
				inspector.AssignBuf(&x1, value, buf)
				(x0)[path[1]] = x1
				return nil
			}
			x.Flags = x0
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			if uvalue, ok := value.(*testobj.TestFinance); ok {
				x0 = uvalue
			}
			if x0 == nil {
				x0 = &testobj.TestFinance{}
				x.Finance = x0
			}
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return nil
				}
				if path[1] == "MoneyIn" {
					inspector.AssignBuf(&x0.MoneyIn, value, buf)
					return nil
				}
				if path[1] == "MoneyOut" {
					inspector.AssignBuf(&x0.MoneyOut, value, buf)
					return nil
				}
				if path[1] == "Balance" {
					inspector.AssignBuf(&x0.Balance, value, buf)
					return nil
				}
				if path[1] == "AllowBuy" {
					inspector.AssignBuf(&x0.AllowBuy, value, buf)
					return nil
				}
				if path[1] == "History" {
					x1 := x0.History
					if uvalue, ok := value.(*[]testobj.TestHistory); ok {
						x1 = *uvalue
					}
					if x1 == nil {
						z := make([]testobj.TestHistory, 0)
						x1 = z
						x0.History = x1
					}
					_ = x1
					if len(path) > 2 {
						var i int
						t36, err36 := strconv.ParseInt(path[2], 0, 0)
						if err36 != nil {
							return err36
						}
						i = int(t36)
						if len(x1) > i {
							x2 := &(x1)[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									inspector.AssignBuf(&x2.DateUnix, value, buf)
									return nil
								}
								if path[3] == "Cost" {
									inspector.AssignBuf(&x2.Cost, value, buf)
									return nil
								}
								if path[3] == "Comment" {
									inspector.AssignBuf(&x2.Comment, value, buf)
									return nil
								}
							}
							(x1)[i] = *x2
							return nil
						}
					}
					x0.History = x1
				}
			}
			x.Finance = x0
		}
	}
	return nil
}

func (i3 *TestObjectInspector) Set(dst, value interface{}, path ...string) error {
	return i3.SetWB(dst, value, nil, path...)
}

func (i3 *TestObjectInspector) DeepEqual(l, r interface{}) bool {
	return i3.DeepEqualWithOptions(l, r, nil)
}

func (i3 *TestObjectInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestObject
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestObject); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestObject); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestObject); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestObject); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestObject); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestObject); ok {
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

	if lx.Id != rx.Id && inspector.DEQMustCheck("Id", opts) {
		return false
	}
	if !bytes.Equal(lx.Name, rx.Name) && inspector.DEQMustCheck("Name", opts) {
		return false
	}
	if lx.Status != rx.Status && inspector.DEQMustCheck("Status", opts) {
		return false
	}
	if lx.Ustate != rx.Ustate && inspector.DEQMustCheck("Ustate", opts) {
		return false
	}
	if lx.Cost != rx.Cost && inspector.DEQMustCheck("Cost", opts) {
		return false
	}
	lx1 := lx.Permission
	rx1 := rx.Permission
	_, _ = lx1, rx1
	if (lx1 == nil && rx1 != nil) || (lx1 != nil && rx1 == nil) {
		return false
	}
	if lx1 != nil && rx1 != nil {
		if len(*lx1) != len(*rx1) {
			return false
		}
		for k := range *lx1 {
			lx2 := (*lx1)[k]
			rx2, ok2 := (*rx1)[k]
			_, _, _ = lx2, rx2, ok2
			if !ok2 {
				return false
			}
			if lx2 != rx2 && inspector.DEQMustCheck("Permission", opts) {
				return false
			}
		}
	}
	lx3 := lx.HistoryTree
	rx3 := rx.HistoryTree
	_, _ = lx3, rx3
	if len(lx3) != len(rx3) {
		return false
	}
	for k := range lx3 {
		lx4 := (lx3)[k]
		rx4, ok4 := (rx3)[k]
		_, _, _ = lx4, rx4, ok4
		if !ok4 {
			return false
		}
		if (lx4 == nil && rx4 != nil) || (lx4 != nil && rx4 == nil) {
			return false
		}
		if lx4 != nil && rx4 != nil {
			if lx4.DateUnix != rx4.DateUnix && inspector.DEQMustCheck("HistoryTree.DateUnix", opts) {
				return false
			}
			if lx4.Cost != rx4.Cost && inspector.DEQMustCheck("HistoryTree.Cost", opts) {
				return false
			}
			if !bytes.Equal(lx4.Comment, rx4.Comment) && inspector.DEQMustCheck("HistoryTree.Comment", opts) {
				return false
			}
		}
	}
	lx5 := lx.Flags
	rx5 := rx.Flags
	_, _ = lx5, rx5
	if len(lx5) != len(rx5) {
		return false
	}
	for k := range lx5 {
		lx6 := (lx5)[k]
		rx6, ok6 := (rx5)[k]
		_, _, _ = lx6, rx6, ok6
		if !ok6 {
			return false
		}
		if lx6 != rx6 && inspector.DEQMustCheck("Flags", opts) {
			return false
		}
	}
	lx7 := lx.Finance
	rx7 := rx.Finance
	_, _ = lx7, rx7
	if (lx7 == nil && rx7 != nil) || (lx7 != nil && rx7 == nil) {
		return false
	}
	if lx7 != nil && rx7 != nil {
		if lx7.MoneyIn != rx7.MoneyIn && inspector.DEQMustCheck("Finance.MoneyIn", opts) {
			return false
		}
		if lx7.MoneyOut != rx7.MoneyOut && inspector.DEQMustCheck("Finance.MoneyOut", opts) {
			return false
		}
		if lx7.Balance != rx7.Balance && inspector.DEQMustCheck("Finance.Balance", opts) {
			return false
		}
		if lx7.AllowBuy != rx7.AllowBuy && inspector.DEQMustCheck("Finance.AllowBuy", opts) {
			return false
		}
		lx8 := lx7.History
		rx8 := rx7.History
		_, _ = lx8, rx8
		if len(lx8) != len(rx8) {
			return false
		}
		for i := 0; i < len(lx8); i++ {
			lx9 := (lx8)[i]
			rx9 := (rx8)[i]
			_, _ = lx9, rx9
			if lx9.DateUnix != rx9.DateUnix && inspector.DEQMustCheck("Finance.History.DateUnix", opts) {
				return false
			}
			if lx9.Cost != rx9.Cost && inspector.DEQMustCheck("Finance.History.Cost", opts) {
				return false
			}
			if !bytes.Equal(lx9.Comment, rx9.Comment) && inspector.DEQMustCheck("Finance.History.Comment", opts) {
				return false
			}
		}
	}
	return true
}
