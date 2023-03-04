//go:build debug

package xunsafe

import (
	"fmt"
	"reflect"
	"unsafe"
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

//Set sets only non pointer value, the reason for this limited functionality method is speed,
//its 20x faster than SetValue
//go:nocheckptr
func (f *Field) Set(structPtr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)

	ptr := f.Pointer(structPtr)
	switch f.kind {
	case reflect.String:
		*(*string)(ptr) = source.(string)
	case reflect.Int:
		*(*int)(ptr) = source.(int)
	case reflect.Int64:
		*(*int64)(ptr) = source.(int64)
	case reflect.Float64:
		*(*float64)(ptr) = source.(float64)
	case reflect.Float32:
		*(*float32)(ptr) = source.(float32)
	case reflect.Bool:
		*(*bool)(ptr) = source.(bool)
	case reflect.Ptr: //had to comment out this cast since this suppresses inlining
		//*(*unsafe.Pointer)(ptr) = AsPointer(source)
	default:
		*(*unsafe.Pointer)(ptr) = *(*unsafe.Pointer)(asPointer(source))
	}
}
