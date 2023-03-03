//go:build debug

package xunsafe

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

//SetError sets field error
func (f *Field) SetError(structPtr unsafe.Pointer, val error) {
	f.MustBeAssignable(val)
	*AsErrorPtr(f.Pointer(structPtr)) = val
}

//SetInt sets field int
func (f *Field) SetInt(structPtr unsafe.Pointer, val int) {
	f.MustBeAssignable(val)
	*AsIntPtr(f.Pointer(structPtr)) = val
}

//SetIntPtr sets field *int
func (f *Field) SetIntPtr(structPtr unsafe.Pointer, val *int) {
	f.MustBeAssignable(val)
	*AsIntAddrPtr(f.Pointer(structPtr)) = val
}

//SetInt64 sets field int
func (f *Field) SetInt64(structPtr unsafe.Pointer, val int64) {
	f.MustBeAssignable(val)
	*AsInt64Ptr(f.Pointer(structPtr)) = val
}

//SetInt64Ptr sets field *int
func (f *Field) SetInt64Ptr(structPtr unsafe.Pointer, val *int64) {
	f.MustBeAssignable(val)
	*AsInt64AddrPtr(f.Pointer(structPtr)) = val
}

//SetInt32 sets field int
func (f *Field) SetInt32(structPtr unsafe.Pointer, val int32) {
	f.MustBeAssignable(val)
	*AsInt32Ptr(f.Pointer(structPtr)) = val
}

//SetInt32Ptr sets field *int
func (f *Field) SetInt32Ptr(structPtr unsafe.Pointer, val *int32) {
	f.MustBeAssignable(val)
	*AsInt32AddrPtr(f.Pointer(structPtr)) = val
}

//SetInt16 sets field int
func (f *Field) SetInt16(structPtr unsafe.Pointer, val int16) {
	f.MustBeAssignable(val)
	*AsInt16Ptr(f.Pointer(structPtr)) = val
}

//SetInt16Ptr sets field *int
func (f *Field) SetInt16Ptr(structPtr unsafe.Pointer, val *int16) {
	f.MustBeAssignable(val)
	*AsInt16AddrPtr(f.Pointer(structPtr)) = val
}

//SetInt8 sets field int
func (f *Field) SetInt8(structPtr unsafe.Pointer, val int8) {
	f.MustBeAssignable(val)
	*AsInt8Ptr(f.Pointer(structPtr)) = val
}

//SetInt8Ptr sets field *int
func (f *Field) SetInt8Ptr(structPtr unsafe.Pointer, val *int8) {
	f.MustBeAssignable(val)
	*AsInt8AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint sets field uint
func (f *Field) SetUint(structPtr unsafe.Pointer, val uint) {
	f.MustBeAssignable(val)
	*AsUintPtr(f.Pointer(structPtr)) = val
}

//SetUintPtr sets field *uint
func (f *Field) SetUintPtr(structPtr unsafe.Pointer, val *uint) {
	f.MustBeAssignable(val)
	*AsUintAddrPtr(f.Pointer(structPtr)) = val
}

//SetUint64 sets field uint
func (f *Field) SetUint64(structPtr unsafe.Pointer, val uint64) {
	f.MustBeAssignable(val)
	*AsUint64Ptr(f.Pointer(structPtr)) = val
}

//SetUint64Ptr sets field *uint
func (f *Field) SetUint64Ptr(structPtr unsafe.Pointer, val *uint64) {
	f.MustBeAssignable(val)
	*AsUint64AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint32 sets field uint
func (f *Field) SetUint32(structPtr unsafe.Pointer, val uint32) {
	f.MustBeAssignable(val)
	*AsUint32Ptr(f.Pointer(structPtr)) = val
}

//SetUint32Ptr sets field *uint
func (f *Field) SetUint32Ptr(structPtr unsafe.Pointer, val *uint32) {
	f.MustBeAssignable(val)
	*AsUint32AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint16 sets field uint
func (f *Field) SetUint16(structPtr unsafe.Pointer, val uint16) {
	f.MustBeAssignable(val)
	*AsUint16Ptr(f.Pointer(structPtr)) = val
}

//SetUint16Ptr sets field *uint
func (f *Field) SetUint16Ptr(structPtr unsafe.Pointer, val *uint16) {
	f.MustBeAssignable(val)
	*AsUint16AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint8 sets field uint
func (f *Field) SetUint8(structPtr unsafe.Pointer, val uint8) {
	f.MustBeAssignable(val)
	*AsUint8Ptr(f.Pointer(structPtr)) = val
}

//SetUint8Ptr sets field *uint
func (f *Field) SetUint8Ptr(structPtr unsafe.Pointer, val *uint8) {
	f.MustBeAssignable(val)
	*AsUint8AddrPtr(f.Pointer(structPtr)) = val
}

//SetFloat64 sets field float64
func (f *Field) SetFloat64(structPtr unsafe.Pointer, val float64) {
	f.MustBeAssignable(val)
	*AsFloat64Ptr(f.Pointer(structPtr)) = val
}

//SetFloat64Ptr sets field *float64
func (f *Field) SetFloat64Ptr(structPtr unsafe.Pointer, val *float64) {
	f.MustBeAssignable(val)
	*AsFloat64AddrPtr(f.Pointer(structPtr)) = val
}

//SetFloat32 sets field float32
func (f *Field) SetFloat32(structPtr unsafe.Pointer, val float32) {
	f.MustBeAssignable(val)
	*AsFloat32Ptr(f.Pointer(structPtr)) = val
}

//SetFloat32Ptr sets field *float32
func (f *Field) SetFloat32Ptr(structPtr unsafe.Pointer, val *float32) {
	f.MustBeAssignable(val)
	*AsFloat32AddrPtr(f.Pointer(structPtr)) = val
}

//SetBool sets field bool
func (f *Field) SetBool(structPtr unsafe.Pointer, val bool) {
	f.MustBeAssignable(val)
	*AsBoolPtr(f.Pointer(structPtr)) = val
}

//SetBoolPtr sets field *bool
func (f *Field) SetBoolPtr(structPtr unsafe.Pointer, val *bool) {
	f.MustBeAssignable(val)
	*AsBoolAddrPtr(f.Pointer(structPtr)) = val
}

//SetString sets field string
func (f *Field) SetString(structPtr unsafe.Pointer, val string) {
	f.MustBeAssignable(val)
	*AsStringPtr(f.Pointer(structPtr)) = val
}

//SetStringPtr sets field *string
func (f *Field) SetStringPtr(structPtr unsafe.Pointer, val *string) {
	f.MustBeAssignable(val)
	*AsStringAddrPtr(f.Pointer(structPtr)) = val
}

//SetBytes sets field []byte
func (f *Field) SetBytes(structPtr unsafe.Pointer, val []byte) {
	f.MustBeAssignable(val)
	*AsBytesPtr(f.Pointer(structPtr)) = val
}

//SetBytesPtr sets field *[]byte
func (f *Field) SetBytesPtr(structPtr unsafe.Pointer, val *[]byte) {
	f.MustBeAssignable(val)
	*AsBytesAddrPtr(f.Pointer(structPtr)) = val
}

//SetTime sets field time.Time
func (f *Field) SetTime(structPtr unsafe.Pointer, val time.Time) {
	f.MustBeAssignable(val)
	*AsTimePtr(f.Pointer(structPtr)) = val
}

//SetTimePtr sets field *time.Time
func (f *Field) SetTimePtr(structPtr unsafe.Pointer, val *time.Time) {
	f.MustBeAssignable(val)
	*AsTimeAddrPtr(f.Pointer(structPtr)) = val
}

//SetInterface set field interface{}, only support empty interfaces
func (f *Field) SetInterface(structPtr unsafe.Pointer, val interface{}) {
	f.MustBeAssignable(val)
	*(*interface{})(f.Pointer(structPtr)) = val
}

//SetValue sets value
//go:nocheckptr
func (f *Field) SetValue(structPtr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	ptr := f.Pointer(structPtr)

	switch f.kind {
	case reflect.String:
		f.SetStringAt(ptr, source)
	case reflect.Int:
		f.SetIntAt(ptr, source)
	case reflect.Float64:
		f.SetFloat64At(ptr, source)
	case reflect.Float32:
		f.SetFloat32At(ptr, source)
	case reflect.Bool:
		f.SetBoolAt(ptr, source)
	case reflect.Ptr:
		*(*unsafe.Pointer)(ptr) = AsPointer(source)
	case reflect.Func:
		addr := f.Addr(ptr)
		reflect.ValueOf(addr).Elem().Set(reflect.ValueOf(source))
	case reflect.Interface:
		newAt := reflect.NewAt(f.Type, ptr)
		value := reflect.ValueOf(source)
		newAt.Elem().Set(value)
	case reflect.Struct:

		srcPtr := AsPointer(source)
		destPtr := f.Pointer(structPtr)
		Copy(destPtr, srcPtr, int(f.Type.Size()))

		//uncomment this to troubleshoot caller issue
		//newAt := reflect.NewAt(f.Type, ptr)
		//value := reflect.ValueOf(source)
		//newAt.Elem().Set(value)

		//for i := 0; i < f.Type.NumField(); i++ {
		//	field := f.Type.Field(i)
		//	fSrcPtr := unsafe.Pointer(uintptr(srcPtr) + field.Offset)
		//	fDstPtr := unsafe.Pointer(uintptr(destPtr) + field.Offset)
		//	switch field.Type.Kind() {
		//	case reflect.String:
		//		*(*string)(fDstPtr) = *(*string)(fSrcPtr)
		//	case reflect.Ptr, reflect.Struct, reflect.Slice:
		//		srcValue := reflect.NewAt(f.Type, fSrcPtr)
		//		NewField(field).SetValue(fDstPtr, srcValue)
		//	default:
		//		*(*unsafe.Pointer)(fDstPtr) = *(*unsafe.Pointer)(fSrcPtr)
		//	}
	//	}
	case reflect.Slice:
		sourceHeader := (*reflect.SliceHeader)(AsPointer(source))
		destHader := (*reflect.SliceHeader)(ptr)
		destHader.Data = sourceHeader.Data
		destHader.Len = sourceHeader.Len
		destHader.Cap = sourceHeader.Cap
	case reflect.Map:
		newAt := reflect.NewAt(f.Type, ptr)
		value := reflect.ValueOf(source)
		newAt.Elem().Set(value)

	default:
		*(*unsafe.Pointer)(ptr) = *(*unsafe.Pointer)(AsPointer(source))
	}
}

func (f *Field) SetIface(structPtr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	ptr := f.Pointer(structPtr)
	f.SetInterfaceAt(ptr, source)
}

func (f *Field) SetInterfaceAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	newAt := reflect.NewAt(f.Type, ptr)
	value := reflect.ValueOf(source)
	newAt.Elem().Set(value)
}

func (f *Field) SetFunc(structPtr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	ptr := f.Pointer(structPtr)
	f.SetFuncAt(ptr, source)
}

func (f *Field) SetFuncAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	addr := f.Addr(ptr)
	reflect.ValueOf(addr).Elem().Set(reflect.ValueOf(source))
}

func (f *Field) SetStruct(structPtr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	ptr := f.Pointer(structPtr)
	f.SetStructAt(ptr, source)
}

func (f *Field) SetStructAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	srcPtr := AsPointer(source)
	destPtr := ptr
	Copy(destPtr, srcPtr, int(f.Type.Size()))
}

func (f *Field) SetSlice(structPtr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	ptr := f.Pointer(structPtr)
	f.SetSliceAt(ptr, source)
}

func (f *Field) SetSliceAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	sourceHeader := (*reflect.SliceHeader)(AsPointer(source))
	destHader := (*reflect.SliceHeader)(ptr)
	destHader.Data = sourceHeader.Data
	destHader.Len = sourceHeader.Len
	destHader.Cap = sourceHeader.Cap
}

func (f *Field) SetMap(structPtr unsafe.Pointer, aMap interface{}) {
	f.MustBeAssignable(aMap)
	ptr := f.Pointer(structPtr)
	f.SetMapAt(ptr, aMap)
}

func (f *Field) SetMapAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	newAt := reflect.NewAt(f.Type, ptr)
	value := reflect.ValueOf(source)
	newAt.Elem().Set(value)
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

func (f *Field) SetBoolAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	*(*bool)(ptr) = *(*bool)((*emptyInterface)(unsafe.Pointer(&source)).word)
}

func (f *Field) SetFloat32At(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	*(*float32)(ptr) = *(*float32)((*emptyInterface)(unsafe.Pointer(&source)).word)
}

func (f *Field) SetFloat64At(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	*(*float64)(ptr) = *(*float64)((*emptyInterface)(unsafe.Pointer(&source)).word)
}

func (f *Field) SetIntAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	*(*int)(ptr) = *(*int)((*emptyInterface)(unsafe.Pointer(&source)).word)
}

func (f *Field) SetStringAt(ptr unsafe.Pointer, source interface{}) {
	f.MustBeAssignable(source)
	*(*string)(ptr) = *(*string)((*emptyInterface)(unsafe.Pointer(&source)).word)
}

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
