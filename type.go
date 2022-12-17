package xunsafe

import (
	"reflect"
	"unsafe"
)

//Type represents a Type
type Type struct {
	typ      reflect.Type
	isError  bool
	rtype    *rtype
	rtypePtr *rtype
	kind     reflect.Kind
	canDeref bool
}

//Type returns reflect type
func (t *Type) Type() reflect.Type {
	return t.typ
}

//Kind returns reflect kind
func (t *Type) Kind() reflect.Kind {
	return t.kind
}

//Value returns an interface for the pointer
func (t *Type) Value(ptr unsafe.Pointer, deref bool) interface{} {
	if t.isError {
		return AsError(ptr)
	}
	return asInterface(ptr, t.rtype, deref)
}

//Interface returns an interface for the pointer
func (t *Type) Interface(ptr unsafe.Pointer) interface{} {
	if t.isError {
		return AsError(ptr)
	}
	return asInterface(ptr, t.rtype, false)
}

//Deref dereference pointer
func (t *Type) Deref(val interface{}) interface{} {
	ptr := AsPointer(val)
	return asInterface(ptr, t.rtype, true)
}

//Pointer returns a pointer
func (t *Type) Pointer(value interface{}) unsafe.Pointer {
	return AsPointer(value)
}

//Ref returns a reference to value
func (t *Type) Ref(value interface{}) interface{} {
	e := (*emptyInterface)(unsafe.Pointer(&value))
	if e.typ.kind&kindDirectIface != 0 {
		e.word = RefPointer(e.word)
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

	canDeref := t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Ptr
	result := &Type{
		isError:  t.Name() == "error",
		typ:      t,
		kind:     t.Kind(),
		canDeref: canDeref,
		rtypePtr: ((*emptyInterface)(unsafe.Pointer(&valPtr))).typ,
		rtype:    ((*emptyInterface)(unsafe.Pointer(&val))).typ,
	}
	return result
}
