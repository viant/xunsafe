package xunsafe

import (
	"reflect"
	"unsafe"
)

//EnsurePointer returns unsafe pointer for src value or value pointer
//if you guarantee src is already pointer AsPointer is much faster option
func EnsurePointer(src interface{}) unsafe.Pointer {
	value := reflect.ValueOf(src)
	switch value.Kind() {
	case reflect.UnsafePointer:
		return src.(unsafe.Pointer)
	case reflect.Ptr:
		return valuePointer(&value)
	default:
		newPtr := reflect.New(value.Type())
		newPtr.Elem().Set(value)
		return valuePointer(&newPtr)
	}
}

//DerefPointer returns deref pointer
func DerefPointer(pointer unsafe.Pointer) unsafe.Pointer {
	return *(*unsafe.Pointer)(pointer)
}

//AsPointer returns a  pointer for an interface
func AsPointer(v interface{}) unsafe.Pointer {
	empty := (*emptyInterface)(unsafe.Pointer(&v))
	return empty.word
}
