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
		return ValuePointer(&value)
	default:
		newPtr := reflect.New(value.Type())
		newPtr.Elem().Set(value)
		return ValuePointer(&newPtr)
	}
}

//SafeDerefPointer returns deref pointer (**T -> *T), pType has to be *T
func SafeDerefPointer(pointer unsafe.Pointer, pType reflect.Type) unsafe.Pointer {
	ptr := (*unsafe.Pointer)(pointer)
	if *ptr == nil {
		n := reflect.New(pType.Elem())
		nPtr := ValuePointer(&n)
		*ptr = nPtr
	}
	return *ptr
}

//DerefPointer returns deref pointer (**T -> *T)
func DerefPointer(pointer unsafe.Pointer) unsafe.Pointer {
	return *(*unsafe.Pointer)(pointer)
}

//RefPointer returns reference to the pointer (*T -> **T)
func RefPointer(pointer unsafe.Pointer) unsafe.Pointer {
	var newPtr unsafe.Pointer
	updated := unsafe.Pointer(&newPtr)
	*(*unsafe.Pointer)(updated) = pointer
	return updated
}

type (
	wrapper struct {
		value interface{}
	}
)

//AsPointer returns a  pointer for an empty interface
func AsPointer(v interface{}) unsafe.Pointer {
	empty := (*emptyInterface)(unsafe.Pointer(&v))
	word := empty.word
	//32 directIface flag
	if empty.typ != nil && empty.typ.kind&32 == 32 && empty.typ.kind&byte(reflect.Struct) == 25 {
		word = RefPointer(word)
		//aValue := wrapper{value: v}
		//word = unsafe.Pointer(&aValue)
	}

	return word
}

func asPointer(v interface{}) unsafe.Pointer {
	empty := (*emptyInterface)(unsafe.Pointer(&v))
	return empty.word
}

//EnsureAddressPointer ensure that address pointer is not nil, ptr has to be address pointer
func EnsureAddressPointer(addrPtr unsafe.Pointer, target reflect.Type) *unsafe.Pointer {
	itemPtr := (*unsafe.Pointer)(addrPtr)
	if *itemPtr != nil {
		return itemPtr
	}

	newValue := reflect.New(target)
	newPtr := ValuePointer(&newValue)
	*itemPtr = unsafe.Pointer(&newPtr)
	return itemPtr
}

const n = 8192

//Copy k bytes from src to dest
//go:nocheckptr
func Copy(dest, src unsafe.Pointer, k int) {
	bsLen := k
	chunks := bsLen / n
	offset := uintptr(0)
	for i := 0; i < chunks; i++ {
		copy((*(*[n]byte)(unsafe.Pointer(uintptr(dest) + offset)))[:n], (*(*[n]byte)(unsafe.Pointer(uintptr(src) + offset)))[:n])
		offset += n
	}
	limit := bsLen % n
	if limit == 0 {
		return
	}
	copy((*(*[n]byte)(unsafe.Pointer(uintptr(dest) + offset)))[:limit], (*(*[n]byte)(unsafe.Pointer(uintptr(src) + offset)))[:limit])
}
