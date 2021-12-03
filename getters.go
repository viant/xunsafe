package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

type GetterInto = func(structAddr unsafe.Pointer, dest interface{})

var gettersInto = initGettersInto()

func initGettersInto() []func(f *Field) GetterInto {
	result := make([]func(field *Field) GetterInto, reflect.UnsafePointer)

	result[reflect.Bool] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*bool) = f.Bool(structAddr)
		}
	}

	result[reflect.Int] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*int) = f.Int(structAddr)
		}
	}

	result[reflect.Int64] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*int64) = f.Int64(structAddr)
		}
	}

	result[reflect.Int32] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*int32) = f.Int32(structAddr)
		}
	}

	result[reflect.Int16] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*int16) = f.Int16(structAddr)
		}
	}

	result[reflect.Int8] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*int8) = f.Int8(structAddr)
		}
	}

	result[reflect.Uint] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*uint) = f.Uint(structAddr)
		}
	}

	result[reflect.Uint64] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*uint64) = f.Uint64(structAddr)
		}
	}

	result[reflect.Uint32] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*uint32) = f.Uint32(structAddr)
		}
	}

	result[reflect.Uint16] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*uint16) = f.Uint16(structAddr)
		}
	}

	result[reflect.Uint8] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*uint8) = f.Uint8(structAddr)
		}
	}

	result[reflect.Float64] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*float64) = f.Float64(structAddr)
		}
	}

	result[reflect.Float32] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*float32) = f.Float32(structAddr)
		}
	}

	result[reflect.String] = func(f *Field) GetterInto {
		return func(structAddr unsafe.Pointer, dest interface{}) {
			*dest.(*string) = f.String(structAddr)
		}
	}

	return result
}

func FieldGetterInto(f *Field) GetterInto {
	if getterInto := gettersInto[f.kind]; getterInto != nil {
		return getterInto(f)
	} else {
		if f.field.Type.ConvertibleTo(timeType) {
			return func(structAddr unsafe.Pointer, val interface{}) {
				*val.(*time.Time) = f.Time(structAddr)
			}
		}
		if f.field.Type.ConvertibleTo(timeTypePtr) {
			return func(structAddr unsafe.Pointer, val interface{}) {
				*val.(**time.Time) = f.TimePtr(structAddr)
			}
		}
		return func(structAddr unsafe.Pointer, val interface{}) {
			value := f.Value(structAddr)
			*val.(*interface{}) = value
		}
	}
}
