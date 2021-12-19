package xunsafe

import "unsafe"

//PtrSliceUint returns uint slice pointer
func (f *Field) PtrSliceUint(structAddr unsafe.Pointer) *[]uint {
	offset := f.field.Offset
	result := (**[]uint)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceInt64 returns int64 slice pointer
func (f *Field) PtrSliceInt64(structAddr unsafe.Pointer) *[]int64 {
	offset := f.field.Offset
	result := (**[]int64)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceInt32 returns int32 slice pointer
func (f *Field) PtrSliceInt32(structAddr unsafe.Pointer) *[]int32 {
	offset := f.field.Offset
	result := (**[]int32)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceInt16 returns int16 slice pointer
func (f *Field) PtrSliceInt16(structAddr unsafe.Pointer) *[]int16 {
	offset := f.field.Offset
	result := (**[]int16)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceInt8 returns int8 slice pointer
func (f *Field) PtrSliceInt8(structAddr unsafe.Pointer) *[]int8 {
	offset := f.field.Offset
	result := (**[]int8)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceUInt64 returns uint64 slice pointer
func (f *Field) PtrSliceUInt64(structAddr unsafe.Pointer) *[]uint64 {
	offset := f.field.Offset
	result := (**[]uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceUInt32 returns uint32 slice pointer
func (f *Field) PtrSliceUInt32(structAddr unsafe.Pointer) *[]uint32 {
	offset := f.field.Offset
	result := (**[]uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceUInt16 returns uint16 slice pointer
func (f *Field) PtrSliceUInt16(structAddr unsafe.Pointer) *[]uint16 {
	offset := f.field.Offset
	result := (**[]uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceUInt8 returns uint8 slice pointer
func (f *Field) PtrSliceUInt8(structAddr unsafe.Pointer) *[]uint8 {
	offset := f.field.Offset
	result := (**[]uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceString returns string slice pointer
func (f *Field) PtrSliceString(structAddr unsafe.Pointer) *[]string {
	offset := f.field.Offset
	result := (**[]string)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceFloat64 returns float64 slice pointer
func (f *Field) PtrSliceFloat64(structAddr unsafe.Pointer) *[]float64 {
	offset := f.field.Offset
	result := (**[]float64)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceFloat32 returns float32 slice pointer
func (f *Field) PtrSliceFloat32(structAddr unsafe.Pointer) *[]float32 {
	offset := f.field.Offset
	result := (**[]float32)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}

//PtrSliceBool returns bool slice pointer
func (f *Field) PtrSliceBool(structAddr unsafe.Pointer) *[]bool {
	offset := f.field.Offset
	result := (**[]bool)(unsafe.Pointer(uintptr(structAddr) + offset))
	if result == nil {
		return nil
	}
	return *result
}
