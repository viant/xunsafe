package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

var (
	typeTime    = reflect.TypeOf(time.Time{})
	typeTimePtr = reflect.TypeOf(&time.Time{})
)

//AsInt casts pointer to int
func AsInt(pointer unsafe.Pointer) int {
	return *(*int)(pointer)
}

//AsIntPtr casts pointer to *int
func AsIntPtr(pointer unsafe.Pointer) *int {
	return (*int)(pointer)
}

//AsIntAddrPtr casts pointer to **int
func AsIntAddrPtr(pointer unsafe.Pointer) **int {
	return (**int)(pointer)
}

//AsUint casts pointer to uint
func AsUint(pointer unsafe.Pointer) uint {
	return *(*uint)(pointer)
}

//AsUintPtr casts pointer to *uint
func AsUintPtr(pointer unsafe.Pointer) *uint {
	return (*uint)(pointer)
}

//AsUintAddrPtr casts pointer to **uint
func AsUintAddrPtr(pointer unsafe.Pointer) **uint {
	return (**uint)(pointer)
}

//AsInt64 casts pointer to int64
func AsInt64(pointer unsafe.Pointer) int64 {
	return *(*int64)(pointer)
}

//AsInt64Ptr casts pointer to *int64
func AsInt64Ptr(pointer unsafe.Pointer) *int64 {
	return (*int64)(pointer)
}

//AsInt64AddrPtr casts pointer to **int64
func AsInt64AddrPtr(pointer unsafe.Pointer) **int64 {
	return (**int64)(pointer)
}

//AsUint64 casts pointer to uint64
func AsUint64(pointer unsafe.Pointer) uint64 {
	return *(*uint64)(pointer)
}

//AsUint64Ptr casts pointer to *uint
func AsUint64Ptr(pointer unsafe.Pointer) *uint64 {
	return (*uint64)(pointer)
}

//AsUint64AddrPtr casts pointer to *uint
func AsUint64AddrPtr(pointer unsafe.Pointer) **uint64 {
	return (**uint64)(pointer)
}

//AsInt32 casts pointer to int32
func AsInt32(pointer unsafe.Pointer) int32 {
	return *(*int32)(pointer)
}

//AsInt32Ptr casts pointer to *int32
func AsInt32Ptr(pointer unsafe.Pointer) *int32 {
	return (*int32)(pointer)
}

//AsInt32AddrPtr casts pointer to **int32
func AsInt32AddrPtr(pointer unsafe.Pointer) **int32 {
	return (**int32)(pointer)
}

//AsUint32 casts pointer to uint32
func AsUint32(pointer unsafe.Pointer) uint32 {
	return *(*uint32)(pointer)
}

//AsUint32Ptr casts pointer to *uint
func AsUint32Ptr(pointer unsafe.Pointer) *uint32 {
	return (*uint32)(pointer)
}

//AsUint32AddrPtr casts pointer to *uint
func AsUint32AddrPtr(pointer unsafe.Pointer) **uint32 {
	return (**uint32)(pointer)
}

//AsInt16 casts pointer to int16
func AsInt16(pointer unsafe.Pointer) int16 {
	return *(*int16)(pointer)
}

//AsInt16Ptr casts pointer to *int16
func AsInt16Ptr(pointer unsafe.Pointer) *int16 {
	return (*int16)(pointer)
}

//AsInt16AddrPtr casts pointer to *int16
func AsInt16AddrPtr(pointer unsafe.Pointer) **int16 {
	return (**int16)(pointer)
}

//AsUint16 casts pointer to uint16
func AsUint16(pointer unsafe.Pointer) uint16 {
	return *(*uint16)(pointer)
}

//AsUintptr casts pointer to uint16
func AsUintptr(pointer unsafe.Pointer) uintptr {
	return *(*uintptr)(pointer)
}

//AsUint16Ptr casts pointer to *uint
func AsUint16Ptr(pointer unsafe.Pointer) *uint16 {
	return (*uint16)(pointer)
}

//AsUint16AddrPtr casts pointer to **uint
func AsUint16AddrPtr(pointer unsafe.Pointer) **uint16 {
	return (**uint16)(pointer)
}

//AsInt8 casts pointer to int8
func AsInt8(pointer unsafe.Pointer) int8 {
	return *(*int8)(pointer)
}

//AsInt8Ptr casts pointer to *int8
func AsInt8Ptr(pointer unsafe.Pointer) *int8 {
	return (*int8)(pointer)
}

//AsInt8AddrPtr casts pointer to *int8
func AsInt8AddrPtr(pointer unsafe.Pointer) **int8 {
	return (**int8)(pointer)
}

//AsUint8 casts pointer to uint8
func AsUint8(pointer unsafe.Pointer) uint8 {
	return *(*uint8)(pointer)
}

//AsUint8Ptr casts pointer to *uint
func AsUint8Ptr(pointer unsafe.Pointer) *uint8 {
	return (*uint8)(pointer)
}

//AsUint8AddrPtr casts pointer to **uint
func AsUint8AddrPtr(pointer unsafe.Pointer) **uint8 {
	return (**uint8)(pointer)
}

//AsFloat32 casts pointer to float32
func AsFloat32(pointer unsafe.Pointer) float32 {
	return *(*float32)(pointer)
}

//AsFloat32Ptr casts pointer to *float32
func AsFloat32Ptr(pointer unsafe.Pointer) *float32 {
	return (*float32)(pointer)
}

//AsFloat32AddrPtr casts pointer to **float32
func AsFloat32AddrPtr(pointer unsafe.Pointer) **float32 {
	return (**float32)(pointer)
}

//AsFloat64 casts pointer to float64
func AsFloat64(pointer unsafe.Pointer) float64 {
	return *(*float64)(pointer)
}

//AsFloat64Ptr casts pointer to *float64
func AsFloat64Ptr(pointer unsafe.Pointer) *float64 {
	return (*float64)(pointer)
}

//AsFloat64AddrPtr casts pointer to **float64
func AsFloat64AddrPtr(pointer unsafe.Pointer) **float64 {
	return (**float64)(pointer)
}

//AsBool casts pointer to bool
func AsBool(pointer unsafe.Pointer) bool {
	return *(*bool)(pointer)
}

//AsBoolPtr casts pointer to *bool
func AsBoolPtr(pointer unsafe.Pointer) *bool {
	return (*bool)(pointer)
}

//AsBoolAddrPtr casts pointer to **bool
func AsBoolAddrPtr(pointer unsafe.Pointer) **bool {
	return (**bool)(pointer)
}

//AsString casts pointer to string
func AsString(pointer unsafe.Pointer) string {
	return *(*string)(pointer)
}

//AsStringPtr casts pointer to *string
func AsStringPtr(pointer unsafe.Pointer) *string {
	return (*string)(pointer)
}

//AsStringAddrPtr casts pointer to **string
func AsStringAddrPtr(pointer unsafe.Pointer) **string {
	return (**string)(pointer)
}

//AsStrings casts pointer to string slice
func AsStrings(pointer unsafe.Pointer) []string {
	return *(*[]string)(pointer)
}

//AsStringsPtr casts pointer to string slice pointer
func AsStringsPtr(pointer unsafe.Pointer) *[]string {
	return (*[]string)(pointer)
}

//AsInts casts pointer to []int
func AsInts(pointer unsafe.Pointer) []int {
	return *(*[]int)(pointer)
}

//AsIntsPtr casts pointer to int slice pointer
func AsIntsPtr(pointer unsafe.Pointer) *[]int {
	return (*[]int)(pointer)
}

//AsUints casts pointer to uint slice
func AsUints(pointer unsafe.Pointer) []uint {
	return *(*[]uint)(pointer)
}

//AsInt64s casts pointer to int64 slice
func AsInt64s(pointer unsafe.Pointer) []int64 {
	return *(*[]int64)(pointer)
}

//AsUint64s casts pointer to uint64 slice
func AsUint64s(pointer unsafe.Pointer) []uint64 {
	return *(*[]uint64)(pointer)
}

//AsInt32s casts pointer to int32 slice
func AsInt32s(pointer unsafe.Pointer) []int32 {
	return *(*[]int32)(pointer)
}

//AsUint32s casts pointer to uint32 slice
func AsUint32s(pointer unsafe.Pointer) []uint32 {
	return *(*[]uint32)(pointer)
}

//AsInt16s casts pointer to int16 slice
func AsInt16s(pointer unsafe.Pointer) []int16 {
	return *(*[]int16)(pointer)
}

//AsUint16s casts pointer to uint16 slice
func AsUint16s(pointer unsafe.Pointer) []uint16 {
	return *(*[]uint16)(pointer)
}

//AsInt8s casts pointer to int8 slice
func AsInt8s(pointer unsafe.Pointer) []int8 {
	return *(*[]int8)(pointer)
}

//AsUint8s casts pointer to uint8 slice
func AsUint8s(pointer unsafe.Pointer) []uint8 {
	return *(*[]uint8)(pointer)
}

//AsBytesPtr casts pointer to []byte  pointer
func AsBytesPtr(pointer unsafe.Pointer) *[]byte {
	return (*[]uint8)(pointer)
}

//AsBytesAddrPtr casts pointer to []byte  pointer
func AsBytesAddrPtr(pointer unsafe.Pointer) **[]byte {
	return (**[]uint8)(pointer)
}

//AsUint8Ptrs casts pointer to uint slice pointer
func AsUint8Ptrs(pointer unsafe.Pointer) *[]uint8 {
	return (*[]uint8)(pointer)
}

//AsFloat32s casts pointer to float32 slice
func AsFloat32s(pointer unsafe.Pointer) []float32 {
	return *(*[]float32)(pointer)
}

//AsFloat32sPtr casts pointer to float32 slice
func AsFloat32sPtr(pointer unsafe.Pointer) *[]float32 {
	return (*[]float32)(pointer)
}

//AsFloat64s casts pointer to float64 slice
func AsFloat64s(pointer unsafe.Pointer) []float64 {
	return *(*[]float64)(pointer)
}

//AsFloat64sPtr casts pointer to float64 slice pointer
func AsFloat64sPtr(pointer unsafe.Pointer) *[]float64 {
	return (*[]float64)(pointer)
}

//AsBools casts pointer to bool slice
func AsBools(pointer unsafe.Pointer) []bool {
	return *(*[]bool)(pointer)
}

//AsInterface casts pointer to interface
func AsInterface(pointer unsafe.Pointer) interface{} {
	return *(*interface{})(pointer)
}

//AsInterfaces casts pointer to interface skuce
func AsInterfaces(pointer unsafe.Pointer) []interface{} {
	return *(*[]interface{})(pointer)
}

//AsStringMap casts pointer to map[string]interface
func AsStringMap(pointer unsafe.Pointer) map[string]interface{} {
	return *(*map[string]interface{})(pointer)
}

//AsMap casts pointer to map[string]interface
func AsMap(pointer unsafe.Pointer) map[interface{}]interface{} {
	return *(*map[interface{}]interface{})(pointer)
}

//AsTime cast pointer to time.Time
func AsTime(pointer unsafe.Pointer) time.Time {
	return *(*time.Time)(pointer)
}

//AsTimePtr cast pointer to *time.Time
func AsTimePtr(pointer unsafe.Pointer) *time.Time {
	return (*time.Time)(pointer)
}

//AsTimeAddrPtr cast pointer to **time.Time
func AsTimeAddrPtr(pointer unsafe.Pointer) **time.Time {
	return (**time.Time)(pointer)
}
