package xunsafe

import (
	"reflect"
	"unsafe"
)

//Type represents a Type
type Type struct {
	typ      reflect.Type
	rtype    *rtype
	rtypePtr *rtype
}

//Type returns reflect type
func (t *Type) Type() reflect.Type {
	return t.typ
}

//Deref dereference pointer
func (t *Type) Deref(val interface{}) interface{} {
	ptr := AsPointer(val)
	return asInterface(ptr, t.rtype, true)
}

//Pointer returns a pointer
func (t *Type) Pointer(value interface{}) interface{} {
	e := (*emptyInterface)(unsafe.Pointer(&value))
	if e.typ.kind&kindDirectIface != 0 {
		var newPtr unsafe.Pointer
		updated := unsafe.Pointer(&newPtr)
		*(*unsafe.Pointer)(updated) = e.word
		e.word = updated
	}
	e.typ = t.rtypePtr
	return value
}

//NewType creates a type
func NewType(t reflect.Type) *Type {
	ptrValue := reflect.New(t)
	ptrElemValue := ptrValue.Elem()
	valPtr := ptrValue.Interface()
	val := ptrElemValue.Interface()
	result := &Type{
		typ:      t,
		rtypePtr: ((*emptyInterface)(unsafe.Pointer(&valPtr))).typ,
		rtype:    ((*emptyInterface)(unsafe.Pointer(&val))).typ,
	}
	return result
}
