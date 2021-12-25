package xunsafe

import (
	"unsafe"
)

//Addr returns field addr
func (f *Field) Addr(structPtr unsafe.Pointer) interface{} {
	return asInterface(f.Pointer(structPtr), f.rtypPtr, false)
}
