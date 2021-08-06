package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//SetInt sets field int
func (f *Field) SetInt(structAddr uintptr, val int) {
	result := (*int)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetIntPtr sets field *int
func (f *Field) SetIntPtr(structAddr uintptr, val *int) {
	result := (**int)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt64 sets field int
func (f *Field) SetInt64(structAddr uintptr, val int64) {
	result := (*int64)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt64Ptr sets field *int
func (f *Field) SetInt64Ptr(structAddr uintptr, val *int64) {
	result := (**int64)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt32 sets field int
func (f *Field) SetInt32(structAddr uintptr, val int32) {
	result := (*int32)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt32Ptr sets field *int
func (f *Field) SetInt32Ptr(structAddr uintptr, val *int32) {
	result := (**int32)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt16 sets field int
func (f *Field) SetInt16(structAddr uintptr, val int16) {
	result := (*int16)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt16Ptr sets field *int
func (f *Field) SetInt16Ptr(structAddr uintptr, val *int16) {
	result := (**int16)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt8 sets field int
func (f *Field) SetInt8(structAddr uintptr, val int8) {
	result := (*int8)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInt8Ptr sets field *int
func (f *Field) SetInt8Ptr(structAddr uintptr, val *int8) {
	result := (**int8)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint sets field uint
func (f *Field) SetUint(structAddr uintptr, val uint) {
	result := (*uint)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUintPtr sets field *uint
func (f *Field) SetUintPtr(structAddr uintptr, val *uint) {
	result := (**uint)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint64 sets field uint
func (f *Field) SetUint64(structAddr uintptr, val uint64) {
	result := (*uint64)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint64Ptr sets field *uint
func (f *Field) SetUint64Ptr(structAddr uintptr, val *uint64) {
	result := (**uint64)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint32 sets field uint
func (f *Field) SetUint32(structAddr uintptr, val uint32) {
	result := (*uint32)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint32Ptr sets field *uint
func (f *Field) SetUint32Ptr(structAddr uintptr, val *uint32) {
	result := (**uint32)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint16 sets field uint
func (f *Field) SetUint16(structAddr uintptr, val uint16) {
	result := (*uint16)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint16Ptr sets field *uint
func (f *Field) SetUint16Ptr(structAddr uintptr, val *uint16) {
	result := (**uint16)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint8 sets field uint
func (f *Field) SetUint8(structAddr uintptr, val uint8) {
	result := (*uint8)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetUint8Ptr sets field *uint
func (f *Field) SetUint8Ptr(structAddr uintptr, val *uint8) {
	result := (**uint8)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetFloat64 sets field float64
func (f *Field) SetFloat64(structAddr uintptr, val float64) {
	result := (*float64)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetFloat64Ptr sets field *float64
func (f *Field) SetFloat64Ptr(structAddr uintptr, val *float64) {
	result := (**float64)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetFloat32 sets field float32
func (f *Field) SetFloat32(structAddr uintptr, val float32) {
	result := (*float32)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetFloat32Ptr sets field *float32
func (f *Field) SetFloat32Ptr(structAddr uintptr, val *float32) {
	result := (**float32)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetBool sets field bool
func (f *Field) SetBool(structAddr uintptr, val bool) {
	result := (*bool)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetBoolPtr sets field *bool
func (f *Field) SetBoolPtr(structAddr uintptr, val *bool) {
	result := (**bool)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetString sets field string
func (f *Field) SetString(structAddr uintptr, val string) {
	result := (*string)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetStringPtr sets field *string
func (f *Field) SetStringPtr(structAddr uintptr, val *string) {
	result := (**string)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetBytes sets field []byte
func (f *Field) SetBytes(structAddr uintptr, val []byte) {
	result := (*[]byte)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetBytesPtr sets field *[]byte
func (f *Field) SetBytesPtr(structAddr uintptr, val *[]byte) {
	result := (**[]byte)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetTime sets field time.Time
func (f *Field) SetTime(structAddr uintptr, val time.Time) {
	result := (*time.Time)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetTimePtr sets field *time.Time
func (f *Field) SetTimePtr(structAddr uintptr, val *time.Time) {
	result := (**time.Time)(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetInterface set field interface{}
func (f *Field) SetInterface(structAddr uintptr, val interface{}) {
	result := (*interface{})(unsafe.Pointer(structAddr + f.field.Offset))
	*result = val
}

//SetValue sets value
func (f *Field) SetValue(structAddr uintptr, val interface{}) {
	reflect.ValueOf(f.Addr(structAddr)).Elem().Set(reflect.ValueOf(val))
}
