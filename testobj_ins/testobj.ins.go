// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"bytes"
	"strconv"

	"github.com/koykov/fastconv"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
)

type TestFinanceInspector struct {
	inspector.BaseInspector
}

func (i0 *TestFinanceInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i0.GetTo(src, &buf, path...)
	return buf, err
}

func (i0 *TestFinanceInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFinance
	if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "MoneyIn" {
			*buf = &x.MoneyIn
			return
		}
		if path[0] == "MoneyOut" {
			*buf = &x.MoneyOut
			return
		}
		if path[0] == "Balance" {
			*buf = &x.Balance
			return
		}
		if path[0] == "AllowBuy" {
			*buf = &x.AllowBuy
			return
		}
		if path[0] == "History" {
			x0 := x.History
			_ = x0
			if len(path) > 1 {
				var i int
				t78, err78 := strconv.ParseInt(path[1], 0, 0)
				if err78 != nil {
					return err78
				}
				i = int(t78)
				if len(x0) > i {
					x1 := x0[i]
					_ = x1
					if len(path) > 2 {
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
	}
	return
}

func (i0 *TestFinanceInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFinance
	if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "MoneyIn" {
			var rightExact float64
			t21, err21 := strconv.ParseFloat(right, 0)
			if err21 != nil {
				return err21
			}
			rightExact = float64(t21)
			switch cond {
			case inspector.OpEq:
				*result = x.MoneyIn == rightExact
			case inspector.OpNq:
				*result = x.MoneyIn != rightExact
			case inspector.OpGt:
				*result = x.MoneyIn > rightExact
			case inspector.OpGtq:
				*result = x.MoneyIn >= rightExact
			case inspector.OpLt:
				*result = x.MoneyIn < rightExact
			case inspector.OpLtq:
				*result = x.MoneyIn <= rightExact
			}
			return
		}
		if path[0] == "MoneyOut" {
			var rightExact float64
			t21, err21 := strconv.ParseFloat(right, 0)
			if err21 != nil {
				return err21
			}
			rightExact = float64(t21)
			switch cond {
			case inspector.OpEq:
				*result = x.MoneyOut == rightExact
			case inspector.OpNq:
				*result = x.MoneyOut != rightExact
			case inspector.OpGt:
				*result = x.MoneyOut > rightExact
			case inspector.OpGtq:
				*result = x.MoneyOut >= rightExact
			case inspector.OpLt:
				*result = x.MoneyOut < rightExact
			case inspector.OpLtq:
				*result = x.MoneyOut <= rightExact
			}
			return
		}
		if path[0] == "Balance" {
			var rightExact float64
			t21, err21 := strconv.ParseFloat(right, 0)
			if err21 != nil {
				return err21
			}
			rightExact = float64(t21)
			switch cond {
			case inspector.OpEq:
				*result = x.Balance == rightExact
			case inspector.OpNq:
				*result = x.Balance != rightExact
			case inspector.OpGt:
				*result = x.Balance > rightExact
			case inspector.OpGtq:
				*result = x.Balance >= rightExact
			case inspector.OpLt:
				*result = x.Balance < rightExact
			case inspector.OpLtq:
				*result = x.Balance <= rightExact
			}
			return
		}
		if path[0] == "AllowBuy" {
			var rightExact bool
			t23, err23 := strconv.ParseBool(right)
			if err23 != nil {
				return err23
			}
			rightExact = bool(t23)
			if cond == inspector.OpEq {
				*result = x.AllowBuy == rightExact
			} else {
				*result = x.AllowBuy != rightExact
			}
			return
		}
		if path[0] == "History" {
			x0 := x.History
			_ = x0
			if len(path) > 1 {
				var i int
				t78, err78 := strconv.ParseInt(path[1], 0, 0)
				if err78 != nil {
					return err78
				}
				i = int(t78)
				if len(x0) > i {
					x1 := x0[i]
					_ = x1
					if len(path) > 2 {
						if path[2] == "DateUnix" {
							var rightExact int64
							t15, err15 := strconv.ParseInt(right, 0, 0)
							if err15 != nil {
								return err15
							}
							rightExact = int64(t15)
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
							t21, err21 := strconv.ParseFloat(right, 0)
							if err21 != nil {
								return err21
							}
							rightExact = float64(t21)
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
	}
	return
}

func (i0 *TestFinanceInspector) Set(dst, value interface{}, path ...string) {
}

type TestFlagInspector struct {
	inspector.BaseInspector
}

func (i1 *TestFlagInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i1.GetTo(src, &buf, path...)
	return buf, err
}

func (i1 *TestFlagInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFlag
	if p, ok := src.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := src.(testobj.TestFlag); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[path[0]]; ok {
			_ = x0
			*buf = &x0
			return
		}
	}
	*buf = x
	return
}

func (i1 *TestFlagInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFlag
	if p, ok := src.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := src.(testobj.TestFlag); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[path[0]]; ok {
			_ = x0
			var rightExact int32
			t13, err13 := strconv.ParseInt(right, 0, 0)
			if err13 != nil {
				return err13
			}
			rightExact = int32(t13)
			switch cond {
			case inspector.OpEq:
				*result = x0 == rightExact
			case inspector.OpNq:
				*result = x0 != rightExact
			case inspector.OpGt:
				*result = x0 > rightExact
			case inspector.OpGtq:
				*result = x0 >= rightExact
			case inspector.OpLt:
				*result = x0 < rightExact
			case inspector.OpLtq:
				*result = x0 <= rightExact
			}
			return
		}
	}
	return
}

func (i1 *TestFlagInspector) Set(dst, value interface{}, path ...string) {
}

type TestHistoryInspector struct {
	inspector.BaseInspector
}

func (i2 *TestHistoryInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	return nil
}

func (i2 *TestHistoryInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i2.GetTo(src, &buf, path...)
	return buf, err
}

func (i2 *TestHistoryInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			*buf = &x.DateUnix
			return
		}
		if path[0] == "Cost" {
			*buf = &x.Cost
			return
		}
		if path[0] == "Comment" {
			*buf = &x.Comment
			return
		}
	}
	return
}

func (i2 *TestHistoryInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			var rightExact int64
			t15, err15 := strconv.ParseInt(right, 0, 0)
			if err15 != nil {
				return err15
			}
			rightExact = int64(t15)
			switch cond {
			case inspector.OpEq:
				*result = x.DateUnix == rightExact
			case inspector.OpNq:
				*result = x.DateUnix != rightExact
			case inspector.OpGt:
				*result = x.DateUnix > rightExact
			case inspector.OpGtq:
				*result = x.DateUnix >= rightExact
			case inspector.OpLt:
				*result = x.DateUnix < rightExact
			case inspector.OpLtq:
				*result = x.DateUnix <= rightExact
			}
			return
		}
		if path[0] == "Cost" {
			var rightExact float64
			t21, err21 := strconv.ParseFloat(right, 0)
			if err21 != nil {
				return err21
			}
			rightExact = float64(t21)
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
		if path[0] == "Comment" {
			var rightExact []byte
			rightExact = fastconv.S2B(right)

			if cond == inspector.OpEq {
				*result = bytes.Equal(x.Comment, rightExact)
			} else {
				*result = !bytes.Equal(x.Comment, rightExact)
			}
			return
		}
	}
	return
}

func (i2 *TestHistoryInspector) Set(dst, value interface{}, path ...string) {
}

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
						t78, err78 := strconv.ParseInt(path[2], 0, 0)
						if err78 != nil {
							return err78
						}
						i = int(t78)
						if len(x1) > i {
							x2 := x1[i]
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
			t13, err13 := strconv.ParseInt(right, 0, 0)
			if err13 != nil {
				return err13
			}
			rightExact = int32(t13)
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
			t21, err21 := strconv.ParseFloat(right, 0)
			if err21 != nil {
				return err21
			}
			rightExact = float64(t21)
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
				t13, err13 := strconv.ParseInt(path[1], 0, 0)
				if err13 != nil {
					return err13
				}
				k = int32(t13)
				x1 := (*x0)[k]
				_ = x1
				var rightExact bool
				t23, err23 := strconv.ParseBool(right)
				if err23 != nil {
					return err23
				}
				rightExact = bool(t23)
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
							t15, err15 := strconv.ParseInt(right, 0, 0)
							if err15 != nil {
								return err15
							}
							rightExact = int64(t15)
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
							t21, err21 := strconv.ParseFloat(right, 0)
							if err21 != nil {
								return err21
							}
							rightExact = float64(t21)
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
					t13, err13 := strconv.ParseInt(right, 0, 0)
					if err13 != nil {
						return err13
					}
					rightExact = int32(t13)
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
					t21, err21 := strconv.ParseFloat(right, 0)
					if err21 != nil {
						return err21
					}
					rightExact = float64(t21)
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
					t21, err21 := strconv.ParseFloat(right, 0)
					if err21 != nil {
						return err21
					}
					rightExact = float64(t21)
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
					t21, err21 := strconv.ParseFloat(right, 0)
					if err21 != nil {
						return err21
					}
					rightExact = float64(t21)
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
					t23, err23 := strconv.ParseBool(right)
					if err23 != nil {
						return err23
					}
					rightExact = bool(t23)
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
						t78, err78 := strconv.ParseInt(path[2], 0, 0)
						if err78 != nil {
							return err78
						}
						i = int(t78)
						if len(x1) > i {
							x2 := x1[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									var rightExact int64
									t15, err15 := strconv.ParseInt(right, 0, 0)
									if err15 != nil {
										return err15
									}
									rightExact = int64(t15)
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
									t21, err21 := strconv.ParseFloat(right, 0)
									if err21 != nil {
										return err21
									}
									rightExact = float64(t21)
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
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				for k := range *x0 {
					l.Set("k", &k, &inspector.StaticInspector{})
					l.Set("item", (*x0)[k], &inspector.StaticInspector{})
					l.Loop()
				}
				return
			}
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			if len(path) > 1 {
				for k := range x0 {
					l.Set("k", &k, &inspector.StaticInspector{})
					l.Set("item", x0[k], &TestHistoryInspector{})
					l.Loop()
				}
			}
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			// if len(path) > 1 {
			for k := range x0 {
				*buf = append((*buf)[:0], k...)
				l.Set("k", buf, &inspector.StaticInspector{})
				l.Set("v", x0[k], &inspector.StaticInspector{})
				l.Loop()
			}
			// }
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
					for i := range x1 {
						*buf = strconv.AppendInt((*buf)[:0], int64(i), 10)
						l.Set("k", buf, &inspector.StaticInspector{})
						l.Set("item", &x1[i], &TestHistoryInspector{})
						l.Loop()
					}
				}
			}
		}
	}
	return
}

func (i3 *TestObjectInspector) Set(dst, value interface{}, path ...string) {
}

type TestPermissionInspector struct {
	inspector.BaseInspector
}

func (i4 *TestPermissionInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i4.GetTo(src, &buf, path...)
	return buf, err
}

func (i4 *TestPermissionInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var k int32
		t13, err13 := strconv.ParseInt(path[0], 0, 0)
		if err13 != nil {
			return err13
		}
		k = int32(t13)
		x0 := (*x)[k]
		_ = x0
		*buf = &x0
		return
	}
	*buf = x
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
	if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var k int32
		t13, err13 := strconv.ParseInt(path[0], 0, 0)
		if err13 != nil {
			return err13
		}
		k = int32(t13)
		x0 := (*x)[k]
		_ = x0
		var rightExact bool
		t23, err23 := strconv.ParseBool(right)
		if err23 != nil {
			return err23
		}
		rightExact = bool(t23)
		if cond == inspector.OpEq {
			*result = x0 == rightExact
		} else {
			*result = x0 != rightExact
		}
		return
	}
	return
}

func (i4 *TestPermissionInspector) Set(dst, value interface{}, path ...string) {
}
