package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//AsInt casts pointer to int
func AsInt(pointer unsafe.Pointer) int {
	return *(*int)(pointer)
}

//AsIntPtr casts pointer to *int
func AsIntPtr(pointer unsafe.Pointer) *int {
	return *(**int)(pointer)
}

//AsUInt casts pointer to uint
func AsUInt(pointer unsafe.Pointer) uint {
	return *(*uint)(pointer)
}

//AsUIntPtr casts pointer to *uint
func AsUIntPtr(pointer unsafe.Pointer) *uint {
	return *(**uint)(pointer)
}

//AsInt64 casts pointer to int64
func AsInt64(pointer unsafe.Pointer) int64 {
	return *(*int64)(pointer)
}

//AsInt64Ptr casts pointer to *int64
func AsInt64Ptr(pointer unsafe.Pointer) *int64 {
	return *(**int64)(pointer)
}

//AsUInt64 casts pointer to uint64
func AsUInt64(pointer unsafe.Pointer) uint64 {
	return *(*uint64)(pointer)
}

//AsUInt64Ptr casts pointer to *uint
func AsUInt64Ptr(pointer unsafe.Pointer) *uint64 {
	return *(**uint64)(pointer)
}

//AsInt32 casts pointer to int32
func AsInt32(pointer unsafe.Pointer) int32 {
	return *(*int32)(pointer)
}

//AsInt32Ptr casts pointer to *int32
func AsInt32Ptr(pointer unsafe.Pointer) *int32 {
	return *(**int32)(pointer)
}

//AsUInt32 casts pointer to uint32
func AsUInt32(pointer unsafe.Pointer) uint32 {
	return *(*uint32)(pointer)
}

//AsUInt32Ptr casts pointer to *uint
func AsUInt32Ptr(pointer unsafe.Pointer) *uint32 {
	return *(**uint32)(pointer)
}

//AsInt16 casts pointer to int16
func AsInt16(pointer unsafe.Pointer) int16 {
	return *(*int16)(pointer)
}

//AsInt16Ptr casts pointer to *int16
func AsInt16Ptr(pointer unsafe.Pointer) *int16 {
	return *(**int16)(pointer)
}

//AsUInt16 casts pointer to uint16
func AsUInt16(pointer unsafe.Pointer) uint16 {
	return *(*uint16)(pointer)
}

//AsUInt16Ptr casts pointer to *uint
func AsUInt16Ptr(pointer unsafe.Pointer) *uint16 {
	return *(**uint16)(pointer)
}

//AsInt8 casts pointer to int8
func AsInt8(pointer unsafe.Pointer) int8 {
	return *(*int8)(pointer)
}

//AsInt8Ptr casts pointer to *int8
func AsInt8Ptr(pointer unsafe.Pointer) *int8 {
	return *(**int8)(pointer)
}

//AsUInt8 casts pointer to uint8
func AsUInt8(pointer unsafe.Pointer) uint8 {
	return *(*uint8)(pointer)
}

//AsUInt8Ptr casts pointer to *uint
func AsUInt8Ptr(pointer unsafe.Pointer) *uint8 {
	return *(**uint8)(pointer)
}

//AsFloat32 casts pointer to int
func AsFloat32(pointer unsafe.Pointer) float32 {
	return *(*float32)(pointer)
}

//AsFloat32Ptr casts pointer to *int
func AsFloat32Ptr(pointer unsafe.Pointer) *float32 {
	return *(**float32)(pointer)
}

//AsFloat64 casts pointer to int
func AsFloat64(pointer unsafe.Pointer) float64 {
	return *(*float64)(pointer)
}

//AsFloat64Ptr casts pointer to *int
func AsFloat64Ptr(pointer unsafe.Pointer) *float64 {
	return *(**float64)(pointer)
}

//AsBool casts pointer to int
func AsBool(pointer unsafe.Pointer) bool {
	return *(*bool)(pointer)
}

//AsBoolPtr casts pointer to *int
func AsBoolPtr(pointer unsafe.Pointer) *bool {
	return *(**bool)(pointer)
}

//AsInts casts pointer to int
func AsInts(pointer unsafe.Pointer) []int {
	return *(*[]int)(pointer)
}

//AsIntPtrs casts pointer to *int slice
func AsIntPtrs(pointer unsafe.Pointer) []*int {
	return *(*[]*int)(pointer)
}

//AsUInts casts pointer to uint slice
func AsUInts(pointer unsafe.Pointer) []uint {
	return *(*[]uint)(pointer)
}

//AsUIntPtrs casts pointer to *uint slice
func AsUIntPtrs(pointer unsafe.Pointer) []*uint {
	return *(*[]*uint)(pointer)
}

//AsInt64s casts pointer to int64 slice
func AsInt64s(pointer unsafe.Pointer) []int64 {
	return *(*[]int64)(pointer)
}

//AsInt64Ptrs casts pointer to *int64 slice
func AsInt64Ptrs(pointer unsafe.Pointer) []*int64 {
	return *(*[]*int64)(pointer)
}

//AsUInt64s casts pointer to uint64 slice
func AsUInt64s(pointer unsafe.Pointer) []uint64 {
	return *(*[]uint64)(pointer)
}

//AsUInt64Ptrs casts pointer to *uint slice
func AsUInt64Ptrs(pointer unsafe.Pointer) []*uint64 {
	return *(*[]*uint64)(pointer)
}

//AsInt32s casts pointer to int32 slice
func AsInt32s(pointer unsafe.Pointer) []int32 {
	return *(*[]int32)(pointer)
}

//AsInt32Ptrs casts pointer to *int32 slice
func AsInt32Ptrs(pointer unsafe.Pointer) []*int32 {
	return *(*[]*int32)(pointer)
}

//AsUInt32s casts pointer to uint32 slice
func AsUInt32s(pointer unsafe.Pointer) []uint32 {
	return *(*[]uint32)(pointer)
}

//AsUInt32Ptrs casts pointer to *uint slice
func AsUInt32Ptrs(pointer unsafe.Pointer) []*uint32 {
	return *(*[]*uint32)(pointer)
}

//AsInt16s casts pointer to int16 slice
func AsInt16s(pointer unsafe.Pointer) []int16 {
	return *(*[]int16)(pointer)
}

//AsInt16Ptrs casts pointer to *int16 slice
func AsInt16Ptrs(pointer unsafe.Pointer) []*int16 {
	return *(*[]*int16)(pointer)
}

//AsUInt16s casts pointer to uint16 slice
func AsUInt16s(pointer unsafe.Pointer) []uint16 {
	return *(*[]uint16)(pointer)
}

//AsUInt16Ptrs casts pointer to *uint slice
func AsUInt16Ptrs(pointer unsafe.Pointer) []*uint16 {
	return *(*[]*uint16)(pointer)
}

//AsInt8s casts pointer to int8 slice
func AsInt8s(pointer unsafe.Pointer) []int8 {
	return *(*[]int8)(pointer)
}

//AsInt8Ptrs casts pointer to *int8 slice
func AsInt8Ptrs(pointer unsafe.Pointer) []*int8 {
	return *(*[]*int8)(pointer)
}

//AsUInt8s casts pointer to uint8 slice
func AsUInt8s(pointer unsafe.Pointer) []uint8 {
	return *(*[]uint8)(pointer)
}

//AsUInt8Ptrs casts pointer to *uint slice
func AsUInt8Ptrs(pointer unsafe.Pointer) []*uint8 {
	return *(*[]*uint8)(pointer)
}

//AsFloat32s casts pointer to float32 slice
func AsFloat32s(pointer unsafe.Pointer) []float32 {
	return *(*[]float32)(pointer)
}

//AsFloat32Ptrs casts pointer to *float32 slice
func AsFloat32Ptrs(pointer unsafe.Pointer) []*float32 {
	return *(*[]*float32)(pointer)
}

//AsFloat64s casts pointer to float64 slice
func AsFloat64s(pointer unsafe.Pointer) []float64 {
	return *(*[]float64)(pointer)
}

//AsFloat64Ptrs casts pointer to *float64 slice
func AsFloat64Ptrs(pointer unsafe.Pointer) []*float64 {
	return *(*[]*float64)(pointer)
}

//AsBools casts pointer to bool slice
func AsBools(pointer unsafe.Pointer) []bool {
	return *(*[]bool)(pointer)
}

//AsBoolPtrs casts pointer to *bool slice
func AsBoolPtrs(pointer unsafe.Pointer) []*bool {
	return *(*[]*bool)(pointer)
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
	return *(**time.Time)(pointer)
}

//AsTyped return function casting pointer to the specified type
func AsTyped(destType reflect.Type) func(pointer unsafe.Pointer) interface{} {
	return func(pointer unsafe.Pointer) interface{} {
		//using reflect.NewAt seems to be way slower than just using a reflect.New
		ptr := reflect.New(destType)
		newPointer := unsafe.Pointer(ptr.Pointer())
		*(*unsafe.Pointer)(newPointer) = *(*unsafe.Pointer)(pointer)
		return ptr.Interface()
	}
}
