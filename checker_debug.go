//go:build debug

package xunsafe

import (
	"fmt"
	"reflect"
)

func (f *Field) MustBeAssignable(y interface{}) {
	xType := f.Type
	var yType reflect.Type
	var ok bool
	yType, ok = y.(reflect.Type)
	if !ok {
		yType = reflect.TypeOf(y)
	}

	if xType != yType && xType.Kind() != reflect.Interface {
		panic(fmt.Errorf("xunsafe.SetValue: types mismatch: wanted %v, got %v", xType.String(), yType.String()))
	}
}
