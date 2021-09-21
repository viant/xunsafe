package eval

import (
	"fmt"
	"github.com/viant/xunsafe"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type Selector struct {
	child            *Selector
	name             string
	sliceDataAddress func(structAddr unsafe.Pointer) unsafe.Pointer
	index            *uintptr
	itemType         reflect.Type
	field            *xunsafe.Field
}

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

//StringAddr returns field *string addr
func (s *Selector) StringAddr(structAddr unsafe.Pointer) *string {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.StringAddr(structAddr)
	}
	return s.child.StringAddr(s.field.UnsafeAddr(structAddr))
}

//IntAddr returns field *int address
func (s *Selector) BoolAddr(structAddr unsafe.Pointer) *bool {
	if s.index != nil {
		structAddr = s.field.UnsafeAddr(s.sliceDataAddress(structAddr))
	} else if s.child == nil {
		return s.field.BoolAddr(structAddr)
	}
	return s.child.BoolAddr(s.field.UnsafeAddr(structAddr))
}

func NewSelector(owner reflect.Type, name string) (*Selector, error) {
	subNode := strings.Index(name, ".")
	itemNode := strings.Index(name, "[")
	child := ""
	var idx *uintptr
	if itemNode != -1 && itemNode < subNode {
		itemIdx, err := strconv.Atoi(name[itemNode+1 : subNode-1])
		if err != nil {
			return nil, fmt.Errorf("invalid selector: %v index: %v", name, err)
		}
		offset := uintptr(itemIdx)
		idx = &offset
		child = name[subNode+1:]
		name = name[:itemNode]
		subNode = strings.Index(name, ".")
	}

	if subNode != -1 {
		child = name[subNode+1:]
		name = name[:subNode]
	}

	result := &Selector{name: name, index: idx}
	result.field = xunsafe.FieldByName(owner, result.name)
	if result.field == nil {
		return nil, fmt.Errorf("failed to lookup %v.%v", owner.Name(), name)
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
			return nil, fmt.Errorf("failed to lookup %v.%v, %w", owner.Name(), name, err)
		}
	}
	return result, err
}
