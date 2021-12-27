package xunsafe

import (
	"reflect"
	"unsafe"
)

//Field represents a field
type Field struct {
	Name string
	reflect.Type
	offset  uintptr
	kind    reflect.Kind
	rtype   *rtype
	rtypPtr *rtype
}

//Pointer return  field pointer (structPtr + field.Offset)
func (f *Field) Pointer(structPtr unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(structPtr) + f.offset)
}

//SafePointer returns field pointer, if field pointer is a pointer this method initialises that pointer
func (f *Field) SafePointer(structPtr unsafe.Pointer) unsafe.Pointer {
	if f.kind == reflect.Ptr {
		ptr := (*unsafe.Pointer)(f.Pointer(structPtr))
		if *ptr == nil {
			var newPointer unsafe.Pointer
			*ptr = unsafe.Pointer(&newPointer)
		}
	}
	return f.Pointer(structPtr)
}

//EnsurePointer initialises field type pointer if needed, and return filed type value pointer rather field pointer.
//for example if field is of T type this method returns *T, in case field is of *T, this method
//also return *T, if you need always field pointer use Field.Pointer method
func (f *Field) EnsurePointer(structPtr unsafe.Pointer) unsafe.Pointer {
	addr := f.Pointer(structPtr)
	ptr := (*unsafe.Pointer)(addr)
	if f.kind != reflect.Ptr {
		return addr
	}
	if *ptr == nil {
		var newPointer unsafe.Pointer
		*ptr = unsafe.Pointer(&newPointer)
	}
	return *ptr
}

func (f *Field) initType() {
	fType := f.Type
	ptrValue := reflect.New(fType)
	ptrElemValue := ptrValue.Elem()
	valPtr := ptrValue.Interface()
	val := ptrElemValue.Interface()
	f.rtypPtr = ((*emptyInterface)(unsafe.Pointer(&valPtr))).typ
	f.rtype = ((*emptyInterface)(unsafe.Pointer(&val))).typ
}

//NewField creates a new filed
func NewField(field reflect.StructField) *Field {
	fieldType := field.Type
	f := &Field{
		Name:   field.Name,
		Type:   fieldType,
		offset: field.Offset,
		kind:   fieldType.Kind(),
	}
	f.initType()
	return f
}

//FieldByIndex creates a field for supplied struct type and field indexAddr
func FieldByIndex(structType reflect.Type, index int) *Field {
	return NewField(structType.Field(index))
}

//FieldByName creates a field for supplied struct type and field name
func FieldByName(structType reflect.Type, name string) *Field {
	switch structType.Kind() {
	case reflect.Ptr:
		return FieldByName(structType.Elem(), name)
	case reflect.Slice:
		return FieldByName(structType.Elem(), name)
	}
	structField, ok := structType.FieldByName(name)
	if !ok {
		return nil
	}
	return NewField(structField)
}
