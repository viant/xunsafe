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

//AsPointer returns a  pointer for an empty interface
func AsPointer(v interface{}) unsafe.Pointer {
	empty := (*emptyInterface)(unsafe.Pointer(&v))
	return empty.word
}

//EnsureAddressPointer ensure that address pointer is not nil, ptr has to be address pointer
func EnsureAddressPointer(addrPtr unsafe.Pointer) *unsafe.Pointer {
	itemPtr := (*unsafe.Pointer)(addrPtr)
	if *itemPtr != nil {
		return itemPtr
	}
	var newPtr unsafe.Pointer
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
