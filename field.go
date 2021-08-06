package xunsafe

import "reflect"

//Field represent a field
type Field struct {
	Field   *Field
	address Getter
	value   Getter
	field   reflect.StructField
}

//NewField creates a new filed
func NewField(field reflect.StructField) *Field {
	return &Field{
		field: field,
	}
}

//FieldByIndex creates a field for supplied struct type and field index
func FieldByIndex(structType reflect.Type, index int) *Field {
	return NewField(structType.Field(index))
}

//FieldByName creates a field for supplied struct type and field name
func FieldByName(structType reflect.Type, name string) *Field {
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
		value:   value,
	}
}




