package xunsafe

import (
	"reflect"
	"time"
)

// Precomputed reflect.Type values mirroring github.com/viant/xreflect/buildin.go.
// Keeping them here removes the external xreflect dependency while preserving
// the same runtime type-checking behaviour in debug builds.

var IntType = reflect.TypeOf(0)
var IntPtrType = reflect.PointerTo(IntType)
var Int8Type = reflect.TypeOf(int8(0))
var Int8PtrType = reflect.PointerTo(Int8Type)
var Int16Type = reflect.TypeOf(int16(0))
var Int16PtrType = reflect.PointerTo(Int16Type)
var Int32Type = reflect.TypeOf(int32(0))
var Int32PtrType = reflect.PointerTo(Int32Type)
var Int64Type = reflect.TypeOf(int64(0))
var Int64PtrType = reflect.PointerTo(Int64Type)

var UintType = reflect.TypeOf(uint(0))
var UintPtrType = reflect.PointerTo(UintType)
var Uint8Type = reflect.TypeOf(uint8(0))
var Uint8PtrType = reflect.PointerTo(Uint8Type)
var Uint16Type = reflect.TypeOf(uint16(0))
var Uint16PtrType = reflect.PointerTo(Uint16Type)
var Uint32Type = reflect.TypeOf(uint32(0))
var Uint32PtrType = reflect.PointerTo(Uint32Type)
var Uint64Type = reflect.TypeOf(uint64(0))
var Uint64PtrType = reflect.PointerTo(Uint64Type)

var Float32Type = reflect.TypeOf(float32(0.0))
var Float32PtrType = reflect.PointerTo(Float32Type)
var Float64Type = reflect.TypeOf(0.0)
var Float64PtrType = reflect.PointerTo(Float64Type)

var StringType = reflect.TypeOf("")
var StringPtrType = reflect.PointerTo(StringType)
var BoolType = reflect.TypeOf(false)
var BoolPtrType = reflect.PointerTo(BoolType)

var TimeType = reflect.TypeOf(time.Time{})
var TimePtrType = reflect.PointerTo(TimeType)

var InterfaceType = reflect.TypeOf(typeHolder{}).Field(0).Type
var InterfacePtrType = reflect.PointerTo(InterfaceType)

var ErrorType = reflect.TypeOf(typeHolder{}).Field(1).Type
var ErrorPtrType = reflect.PointerTo(ErrorType)

var BytesType = reflect.TypeOf([]byte{})
var BytesPtrType = reflect.PointerTo(BytesType)

type typeHolder struct {
	ifaceField interface{}
	errField   error
}
