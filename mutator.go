package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//Setter represents a func setting field value
type Setter func(structAddr unsafe.Pointer, val interface{})

//SetInt sets field int
func (f *Field) SetInt(structAddr unsafe.Pointer, val int) {
	result := (*int)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetIntPtr sets field *int
func (f *Field) SetIntPtr(structAddr unsafe.Pointer, val *int) {
	result := (**int)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt64 sets field int
func (f *Field) SetInt64(structAddr unsafe.Pointer, val int64) {
	result := (*int64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt64Ptr sets field *int
func (f *Field) SetInt64Ptr(structAddr unsafe.Pointer, val *int64) {
	result := (**int64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt32 sets field int
func (f *Field) SetInt32(structAddr unsafe.Pointer, val int32) {
	result := (*int32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt32Ptr sets field *int
func (f *Field) SetInt32Ptr(structAddr unsafe.Pointer, val *int32) {
	result := (**int32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt16 sets field int
func (f *Field) SetInt16(structAddr unsafe.Pointer, val int16) {
	result := (*int16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt16Ptr sets field *int
func (f *Field) SetInt16Ptr(structAddr unsafe.Pointer, val *int16) {
	result := (**int16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt8 sets field int
func (f *Field) SetInt8(structAddr unsafe.Pointer, val int8) {
	result := (*int8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInt8Ptr sets field *int
func (f *Field) SetInt8Ptr(structAddr unsafe.Pointer, val *int8) {
	result := (**int8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint sets field uint
func (f *Field) SetUint(structAddr unsafe.Pointer, val uint) {
	result := (*uint)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUintPtr sets field *uint
func (f *Field) SetUintPtr(structAddr unsafe.Pointer, val *uint) {
	result := (**uint)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint64 sets field uint
func (f *Field) SetUint64(structAddr unsafe.Pointer, val uint64) {
	result := (*uint64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint64Ptr sets field *uint
func (f *Field) SetUint64Ptr(structAddr unsafe.Pointer, val *uint64) {
	result := (**uint64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint32 sets field uint
func (f *Field) SetUint32(structAddr unsafe.Pointer, val uint32) {
	result := (*uint32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint32Ptr sets field *uint
func (f *Field) SetUint32Ptr(structAddr unsafe.Pointer, val *uint32) {
	result := (**uint32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint16 sets field uint
func (f *Field) SetUint16(structAddr unsafe.Pointer, val uint16) {
	result := (*uint16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint16Ptr sets field *uint
func (f *Field) SetUint16Ptr(structAddr unsafe.Pointer, val *uint16) {
	result := (**uint16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint8 sets field uint
func (f *Field) SetUint8(structAddr unsafe.Pointer, val uint8) {
	result := (*uint8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetUint8Ptr sets field *uint
func (f *Field) SetUint8Ptr(structAddr unsafe.Pointer, val *uint8) {
	result := (**uint8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetFloat64 sets field float64
func (f *Field) SetFloat64(structAddr unsafe.Pointer, val float64) {
	result := (*float64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetFloat64Ptr sets field *float64
func (f *Field) SetFloat64Ptr(structAddr unsafe.Pointer, val *float64) {
	result := (**float64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetFloat32 sets field float32
func (f *Field) SetFloat32(structAddr unsafe.Pointer, val float32) {
	result := (*float32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetFloat32Ptr sets field *float32
func (f *Field) SetFloat32Ptr(structAddr unsafe.Pointer, val *float32) {
	result := (**float32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetBool sets field bool
func (f *Field) SetBool(structAddr unsafe.Pointer, val bool) {
	result := (*bool)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetBoolPtr sets field *bool
func (f *Field) SetBoolPtr(structAddr unsafe.Pointer, val *bool) {
	result := (**bool)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetString sets field string
func (f *Field) SetString(structAddr unsafe.Pointer, val string) {
	result := (*string)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetStringPtr sets field *string
func (f *Field) SetStringPtr(structAddr unsafe.Pointer, val *string) {
	result := (**string)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetBytes sets field []byte
func (f *Field) SetBytes(structAddr unsafe.Pointer, val []byte) {
	result := (*[]byte)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetBytesPtr sets field *[]byte
func (f *Field) SetBytesPtr(structAddr unsafe.Pointer, val *[]byte) {
	result := (**[]byte)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetTime sets field time.Time
func (f *Field) SetTime(structAddr unsafe.Pointer, val time.Time) {
	result := (*time.Time)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetTimePtr sets field *time.Time
func (f *Field) SetTimePtr(structAddr unsafe.Pointer, val *time.Time) {
	result := (**time.Time)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetInterface set field interface{}
func (f *Field) SetInterface(structAddr unsafe.Pointer, val interface{}) {
	result := (*interface{})(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	*result = val
}

//SetValue sets value
func (f *Field) SetValue(structAddr unsafe.Pointer, val interface{}) {
	refValue := reflect.ValueOf(val)
	switch refValue.Kind() {
	case reflect.Ptr:
		ptr := (*unsafe.Pointer)(unsafe.Add(structAddr, f.field.Offset))
		if ptr == nil {
			return
		}
		if refValue.IsZero() {
			*ptr = nil
		} else {
			*ptr = unsafe.Pointer(refValue.Elem().UnsafeAddr())
		}
	case reflect.Func:
		addr := f.Addr(structAddr)
		reflect.ValueOf(addr).Elem().Set(refValue)
	default:
		addr := f.Addr(structAddr)
		if addr == nil || !refValue.IsValid() {
			return
		}
		reflect.ValueOf(addr).Elem().Set(refValue)
	}

}
