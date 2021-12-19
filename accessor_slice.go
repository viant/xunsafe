package xunsafe

import (
	"reflect"
	"unsafe"
)

//SliceInt returns int slice
func (f *Field) SliceInt(structAddr unsafe.Pointer) []int {
	offset := f.field.Offset
	result := (*[]int)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceUint returns uint slice
func (f *Field) SliceUint(structAddr unsafe.Pointer) []uint {
	offset := f.field.Offset
	result := (*[]uint)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceInt64 returns int64 slice
func (f *Field) SliceInt64(structAddr unsafe.Pointer) []int64 {
	offset := f.field.Offset
	result := (*[]int64)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceInt32 returns int32 slice
func (f *Field) SliceInt32(structAddr unsafe.Pointer) []int32 {
	offset := f.field.Offset
	result := (*[]int32)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceInt16 returns int16 slice
func (f *Field) SliceInt16(structAddr unsafe.Pointer) []int16 {
	offset := f.field.Offset
	result := (*[]int16)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceInt8 returns int8 slice
func (f *Field) SliceInt8(structAddr unsafe.Pointer) []int8 {
	offset := f.field.Offset
	result := (*[]int8)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceUint64 returns uint64 slice
func (f *Field) SliceUint64(structAddr unsafe.Pointer) []uint64 {
	offset := f.field.Offset
	result := (*[]uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceUint32 returns uint32 slice
func (f *Field) SliceUint32(structAddr unsafe.Pointer) []uint32 {
	offset := f.field.Offset
	result := (*[]uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceUint16 returns uint16 slice
func (f *Field) SliceUint16(structAddr unsafe.Pointer) []uint16 {
	offset := f.field.Offset
	result := (*[]uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceUint8 returns uint8 slice
func (f *Field) SliceUint8(structAddr unsafe.Pointer) []uint8 {
	offset := f.field.Offset
	result := (*[]uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceString returns string slice
func (f *Field) SliceString(structAddr unsafe.Pointer) []string {
	offset := f.field.Offset
	result := (*[]string)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceFloat64 returns float64 slice
func (f *Field) SliceFloat64(structAddr unsafe.Pointer) []float64 {
	offset := f.field.Offset
	result := (*[]float64)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceFloat32 returns float32 slice
func (f *Field) SliceFloat32(structAddr unsafe.Pointer) []float32 {
	offset := f.field.Offset
	result := (*[]float32)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//SliceBool returns bool slice
func (f *Field) SliceBool(structAddr unsafe.Pointer) []bool {
	offset := f.field.Offset
	result := (*[]bool)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//InterfaceSlice returns interface slice
func (f *Field) InterfaceSlice(structAddr unsafe.Pointer) interface{} {
	offset := f.field.Offset
	fieldValue := reflect.NewAt(f.field.Type, unsafe.Add(structAddr, offset))
	return fieldValue.Elem().Interface()
}
