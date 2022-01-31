package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//SetError sets field error
func (f *Field) SetError(structPtr unsafe.Pointer, val error) {
	*AsErrorPtr(f.Pointer(structPtr)) = val
}

//SetInt sets field int
func (f *Field) SetInt(structPtr unsafe.Pointer, val int) {
	*AsIntPtr(f.Pointer(structPtr)) = val
}

//SetIntPtr sets field *int
func (f *Field) SetIntPtr(structPtr unsafe.Pointer, val *int) {
	*AsIntAddrPtr(f.Pointer(structPtr)) = val
}

//SetInt64 sets field int
func (f *Field) SetInt64(structPtr unsafe.Pointer, val int64) {
	*AsInt64Ptr(f.Pointer(structPtr)) = val
}

//SetInt64Ptr sets field *int
func (f *Field) SetInt64Ptr(structPtr unsafe.Pointer, val *int64) {
	*AsInt64AddrPtr(f.Pointer(structPtr)) = val
}

//SetInt32 sets field int
func (f *Field) SetInt32(structPtr unsafe.Pointer, val int32) {
	*AsInt32Ptr(f.Pointer(structPtr)) = val
}

//SetInt32Ptr sets field *int
func (f *Field) SetInt32Ptr(structPtr unsafe.Pointer, val *int32) {
	*AsInt32AddrPtr(f.Pointer(structPtr)) = val
}

//SetInt16 sets field int
func (f *Field) SetInt16(structPtr unsafe.Pointer, val int16) {
	*AsInt16Ptr(f.Pointer(structPtr)) = val
}

//SetInt16Ptr sets field *int
func (f *Field) SetInt16Ptr(structPtr unsafe.Pointer, val *int16) {
	*AsInt16AddrPtr(f.Pointer(structPtr)) = val
}

//SetInt8 sets field int
func (f *Field) SetInt8(structPtr unsafe.Pointer, val int8) {
	*AsInt8Ptr(f.Pointer(structPtr)) = val
}

//SetInt8Ptr sets field *int
func (f *Field) SetInt8Ptr(structPtr unsafe.Pointer, val *int8) {
	*AsInt8AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint sets field uint
func (f *Field) SetUint(structPtr unsafe.Pointer, val uint) {
	*AsUintPtr(f.Pointer(structPtr)) = val
}

//SetUintPtr sets field *uint
func (f *Field) SetUintPtr(structPtr unsafe.Pointer, val *uint) {
	*AsUintAddrPtr(f.Pointer(structPtr)) = val
}

//SetUint64 sets field uint
func (f *Field) SetUint64(structPtr unsafe.Pointer, val uint64) {
	*AsUint64Ptr(f.Pointer(structPtr)) = val
}

//SetUint64Ptr sets field *uint
func (f *Field) SetUint64Ptr(structPtr unsafe.Pointer, val *uint64) {
	*AsUint64AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint32 sets field uint
func (f *Field) SetUint32(structPtr unsafe.Pointer, val uint32) {
	*AsUint32Ptr(f.Pointer(structPtr)) = val
}

//SetUint32Ptr sets field *uint
func (f *Field) SetUint32Ptr(structPtr unsafe.Pointer, val *uint32) {
	*AsUint32AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint16 sets field uint
func (f *Field) SetUint16(structPtr unsafe.Pointer, val uint16) {
	*AsUint16Ptr(f.Pointer(structPtr)) = val
}

//SetUint16Ptr sets field *uint
func (f *Field) SetUint16Ptr(structPtr unsafe.Pointer, val *uint16) {
	*AsUint16AddrPtr(f.Pointer(structPtr)) = val
}

//SetUint8 sets field uint
func (f *Field) SetUint8(structPtr unsafe.Pointer, val uint8) {
	*AsUint8Ptr(f.Pointer(structPtr)) = val
}

//SetUint8Ptr sets field *uint
func (f *Field) SetUint8Ptr(structPtr unsafe.Pointer, val *uint8) {
	*AsUint8AddrPtr(f.Pointer(structPtr)) = val
}

//SetFloat64 sets field float64
func (f *Field) SetFloat64(structPtr unsafe.Pointer, val float64) {
	*AsFloat64Ptr(f.Pointer(structPtr)) = val
}

//SetFloat64Ptr sets field *float64
func (f *Field) SetFloat64Ptr(structPtr unsafe.Pointer, val *float64) {
	*AsFloat64AddrPtr(f.Pointer(structPtr)) = val
}

//SetFloat32 sets field float32
func (f *Field) SetFloat32(structPtr unsafe.Pointer, val float32) {
	*AsFloat32Ptr(f.Pointer(structPtr)) = val
}

//SetFloat32Ptr sets field *float32
func (f *Field) SetFloat32Ptr(structPtr unsafe.Pointer, val *float32) {
	*AsFloat32AddrPtr(f.Pointer(structPtr)) = val
}

//SetBool sets field bool
func (f *Field) SetBool(structPtr unsafe.Pointer, val bool) {
	*AsBoolPtr(f.Pointer(structPtr)) = val
}

//SetBoolPtr sets field *bool
func (f *Field) SetBoolPtr(structPtr unsafe.Pointer, val *bool) {
	*AsBoolAddrPtr(f.Pointer(structPtr)) = val
}

//SetString sets field string
func (f *Field) SetString(structPtr unsafe.Pointer, val string) {
	*AsStringPtr(f.Pointer(structPtr)) = val
}

//SetStringPtr sets field *string
func (f *Field) SetStringPtr(structPtr unsafe.Pointer, val *string) {
	*AsStringAddrPtr(f.Pointer(structPtr)) = val
}

//SetBytes sets field []byte
func (f *Field) SetBytes(structPtr unsafe.Pointer, val []byte) {
	*AsBytesPtr(f.Pointer(structPtr)) = val
}

//SetBytesPtr sets field *[]byte
func (f *Field) SetBytesPtr(structPtr unsafe.Pointer, val *[]byte) {
	*AsBytesAddrPtr(f.Pointer(structPtr)) = val
}

//SetTime sets field time.Time
func (f *Field) SetTime(structPtr unsafe.Pointer, val time.Time) {
	*AsTimePtr(f.Pointer(structPtr)) = val
}

//SetTimePtr sets field *time.Time
func (f *Field) SetTimePtr(structPtr unsafe.Pointer, val *time.Time) {
	*AsTimeAddrPtr(f.Pointer(structPtr)) = val
}

//SetInterface set field interface{}
func (f *Field) SetInterface(structPtr unsafe.Pointer, val interface{}) {
	*(*interface{})(f.Pointer(structPtr)) = val
}

//SetValue sets value
//go:nocheckptr
func (f *Field) SetValue(structPtr unsafe.Pointer, source interface{}) {
	ptr := f.Pointer(structPtr)
	switch f.kind {
	case reflect.String:
		*(*string)(ptr) = source.(string)
	case reflect.Int:
		*(*int)(ptr) = source.(int)
	case reflect.Float64:
		*(*float64)(ptr) = source.(float64)
	case reflect.Float32:
		*(*float32)(ptr) = source.(float32)
	case reflect.Bool:
		*(*bool)(ptr) = source.(bool)
	case reflect.Ptr:
		*(*unsafe.Pointer)(ptr) = AsPointer(source)
	case reflect.Func:
		addr := f.Addr(ptr)
		reflect.ValueOf(addr).Elem().Set(reflect.ValueOf(source))
	case reflect.Interface:
		newAt := reflect.NewAt(f.Type, ptr)
		newAt.Elem().Set(reflect.ValueOf(source))
	case reflect.Slice:
		sourceHeader := (*reflect.SliceHeader)(AsPointer(source))
		destHader := (*reflect.SliceHeader)(ptr)
		destHader.Data = sourceHeader.Data
		destHader.Len = sourceHeader.Len
		destHader.Cap = sourceHeader.Cap
	default:
		*(*unsafe.Pointer)(ptr) = *(*unsafe.Pointer)(AsPointer(source))
	}
}

//Set sets only non pointer value, the reason for this limited functionality method is speed,
//its 20x faster than SetValue
//go:nocheckptr
func (f *Field) Set(structPtr unsafe.Pointer, source interface{}) {
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
		*(*unsafe.Pointer)(ptr) = *(*unsafe.Pointer)(AsPointer(source))
	}
}
