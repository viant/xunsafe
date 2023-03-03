//go:build debug

package xunsafe

import (
	"github.com/viant/xreflect"
	"time"
	"unsafe"
)

//Interface cast field pointer to value
func (f *Field) Interface(structPtr unsafe.Pointer) interface{} {
	ptr := f.Pointer(structPtr)
	if f.ptrKind == ptrKindEmptyInterface {
		return asInterface(ptr, f.rtype, true)
	}
	if f.ptrKind == ptrKindMethodInterface {
		return *(*interface {
			M()
		})(ptr)
	}
	return *(*interface{})(ptr)
}

//Int cast field pointer to int
func (f *Field) Int(structPtr unsafe.Pointer) int {
	f.MustBeAssignable(xreflect.IntType)
	return AsInt(f.Pointer(structPtr))
}

//IntPtr cast field pointer to *int
func (f *Field) IntPtr(structPtr unsafe.Pointer) *int {
	f.MustBeAssignable(xreflect.IntPtrType)
	result := AsIntAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Error cast field pointer to error
func (f *Field) Error(structPtr unsafe.Pointer) error {
	f.MustBeAssignable(xreflect.ErrorType)
	return AsError(f.Pointer(structPtr))
}

//ErrorPtr cast field pointer to *error
func (f *Field) ErrorPtr(structPtr unsafe.Pointer) *error {
	f.MustBeAssignable(xreflect.ErrorPtrType)
	result := AsErrorAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//IntAddr cast field pointer to **int
func (f *Field) IntAddr(structPtr unsafe.Pointer) *int {
	f.MustBeAssignable(xreflect.IntPtrType)
	return AsIntPtr(f.Pointer(structPtr))
}

//Int64 cast field pointer to int64
func (f *Field) Int64(structPtr unsafe.Pointer) int64 {
	f.MustBeAssignable(xreflect.Int64Type)
	return AsInt64(f.Pointer(structPtr))
}

//Int64Ptr cast field pointer to *int
func (f *Field) Int64Ptr(structPtr unsafe.Pointer) *int64 {
	f.MustBeAssignable(xreflect.Int64PtrType)
	result := AsInt64AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int64Addr cast field pointer to *int64
func (f *Field) Int64Addr(structPtr unsafe.Pointer) *int64 {
	f.MustBeAssignable(xreflect.Int64PtrType)
	return AsInt64Ptr(f.Pointer(structPtr))
}

//Int32 cast field pointer to int32
func (f *Field) Int32(structPtr unsafe.Pointer) int32 {
	f.MustBeAssignable(xreflect.Int32Type)
	return AsInt32(f.Pointer(structPtr))
}

//Int32Ptr cast field pointer to *int32
func (f *Field) Int32Ptr(structPtr unsafe.Pointer) *int32 {
	f.MustBeAssignable(xreflect.Int32PtrType)
	result := AsInt32AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int32Addr cast field pointer to *int32
func (f *Field) Int32Addr(structPtr unsafe.Pointer) *int32 {
	f.MustBeAssignable(xreflect.Int32PtrType)
	return AsInt32Ptr(f.Pointer(structPtr))
}

//Int16 cast field pointer to int16
func (f *Field) Int16(structPtr unsafe.Pointer) int16 {
	f.MustBeAssignable(xreflect.Int16Type)
	return AsInt16(f.Pointer(structPtr))
}

//Int16Ptr cast field pointer to *int16
func (f *Field) Int16Ptr(structPtr unsafe.Pointer) *int16 {
	f.MustBeAssignable(xreflect.Int16PtrType)
	result := AsInt16AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int16Addr returns *int16
func (f *Field) Int16Addr(structPtr unsafe.Pointer) *int16 {
	f.MustBeAssignable(xreflect.Int16PtrType)
	return AsInt16Ptr(f.Pointer(structPtr))
}

//Int8 cast field pointer to int8
func (f *Field) Int8(structPtr unsafe.Pointer) int8 {
	f.MustBeAssignable(xreflect.Int8Type)
	return AsInt8(f.Pointer(structPtr))
}

//Int8Ptr cast field pointer to *int8
func (f *Field) Int8Ptr(structPtr unsafe.Pointer) *int8 {
	f.MustBeAssignable(xreflect.Int8PtrType)
	result := AsInt8AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Int8Addr cast field pointer to *int8
func (f *Field) Int8Addr(structPtr unsafe.Pointer) *int8 {
	f.MustBeAssignable(xreflect.Int8PtrType)
	return AsInt8Ptr(f.Pointer(structPtr))
}

//UintAddr cast field pointer to **uint
func (f *Field) UintAddr(structPtr unsafe.Pointer) *uint {
	f.MustBeAssignable(xreflect.UintPtrType)
	return AsUintPtr(f.Pointer(structPtr))
}

//Uint cast field pointer to uint
func (f *Field) Uint(structPtr unsafe.Pointer) uint {
	f.MustBeAssignable(xreflect.UintType)
	return AsUint(f.Pointer(structPtr))
}

//UintPtr cast field pointer to *uint
func (f *Field) UintPtr(structPtr unsafe.Pointer) *uint {
	f.MustBeAssignable(xreflect.UintPtrType)
	result := AsUintAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint64 cast field pointer to uint64
func (f *Field) Uint64(structPtr unsafe.Pointer) uint64 {
	f.MustBeAssignable(xreflect.Uint64Type)
	return AsUint64(f.Pointer(structPtr))
}

//Uint64Ptr cast field pointer to *uint64
func (f *Field) Uint64Ptr(structPtr unsafe.Pointer) *uint64 {
	f.MustBeAssignable(xreflect.Uint64PtrType)
	result := AsUint64AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint64Addr cast field pointer to *uint64
func (f *Field) Uint64Addr(structPtr unsafe.Pointer) *uint64 {
	f.MustBeAssignable(xreflect.Uint64PtrType)
	return AsUint64Ptr(f.Pointer(structPtr))
}

//Uint32 cast field pointer to uint32
func (f *Field) Uint32(structPtr unsafe.Pointer) uint32 {
	f.MustBeAssignable(xreflect.Uint32Type)
	return AsUint32(f.Pointer(structPtr))
}

//Uint32Ptr cast field pointer to *uint32
func (f *Field) Uint32Ptr(structPtr unsafe.Pointer) *uint32 {
	f.MustBeAssignable(xreflect.Uint32PtrType)
	result := AsUint32AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint32Addr cast field pointer to *uint32
func (f *Field) Uint32Addr(structPtr unsafe.Pointer) *uint32 {
	f.MustBeAssignable(xreflect.Uint32PtrType)
	return AsUint32Ptr(f.Pointer(structPtr))
}

//Uint16 cast field pointer to uint16
func (f *Field) Uint16(structPtr unsafe.Pointer) uint16 {
	f.MustBeAssignable(xreflect.Uint16Type)
	return AsUint16(f.Pointer(structPtr))
}

//Uint16Ptr cast field pointer to *uint16
func (f *Field) Uint16Ptr(structPtr unsafe.Pointer) *uint16 {
	f.MustBeAssignable(xreflect.Uint16PtrType)
	result := AsUint16AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint16Addr cast field pointer to *uint16
func (f *Field) Uint16Addr(structPtr unsafe.Pointer) *uint16 {
	f.MustBeAssignable(xreflect.Uint16PtrType)
	return AsUint16Ptr(f.Pointer(structPtr))
}

//Uint8 cast field pointer to uint8
func (f *Field) Uint8(structPtr unsafe.Pointer) uint8 {
	f.MustBeAssignable(xreflect.Uint8Type)
	return AsUint8(f.Pointer(structPtr))
}

//Uint8Ptr cast field pointer to *uint8
func (f *Field) Uint8Ptr(structPtr unsafe.Pointer) *uint8 {
	f.MustBeAssignable(xreflect.Uint8PtrType)
	result := AsUint8AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Uint8Addr cast field pointer to *uint8
func (f *Field) Uint8Addr(structPtr unsafe.Pointer) *uint8 {
	f.MustBeAssignable(xreflect.Uint8PtrType)
	return AsUint8Ptr(f.Pointer(structPtr))
}

//Bool cast field pointer to bool
func (f *Field) Bool(structPtr unsafe.Pointer) bool {
	f.MustBeAssignable(xreflect.BoolType)
	return AsBool(f.Pointer(structPtr))
}

//BoolPtr cast field pointer to *bool
func (f *Field) BoolPtr(structPtr unsafe.Pointer) *bool {
	f.MustBeAssignable(xreflect.BoolPtrType)
	result := AsBoolAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//BoolAddr cast field pointer to **bool
func (f *Field) BoolAddr(structPtr unsafe.Pointer) *bool {
	f.MustBeAssignable(xreflect.BoolPtrType)
	return AsBoolPtr(f.Pointer(structPtr))
}

//Float64 cast field pointer to float64
func (f *Field) Float64(structPtr unsafe.Pointer) float64 {
	f.MustBeAssignable(xreflect.Float64Type)
	return AsFloat64(f.Pointer(structPtr))
}

//Float64Ptr cast field pointer to *float64
func (f *Field) Float64Ptr(structPtr unsafe.Pointer) *float64 {
	f.MustBeAssignable(xreflect.Float64PtrType)
	result := AsFloat64AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Float64Addr cast field pointer to *float64
func (f *Field) Float64Addr(structPtr unsafe.Pointer) *float64 {
	f.MustBeAssignable(xreflect.Float64PtrType)
	return AsFloat64Ptr(f.Pointer(structPtr))
}

//Float32 cast field pointer to float32
func (f *Field) Float32(structPtr unsafe.Pointer) float32 {
	f.MustBeAssignable(xreflect.Float32Type)
	return AsFloat32(f.Pointer(structPtr))
}

//Float32Ptr cast field pointer to *float32
func (f *Field) Float32Ptr(structPtr unsafe.Pointer) *float32 {
	f.MustBeAssignable(xreflect.Float32PtrType)
	result := AsFloat32AddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Float32Addr cast field pointer to *float32
func (f *Field) Float32Addr(structPtr unsafe.Pointer) *float32 {
	f.MustBeAssignable(xreflect.Float32PtrType)
	return AsFloat32Ptr(f.Pointer(structPtr))
}

//String cast field pointer to string
func (f *Field) String(structPtr unsafe.Pointer) string {
	f.MustBeAssignable(xreflect.StringType)
	return AsString(f.Pointer(structPtr))
}

//StringPtr cast field pointer to *string
func (f *Field) StringPtr(structPtr unsafe.Pointer) *string {
	f.MustBeAssignable(xreflect.StringPtrType)
	result := AsStringAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//StringAddr field pointer to *string
func (f *Field) StringAddr(structPtr unsafe.Pointer) *string {
	f.MustBeAssignable(xreflect.StringPtrType)
	return AsStringPtr(f.Pointer(structPtr))
}

//Bytes cast field pointer to []byte
func (f *Field) Bytes(structPtr unsafe.Pointer) []byte {
	f.MustBeAssignable(xreflect.BytesType)
	return AsUint8s(f.Pointer(structPtr))
}

//BytesPtr cast field pointer to *[]byte
func (f *Field) BytesPtr(structPtr unsafe.Pointer) *[]byte {
	f.MustBeAssignable(xreflect.BytesPtrType)
	result := AsBytesAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//TimePtr cast field pointer to *time.Time
func (f *Field) TimePtr(structPtr unsafe.Pointer) *time.Time {
	f.MustBeAssignable(xreflect.TimePtrType)
	result := AsTimeAddrPtr(f.Pointer(structPtr))
	if result == nil {
		return nil
	}
	return *result
}

//Time cast field pointer to time.Time
func (f *Field) Time(structPtr unsafe.Pointer) time.Time {
	f.MustBeAssignable(xreflect.TimeType)
	return AsTime(f.Pointer(structPtr))
}

//Value returns field value
func (f *Field) Value(structPtr unsafe.Pointer) interface{} {
	return f.Interface(structPtr)
}
