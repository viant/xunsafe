package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//IntAddr returns field *int address
func (f *Field) IntAddr(structAddr unsafe.Pointer) *int {
	return (*int)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Int returns field int
func (f *Field) Int(structAddr unsafe.Pointer) int {
	result := (*int)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//IntPtr returns field *int
func (f *Field) IntPtr(structAddr unsafe.Pointer) *int {
	result := (**int)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int64Addr returns field *int64 addr
func (f *Field) Int64Addr(structAddr unsafe.Pointer) *int64 {
	return (*int64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Int64 returns field int64
func (f *Field) Int64(structAddr unsafe.Pointer) int64 {
	result := (*int64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int64Ptr returns field *int64
func (f *Field) Int64Ptr(structAddr unsafe.Pointer) *int64 {
	result := (**int64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int32Addr returns field *int32 addr
func (f *Field) Int32Addr(structAddr unsafe.Pointer) *int32 {
	return (*int32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Int32 returns field int32
func (f *Field) Int32(structAddr unsafe.Pointer) int32 {
	result := (*int32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int32Ptr returns field *int32
func (f *Field) Int32Ptr(structAddr unsafe.Pointer) *int32 {
	result := (**int32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int16Addr returns field *int16 addr
func (f *Field) Int16Addr(structAddr unsafe.Pointer) *int16 {
	return (*int16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Int16 returns field int16
func (f *Field) Int16(structAddr unsafe.Pointer) int16 {
	result := (*int16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int16Ptr returns field *int16
func (f *Field) Int16Ptr(structAddr unsafe.Pointer) *int16 {
	result := (**int16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Int8Addr returns field *int8 addr
func (f *Field) Int8Addr(structAddr unsafe.Pointer) *int8 {
	return (*int8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Int8 returns field int8
func (f *Field) Int8(structAddr unsafe.Pointer) int8 {
	result := (*int8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Int8Ptr returns field *int8
func (f *Field) Int8Ptr(structAddr unsafe.Pointer) *int8 {
	result := (**int8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//UintAddr returns field *uint address
func (f *Field) UintAddr(structAddr unsafe.Pointer) *uint {
	return (*uint)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Uint returns field uint
func (f *Field) Uint(structAddr unsafe.Pointer) uint {
	result := (*uint)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//UintPtr returns field *uint
func (f *Field) UintPtr(structAddr unsafe.Pointer) *uint {
	result := (**uint)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint64Addr returns field *uint64 addr
func (f *Field) Uint64Addr(structAddr unsafe.Pointer) *uint64 {
	return (*uint64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Uint64 returns field uint64
func (f *Field) Uint64(structAddr unsafe.Pointer) uint64 {
	result := (*uint64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint64Ptr returns field *uint64
func (f *Field) Uint64Ptr(structAddr unsafe.Pointer) *uint64 {
	result := (**uint64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint32Addr returns field *uint32 addr
func (f *Field) Uint32Addr(structAddr unsafe.Pointer) *uint32 {
	return (*uint32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Uint32 returns field uint32
func (f *Field) Uint32(structAddr unsafe.Pointer) uint32 {
	result := (*uint32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint32Ptr returns field *uint32
func (f *Field) Uint32Ptr(structAddr unsafe.Pointer) *uint32 {
	result := (**uint32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint16Addr returns field *uint16 addr
func (f *Field) Uint16Addr(structAddr unsafe.Pointer) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Uint16 returns field uint16
func (f *Field) Uint16(structAddr unsafe.Pointer) uint16 {
	result := (*uint16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint16Ptr returns field *uint16
func (f *Field) Uint16Ptr(structAddr unsafe.Pointer) *uint16 {
	result := (**uint16)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Uint8Addr returns field *uint8 addr
func (f *Field) Uint8Addr(structAddr unsafe.Pointer) *uint8 {
	return (*uint8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Uint8 returns field uint8
func (f *Field) Uint8(structAddr unsafe.Pointer) uint8 {
	result := (*uint8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Uint8Ptr returns field *uint8
func (f *Field) Uint8Ptr(structAddr unsafe.Pointer) *uint8 {
	result := (**uint8)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//BoolAddr returns field *bool addr
func (f *Field) BoolAddr(structAddr unsafe.Pointer) *bool {
	return (*bool)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Bool returns field bool
func (f *Field) Bool(structAddr unsafe.Pointer) bool {
	result := (*bool)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return false
	}
	return *result
}

//BoolPtr returns field *bool
func (f *Field) BoolPtr(structAddr unsafe.Pointer) *bool {
	result := (**bool)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Float64Addr returns field *float64 addr
func (f *Field) Float64Addr(structAddr unsafe.Pointer) *float64 {
	return (*float64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Float64 returns field float64
func (f *Field) Float64(structAddr unsafe.Pointer) float64 {
	result := (*float64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Float64Ptr returns field *float64
func (f *Field) Float64Ptr(structAddr unsafe.Pointer) *float64 {
	result := (**float64)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Float32Addr returns field *float32 addr
func (f *Field) Float32Addr(structAddr unsafe.Pointer) *float32 {
	return (*float32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Float32 returns field float32
func (f *Field) Float32(structAddr unsafe.Pointer) float32 {
	result := (*float32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return 0
	}
	return *result
}

//Float32Ptr returns field *float32
func (f *Field) Float32Ptr(structAddr unsafe.Pointer) *float32 {
	result := (**float32)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//StringAddr returns field *string addr
func (f *Field) StringAddr(structAddr unsafe.Pointer) *string {
	return (*string)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//String returns field string
func (f *Field) String(structAddr unsafe.Pointer) string {
	result := (*string)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return ""
	}
	return *result
}

//StringPtr returns field *string
func (f *Field) StringPtr(structAddr unsafe.Pointer) *string {
	result := (**string)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//BytesAddr returns field *[]byte addr
func (f *Field) BytesAddr(structAddr unsafe.Pointer) *[]byte {
	return (*[]byte)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Bytes returns field []byte
func (f *Field) Bytes(structAddr unsafe.Pointer) []byte {
	result := (*[]byte)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//BytesPtr returns field *[]byte
func (f *Field) BytesPtr(structAddr unsafe.Pointer) *[]byte {
	result := (**[]byte)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//TimeAddr returns field *time.Time addr
func (f *Field) TimeAddr(structAddr unsafe.Pointer) *time.Time {
	return (*time.Time)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
}

//Time returns field time.Time
func (f *Field) Time(structAddr unsafe.Pointer) time.Time {
	result := (*time.Time)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return time.Time{}
	}
	return *result
}

//TimePtr returns field *time.Time
func (f *Field) TimePtr(structAddr unsafe.Pointer) *time.Time {
	result := (**time.Time)(unsafe.Pointer(uintptr(structAddr) + f.field.Offset))
	if result == nil {
		return nil
	}
	return *result
}

//Value returns a value getter or error
func (f *Field) Value(structAddr unsafe.Pointer) interface{} {
	if f.value != nil {
		return f.value(structAddr)
	}
	f.value = f.fieldValue()
	return f.value(structAddr)
}

//Addr returns a field addr getter or error
func (f *Field) Addr(structAddr unsafe.Pointer) interface{} {
	if f.address != nil {
		return f.address(structAddr)
	}
	f.address = f.AddrGetter()
	return f.address(structAddr)
}

//Interface returns field address
func (f *Field) Interface(structAddr unsafe.Pointer) interface{} {
	fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+f.field.Offset))
	return fieldValue.Elem().Elem().Interface()
}
