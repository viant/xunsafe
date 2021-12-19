package xunsafe

import (
	"reflect"
)

//Field represent a field
type Field struct {
	Field   *Field
	address Getter
	Value   Getter

	setter Setter
	field  reflect.StructField
	Type   reflect.Type
	kind   reflect.Kind
}

//NewField creates a new filed
func NewField(field reflect.StructField) *Field {
	fType := field.Type
	f := &Field{
		field: field,
		Type:  fType,
		kind:  fType.Kind(),
	}

	f.Value = FieldAccessor(f)
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

//FieldWithGetters creates a field supplied custom address, value getter
func FieldWithGetters(address, value Getter) *Field {
	return &Field{
		address: address,
		Value:   value,
	}
}
