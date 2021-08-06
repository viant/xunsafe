package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//IntAddr returns field *int address
func (f *Field) IntAddr(structAddr uintptr) *int {
	return (*int)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Int returns field int
func (f *Field) Int(structAddr uintptr) int {
	result := (*int)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//IntPtr returns field *int
func (f *Field) IntPtr(structAddr uintptr) *int {
	result := (**int)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int64Addr returns field *int64 addr
func (f *Field) Int64Addr(structAddr uintptr) *int64 {
	return (*int64)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Int64 returns field int64
func (f *Field) Int64(structAddr uintptr) int64 {
	result := (*int64)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int64Ptr returns field *int64
func (f *Field) Int64Ptr(structAddr uintptr) *int64 {
	result := (**int64)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int32Addr returns field *int32 addr
func (f *Field) Int32Addr(structAddr uintptr) *int32 {
	return (*int32)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Int32 returns field int32
func (f *Field) Int32(structAddr uintptr) int32 {
	result := (*int32)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int32Ptr returns field *int32
func (f *Field) Int32Ptr(structAddr uintptr) *int32 {
	result := (**int32)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int16Addr returns field *int16 addr
func (f *Field) Int16Addr(structAddr uintptr) *int16 {
	return (*int16)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Int16 returns field int16
func (f *Field) Int16(structAddr uintptr) int16 {
	result := (*int16)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int16Ptr returns field *int16
func (f *Field) Int16Ptr(structAddr uintptr) *int16 {
	result := (**int16)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int8Addr returns field *int8 addr
func (f *Field) Int8Addr(structAddr uintptr) *int8 {
	return (*int8)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Int8 returns field int8
func (f *Field) Int8(structAddr uintptr) int8 {
	result := (*int8)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int8Ptr returns field *int8
func (f *Field) Int8Ptr(structAddr uintptr) *int8 {
	result := (**int8)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//UintAddr returns field *uint address
func (f *Field) UintAddr(structAddr uintptr) *uint {
	return (*uint)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Uint returns field uint
func (f *Field) Uint(structAddr uintptr) uint {
	result := (*uint)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//UintPtr returns field *uint
func (f *Field) UintPtr(structAddr uintptr) *uint {
	result := (**uint)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint64Addr returns field *uint64 addr
func (f *Field) Uint64Addr(structAddr uintptr) *uint64 {
	return (*uint64)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Uint64 returns field uint64
func (f *Field) Uint64(structAddr uintptr) uint64 {
	result := (*uint64)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint64Ptr returns field *uint64
func (f *Field) Uint64Ptr(structAddr uintptr) *uint64 {
	result := (**uint64)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint32Addr returns field *uint32 addr
func (f *Field) Uint32Addr(structAddr uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Uint32 returns field uint32
func (f *Field) Uint32(structAddr uintptr) uint32 {
	result := (*uint32)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint32Ptr returns field *uint32
func (f *Field) Uint32Ptr(structAddr uintptr) *uint32 {
	result := (**uint32)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint16Addr returns field *uint16 addr
func (f *Field) Uint16Addr(structAddr uintptr) *uint16 {
	return (*uint16)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Uint16 returns field uint16
func (f *Field) Uint16(structAddr uintptr) uint16 {
	result := (*uint16)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint16Ptr returns field *uint16
func (f *Field) Uint16Ptr(structAddr uintptr) *uint16 {
	result := (**uint16)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint8Addr returns field *uint8 addr
func (f *Field) Uint8Addr(structAddr uintptr) *uint8 {
	return (*uint8)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Uint8 returns field uint8
func (f *Field) Uint8(structAddr uintptr) uint8 {
	result := (*uint8)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint8Ptr returns field *uint8
func (f *Field) Uint8Ptr(structAddr uintptr) *uint8 {
	result := (**uint8)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//BoolAddr returns field *bool addr
func (f *Field) BoolAddr(structAddr uintptr) *bool {
	return (*bool)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Bool returns field bool
func (f *Field) Bool(structAddr uintptr) bool {
	result := (*bool)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return false
	}
	return *result
}

//BoolPtr returns field *bool
func (f *Field) BoolPtr(structAddr uintptr) *bool {
	result := (**bool)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Float64Addr returns field *float64 addr
func (f *Field) Float64Addr(structAddr uintptr) *float64 {
	return (*float64)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Float64 returns field float64
func (f *Field) Float64(structAddr uintptr) float64 {
	result := (*float64)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Float64Ptr returns field *float64
func (f *Field) Float64Ptr(structAddr uintptr) *float64 {
	result := (**float64)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Float32Addr returns field *float32 addr
func (f *Field) Float32Addr(structAddr uintptr) *float32 {
	return (*float32)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Float32 returns field float32
func (f *Field) Float32(structAddr uintptr) float32 {
	result := (*float32)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Float32Ptr returns field *float32
func (f *Field) Float32Ptr(structAddr uintptr) *float32 {
	result := (**float32)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//StringAddr returns field *string addr
func (f *Field) StringAddr(structAddr uintptr) *string {
	return (*string)(unsafe.Pointer(structAddr + f.field.Offset))
}

//String returns field string
func (f *Field) String(structAddr uintptr) string {
	result := (*string)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return ""
	}
	return *result
}

//StringPtr returns field *string
func (f *Field) StringPtr(structAddr uintptr) *string {
	result := (**string)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//BytesAddr returns field *[]byte addr
func (f *Field) BytesAddr(structAddr uintptr) *[]byte {
	return (*[]byte)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Bytes returns field []byte
func (f *Field) Bytes(structAddr uintptr) []byte {
	result := (*[]byte)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//BytesPtr returns field *[]byte
func (f *Field) BytesPtr(structAddr uintptr) *[]byte {
	result := (**[]byte)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//TimeAddr returns field *time.Time addr
func (f *Field) TimeAddr(structAddr uintptr) *time.Time {
	return (*time.Time)(unsafe.Pointer(structAddr + f.field.Offset))
}

//Time returns field time.Time
func (f *Field) Time(structAddr uintptr) time.Time {
	result := (*time.Time)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return time.Time{}
	}
	return *result
}

//TimePtr returns field *time.Time
func (f *Field) TimePtr(structAddr uintptr) *time.Time {
	result := (**time.Time)(unsafe.Pointer(structAddr + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Value returns a value getter or error
func (f *Field) Value(structAddr uintptr) interface{} {
	if f.Val != nil {
		return f.Val(structAddr)
	}
	f.Val = f.fieldValue()
	return f.Val(structAddr)
}

//Addr returns a field addr getter or error
func (f *Field) Addr(structAddr uintptr) interface{} {
	if f.Address != nil {
		return f.Address(structAddr)
	}
	f.Address = f.AddrGetter()
	return f.Address(structAddr)
}

//Interface returns field address
func (f *Field) Interface(structAddr uintptr) interface{} {
	fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+f.field.Offset))
	return fieldValue.Elem().Elem().Interface()
}
