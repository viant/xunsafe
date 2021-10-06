package xunsafe

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

//Selector represents nested abstraction selector
type Selector struct {
	child            *Selector
	name             string
	sliceDataAddress func(structAddr unsafe.Pointer, index int) unsafe.Pointer
	index            *int
	itemType         reflect.Type
	field            *Field
	_getValue        func(structAddr unsafe.Pointer) interface{}
	_setValue        func(structAddr unsafe.Pointer, val interface{})
}

//Type returns field type
func (s *Selector) Type() reflect.Type {
	if s.child == nil {
		return s.field.Type
	}
	return s.child.Type()
}

//ISet sets path value with optional slice indexes
func (s *Selector) ISet(structAddr unsafe.Pointer, val interface{}, indexes []int) {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), indexes[0])
		s.child.ISet(structAddr, val, indexes[1:])
		return
	} else if s.child == nil {
		s.field.Set(structAddr, val)
		return
	}
	s.child.ISet(s.field.UnsafeAddr(structAddr), val, indexes)
}

//IValue returns field value
func (s *Selector) IValue(structAddr unsafe.Pointer, indexes []int) interface{} {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), indexes[0])
		return s.child.IValue(structAddr, indexes[1:])
	} else if s.child == nil {
		return s.field.Value(structAddr)
	}
	return s.child.IValue(s.field.UnsafeAddr(structAddr), indexes)
}

//Set sets path value
func (s *Selector) Set(structAddr unsafe.Pointer, val interface{}) {
	if s._setValue != nil {
		s._setValue(structAddr, val)
		return
	}

	var fn func(structAddr unsafe.Pointer, val interface{})
	if s.index != nil {
		fn = func(structAddr unsafe.Pointer, val interface{}) {
			structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
			s.child.Set(structAddr, val)
		}
	} else if s.child == nil {
		fn = func(structAddr unsafe.Pointer, val interface{}) {
			s.field.Set(structAddr, val)
		}
	} else {
		fn = func(structAddr unsafe.Pointer, val interface{}) {
			s.child.Set(s.field.UnsafeAddr(structAddr), val)
		}
	}
	s._setValue = fn
	fn(structAddr, val)
}

//Value returns field value
func (s *Selector) Value(structAddr unsafe.Pointer) interface{} {
	if s._getValue != nil {
		return s._getValue(structAddr)
	}
	var fn func(structAddr unsafe.Pointer) interface{}
	if s.index != nil {
		fn = func(structAddr unsafe.Pointer) interface{} {
			structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
			return s.child.Value(structAddr)
		}
	} else if s.child == nil {
		fn = func(structAddr unsafe.Pointer) interface{} {
			return s.field.Value(structAddr)
		}
	} else {
		fn = func(structAddr unsafe.Pointer) interface{} {
			return s.child.Value(s.field.UnsafeAddr(structAddr))
		}
	}
	s._getValue = fn
	return fn(structAddr)
}

//IntAddr returns field *int address
func (s *Selector) IntAddr(structAddr unsafe.Pointer) *int {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
	} else if s.child == nil {
		return s.field.IntAddr(structAddr)
	}
	return s.child.IntAddr(s.field.UnsafeAddr(structAddr))
}

//Int returns field int value
func (s *Selector) Int(structAddr unsafe.Pointer) int {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.Int(structAddr)
	} else if s.child == nil {
		return s.field.Int(structAddr)
	}
	return s.child.Int(s.field.UnsafeAddr(structAddr))
}

//SetInt sets int value
func (s *Selector) SetInt(structAddr unsafe.Pointer, val int) {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		s.child.SetInt(structAddr, val)
		return
	} else if s.child == nil {
		s.field.SetInt(structAddr, val)
		return
	}
	s.child.SetInt(s.field.UnsafeAddr(structAddr), val)
}

//Float64Addr returns field *float64 address
func (s *Selector) Float64Addr(structAddr unsafe.Pointer) *float64 {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.Float64Addr(structAddr)
	} else if s.child == nil {
		return s.field.Float64Addr(structAddr)
	}
	return s.child.Float64Addr(s.field.UnsafeAddr(structAddr))
}

//Float64 returns field float64 value
func (s *Selector) Float64(structAddr unsafe.Pointer) float64 {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.Float64(structAddr)
	} else if s.child == nil {
		return s.field.Float64(structAddr)
	}
	return s.child.Float64(s.field.UnsafeAddr(structAddr))
}

//SetFloat64 sets int value
func (s *Selector) SetFloat64(structAddr unsafe.Pointer, val float64) {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		s.child.SetFloat64(structAddr, val)
		return
	} else if s.child == nil {
		s.field.SetFloat64(structAddr, val)
		return
	}
	s.child.SetFloat64(s.field.UnsafeAddr(structAddr), val)
}

//StringAddr returns field *string addr
func (s *Selector) StringAddr(structAddr unsafe.Pointer) *string {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.StringAddr(structAddr)
	} else if s.child == nil {
		return s.field.StringAddr(structAddr)
	}
	return s.child.StringAddr(s.field.UnsafeAddr(structAddr))
}

//String returns field int value
func (s *Selector) String(structAddr unsafe.Pointer) string {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.String(structAddr)
	} else if s.child == nil {
		return s.field.String(structAddr)
	}
	return s.child.String(s.field.UnsafeAddr(structAddr))
}

//SetString sets string value
func (s *Selector) SetString(structAddr unsafe.Pointer, val string) {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		s.child.SetString(structAddr, val)
		return
	} else if s.child == nil {
		s.field.SetString(structAddr, val)
		return
	}
	s.child.SetString(s.field.UnsafeAddr(structAddr), val)
}

//BoolAddr returns field *bool address
func (s *Selector) BoolAddr(structAddr unsafe.Pointer) *bool {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.BoolAddr(structAddr)
	} else if s.child == nil {
		return s.field.BoolAddr(structAddr)
	}
	return s.child.BoolAddr(s.field.UnsafeAddr(structAddr))
}

//Bool returns field bool value
func (s *Selector) Bool(structAddr unsafe.Pointer) bool {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		return s.child.Bool(structAddr)
	} else if s.child == nil {
		return s.field.Bool(structAddr)
	}
	return s.child.Bool(s.field.UnsafeAddr(structAddr))
}

//SetBool sets bool value
func (s *Selector) SetBool(structAddr unsafe.Pointer, val bool) {
	if s.index != nil {
		structAddr = s.sliceDataAddress(s.field.UnsafeAddr(structAddr), *s.index)
		s.child.SetBool(structAddr, val)
		return
	} else if s.child == nil {
		s.field.SetBool(structAddr, val)
		return
	}
	s.child.SetBool(s.field.UnsafeAddr(structAddr), val)
}

//SetFiled set selector field
func (s *Selector) SetFiled(field *Field) {
	s.field = field
}

//NewSelector creates a selector for supplied expression
func NewSelector(owner reflect.Type, expr string) (*Selector, error) {
	subNode := strings.Index(expr, ".")
	itemNode := strings.Index(expr, "[")
	child := ""
	var idx *int
	if itemNode != -1 && itemNode < subNode {
		indexLit := expr[itemNode+1 : subNode-1]
		offset := -1
		var err error
		if indexLit != "" {
			offset, err = strconv.Atoi(indexLit)
			if err != nil {
				return nil, fmt.Errorf("invalid selector: %v index: %v", expr, err)
			}
		}
		idx = &offset
		child = expr[subNode+1:]
		expr = expr[:itemNode]
		subNode = strings.Index(expr, ".")
	}

	if subNode != -1 {
		child = expr[subNode+1:]
		expr = expr[:subNode]
	}
	result := &Selector{name: expr, index: idx}
	result.field = FieldByName(owner, result.name)
	if result.field == nil {
		ownerName := owner.Name()
		if ownerName != "" {
			ownerName += "."
		}
		return nil, fmt.Errorf("failed to lookup %v%v", ownerName, expr)
	}

	if idx != nil {
		if owner.Kind() == reflect.Ptr {
			owner = owner.Elem()
		}
		if owner.Kind() == reflect.Slice {
			field, _ := owner.Elem().FieldByName(result.name)
			result.itemType = field.Type.Elem()

		} else {
			field, _ := owner.FieldByName(result.name)
			result.itemType = field.Type.Elem()
		}

		if result.itemType.Kind() == reflect.Ptr {
			result.sliceDataAddress = func(structAddr unsafe.Pointer, index int) unsafe.Pointer {
				header := *(*reflect.SliceHeader)(structAddr)
				offset := header.Data - uintptr(structAddr) + uintptr(index)*result.itemType.Size()
				return *(*unsafe.Pointer)(unsafe.Add(structAddr, offset))
			}
		} else {
			result.sliceDataAddress = func(structAddr unsafe.Pointer, index int) unsafe.Pointer {
				header := *(*reflect.SliceHeader)(structAddr)
				offset := header.Data - uintptr(structAddr) + uintptr(index)*result.itemType.Size()
				return unsafe.Add(structAddr, offset)
			}
		}
	}

	var err error
	if len(child) > 1 {
		result.child, err = NewSelector(result.field.Type, child)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup %v.%v, %w", owner.Name(), expr, err)
		}
	}
	return result, err
}
