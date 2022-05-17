package xunsafe

import (
	"time"
	"unsafe"
)

//Interface cast field pointer to value
func (f *Field) Interface(structPtr unsafe.Pointer) interface{} {
	ptr := f.Pointer(structPtr)
	if f.iface {
		return *(*interface{})(ptr)
	}
	//return reflect.NewAt(f.Type, f.Ref(structPtr)).Elem().Interface()
	return asInterface(ptr, f.rtype, true)
}

//Int cast field pointer to int
func (f *Field) Int(structPtr unsafe.Pointer) int {
	return AsInt(f.Pointer(structPtr))
}

//IntPtr cast field pointer to *int
func (f *Field) IntPtr(structPtr unsafe.Pointer) *int {
	result := AsIntAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Error cast field pointer to error
func (f *Field) Error(structPtr unsafe.Pointer) error {
	return AsError(f.Pointer(structPtr))
}

//ErrorPtr cast field pointer to *error
func (f *Field) ErrorPtr(structPtr unsafe.Pointer) *error {
	result := AsErrorAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//IntAddr cast field pointer to **int
func (f *Field) IntAddr(structPtr unsafe.Pointer) *int {
	return AsIntPtr(f.Pointer(structPtr))
}

//Int64 cast field pointer to int64
func (f *Field) Int64(structPtr unsafe.Pointer) int64 {
	return AsInt64(f.Pointer(structPtr))
}

//Int64Ptr cast field pointer to *int
func (f *Field) Int64Ptr(structPtr unsafe.Pointer) *int64 {
	result := AsInt64AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int64Addr cast field pointer to *int64
func (f *Field) Int64Addr(structPtr unsafe.Pointer) *int64 {
	return AsInt64Ptr(f.Pointer(structPtr))
}

//Int32 cast field pointer to int32
func (f *Field) Int32(structPtr unsafe.Pointer) int32 {
	return AsInt32(f.Pointer(structPtr))
}

//Int32Ptr cast field pointer to *int32
func (f *Field) Int32Ptr(structPtr unsafe.Pointer) *int32 {
	result := AsInt32AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int32Addr cast field pointer to *int32
func (f *Field) Int32Addr(structPtr unsafe.Pointer) *int32 {
	return AsInt32Ptr(f.Pointer(structPtr))
}

//Int16 cast field pointer to int16
func (f *Field) Int16(structPtr unsafe.Pointer) int16 {
	return AsInt16(f.Pointer(structPtr))
}

//Int16Ptr cast field pointer to *int16
func (f *Field) Int16Ptr(structPtr unsafe.Pointer) *int16 {
	result := AsInt16AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int16Addr returns *int16
func (f *Field) Int16Addr(structPtr unsafe.Pointer) *int16 {
	return AsInt16Ptr(f.Pointer(structPtr))
}

//Int8 cast field pointer to int8
func (f *Field) Int8(structPtr unsafe.Pointer) int8 {
	return AsInt8(f.Pointer(structPtr))
}

//Int8Ptr cast field pointer to *int8
func (f *Field) Int8Ptr(structPtr unsafe.Pointer) *int8 {
	result := AsInt8AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int8Addr cast field pointer to *int8
func (f *Field) Int8Addr(structPtr unsafe.Pointer) *int8 {
	return AsInt8Ptr(f.Pointer(structPtr))
}

//UintAddr cast field pointer to **uint
func (f *Field) UintAddr(structPtr unsafe.Pointer) *uint {
	return AsUintPtr(f.Pointer(structPtr))
}

//Uint cast field pointer to uint
func (f *Field) Uint(structPtr unsafe.Pointer) uint {
	return AsUint(f.Pointer(structPtr))
}

//UintPtr cast field pointer to *uint
func (f *Field) UintPtr(structPtr unsafe.Pointer) *uint {
	result := AsUintAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint64 cast field pointer to uint64
func (f *Field) Uint64(structPtr unsafe.Pointer) uint64 {
	return AsUint64(f.Pointer(structPtr))
}

//Uint64Ptr cast field pointer to *uint64
func (f *Field) Uint64Ptr(structPtr unsafe.Pointer) *uint64 {
	result := AsUint64AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint64Addr cast field pointer to *uint64
func (f *Field) Uint64Addr(structPtr unsafe.Pointer) *uint64 {
	return AsUint64Ptr(f.Pointer(structPtr))
}

//Uint32 cast field pointer to uint32
func (f *Field) Uint32(structPtr unsafe.Pointer) uint32 {
	return AsUint32(f.Pointer(structPtr))
}

//Uint32Ptr cast field pointer to *uint32
func (f *Field) Uint32Ptr(structPtr unsafe.Pointer) *uint32 {
	result := AsUint32AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint32Addr cast field pointer to *uint32
func (f *Field) Uint32Addr(structPtr unsafe.Pointer) *uint32 {
	return AsUint32Ptr(f.Pointer(structPtr))
}

//Uint16 cast field pointer to uint16
func (f *Field) Uint16(structPtr unsafe.Pointer) uint16 {
	return AsUint16(f.Pointer(structPtr))
}

//Uint16Ptr cast field pointer to *uint16
func (f *Field) Uint16Ptr(structPtr unsafe.Pointer) *uint16 {
	result := AsUint16AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint16Addr cast field pointer to *uint16
func (f *Field) Uint16Addr(structPtr unsafe.Pointer) *uint16 {
	return AsUint16Ptr(f.Pointer(structPtr))
}

//Uint8 cast field pointer to uint8
func (f *Field) Uint8(structPtr unsafe.Pointer) uint8 {
	return AsUint8(f.Pointer(structPtr))
}

//Uint8Ptr cast field pointer to *uint8
func (f *Field) Uint8Ptr(structPtr unsafe.Pointer) *uint8 {
	result := AsUint8AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint8Addr cast field pointer to *uint8
func (f *Field) Uint8Addr(structPtr unsafe.Pointer) *uint8 {
	return AsUint8Ptr(f.Pointer(structPtr))
}

//Bool cast field pointer to bool
func (f *Field) Bool(structPtr unsafe.Pointer) bool {
	return AsBool(f.Pointer(structPtr))
}

//BoolPtr cast field pointer to *bool
func (f *Field) BoolPtr(structPtr unsafe.Pointer) *bool {
	result := AsBoolAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//BoolAddr cast field pointer to **bool
func (f *Field) BoolAddr(structPtr unsafe.Pointer) *bool {
	return AsBoolPtr(f.Pointer(structPtr))
}

//Float64 cast field pointer to float64
func (f *Field) Float64(structPtr unsafe.Pointer) float64 {
	return AsFloat64(f.Pointer(structPtr))
}

//Float64Ptr cast field pointer to *float64
func (f *Field) Float64Ptr(structPtr unsafe.Pointer) *float64 {
	result := AsFloat64AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Float64Addr cast field pointer to *float64
func (f *Field) Float64Addr(structPtr unsafe.Pointer) *float64 {
	return AsFloat64Ptr(f.Pointer(structPtr))
}

//Float32 cast field pointer to float32
func (f *Field) Float32(structPtr unsafe.Pointer) float32 {
	return AsFloat32(f.Pointer(structPtr))
}

//Float32Ptr cast field pointer to *float32
func (f *Field) Float32Ptr(structPtr unsafe.Pointer) *float32 {
	result := AsFloat32AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Float32Addr cast field pointer to *float32
func (f *Field) Float32Addr(structPtr unsafe.Pointer) *float32 {
	return AsFloat32Ptr(f.Pointer(structPtr))
}

//String cast field pointer to string
func (f *Field) String(structPtr unsafe.Pointer) string {
	return AsString(f.Pointer(structPtr))
}

//StringPtr cast field pointer to *string
func (f *Field) StringPtr(structPtr unsafe.Pointer) *string {
	result := AsStringAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//StringAddr field pointer to *string
func (f *Field) StringAddr(structPtr unsafe.Pointer) *string {
	return AsStringPtr(f.Pointer(structPtr))
}

//Bytes cast field pointer to []byte
func (f *Field) Bytes(structPtr unsafe.Pointer) []byte {
	return AsUint8s(f.Pointer(structPtr))
}

//BytesPtr cast field pointer to *[]byte
func (f *Field) BytesPtr(structPtr unsafe.Pointer) *[]byte {
	result := AsBytesAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//TimePtr cast field pointer to *time.Time
func (f *Field) TimePtr(structPtr unsafe.Pointer) *time.Time {
	result := AsTimeAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Time cast field pointer to time.Time
func (f *Field) Time(structPtr unsafe.Pointer) time.Time {
	return AsTime(f.Pointer(structPtr))
}

//Value returns field value
func (f *Field) Value(structPtr unsafe.Pointer) interface{} {
	return f.Interface(structPtr)
}
