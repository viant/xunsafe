package xunsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Slice struct {
	reflect.Type
	sliceDataAddress func(structAddr unsafe.Pointer, index uintptr) unsafe.Pointer
}

func (r *Slice) Range(addr unsafe.Pointer, visit func(index int, addr unsafe.Pointer) bool) {
	header := *(*reflect.SliceHeader)(addr)
	for i := 0; i < header.Len; i++ {
		if !visit(i, r.sliceDataAddress(addr, uintptr(i))) {
			return
		}
	}
}

func (r *Slice) Index(addr unsafe.Pointer, index int) unsafe.Pointer {
	return r.sliceDataAddress(addr, uintptr(index))
}

//NewSlice creates  slice
func NewSlice(aType reflect.Type) *Slice {
	switch aType.Kind() {
	case reflect.Slice:
	case reflect.Array:
		panic(fmt.Sprintf("unsupported type: %v", aType.Name()))
	default:
		aType = reflect.SliceOf(aType)
	}
	itemType := aType.Elem()
	size := itemType.Size()
	result := &Slice{
		Type: aType,
	}
	if itemType.Kind() == reflect.Ptr {
		result.sliceDataAddress = func(structAddr unsafe.Pointer, index uintptr) unsafe.Pointer {
			header := *(*reflect.SliceHeader)(structAddr)
			offset := header.Data - uintptr(structAddr) + index*size
			return *(*unsafe.Pointer)(unsafe.Add(structAddr, offset))
		}
	} else {
		result.sliceDataAddress = func(structAddr unsafe.Pointer, index uintptr) unsafe.Pointer {
			header := *(*reflect.SliceHeader)(structAddr)
			offset := header.Data - uintptr(structAddr) + index*size
			return unsafe.Add(structAddr, offset)
		}
	}
	return result
}
