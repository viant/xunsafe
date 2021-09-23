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
	sliceDataAddress func(structAddr unsafe.Pointer) unsafe.Pointer
	index            *uintptr
	itemType         reflect.Type
	field            *Field
}

//Type returns field type
func (s *Selector) Type() reflect.Type {
	if s.child == nil {
		return s.field.Type
	}
	return s.child.Type()
}

//Set sets path value
func (s *Selector) Set(structAddr unsafe.Pointer, val interface{}) {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		s.field.Set(structAddr, val)
		return
	}
	s.child.Set(s.field.UnsafeAddr(structAddr), val)
}

//IntAddr returns field *int address
func (s *Selector) IntAddr(structAddr unsafe.Pointer) *int {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.IntAddr(structAddr)
	}
	return s.child.IntAddr(s.field.UnsafeAddr(structAddr))
}

//Value returns field value
func (s *Selector) Value(structAddr unsafe.Pointer) interface{} {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.Value(structAddr)
	}
	return s.child.Value(s.field.UnsafeAddr(structAddr))
}

//Int returns field int value
func (s *Selector) Int(structAddr unsafe.Pointer) int {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.Int(structAddr)
	}
	return s.child.Int(s.field.UnsafeAddr(structAddr))
}

//Float64Addr returns field *float64 address
func (s *Selector) Float64Addr(structAddr unsafe.Pointer) *float64 {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.Float64Addr(structAddr)
	}
	return s.child.Float64Addr(s.field.UnsafeAddr(structAddr))
}

//Float64 returns field float64 value
func (s *Selector) Float64(structAddr unsafe.Pointer) float64 {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.Float64(structAddr)
	}
	return s.child.Float64(s.field.UnsafeAddr(structAddr))
}

//StringAddr returns field *string addr
func (s *Selector) StringAddr(structAddr unsafe.Pointer) *string {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.StringAddr(structAddr)
	}
	return s.child.StringAddr(s.field.UnsafeAddr(structAddr))
}

//String returns field int value
func (s *Selector) String(structAddr unsafe.Pointer) string {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.String(structAddr)
	}
	return s.child.String(s.field.UnsafeAddr(structAddr))
}

//BoolAddr returns field *bool address
func (s *Selector) BoolAddr(structAddr unsafe.Pointer) *bool {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.BoolAddr(structAddr)
	}
	return s.child.BoolAddr(s.field.UnsafeAddr(structAddr))
}

//Bool returns field bool value
func (s *Selector) Bool(structAddr unsafe.Pointer) bool {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.Bool(structAddr)
	}
	return s.child.Bool(s.field.UnsafeAddr(structAddr))
}

//NewSelector creates a selector for supplied expression
func NewSelector(owner reflect.Type, expr string) (*Selector, error) {
	subNode := strings.Index(expr, ".")
	itemNode := strings.Index(expr, "[")
	child := ""
	var idx *uintptr
	if itemNode != -1 && itemNode < subNode {
		itemIdx, err := strconv.Atoi(expr[itemNode+1 : subNode-1])
		if err != nil {
			return nil, fmt.Errorf("invalid selector: %v index: %v", expr, err)
		}
		offset := uintptr(itemIdx)
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
		field, _ := owner.FieldByName(result.name)
		result.itemType = field.Type.Elem()
		if result.itemType.Kind() == reflect.Ptr {
			result.sliceDataAddress = func(structAddr unsafe.Pointer) unsafe.Pointer {
				header := *(*reflect.SliceHeader)(structAddr)
				offset := header.Data - uintptr(structAddr) + *result.index*result.itemType.Size()
				return *(*unsafe.Pointer)(unsafe.Add(structAddr, offset))
			}
		} else {
			result.sliceDataAddress = func(structAddr unsafe.Pointer) unsafe.Pointer {
				header := *(*reflect.SliceHeader)(structAddr)
				offset := header.Data - uintptr(structAddr) + *result.index*result.itemType.Size()
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
