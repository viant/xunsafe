package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

var mutators = initMutators()

func initMutators() []func(field *Field) Setter {
	result := make([]func(field *Field) Setter, reflect.UnsafePointer)

	result[reflect.Int] = func(f *Field) Setter {
		return f.intMutator
	}
	result[reflect.Int64] = func(f *Field) Setter {
		return f.int64Mutator
	}
	result[reflect.Int32] = func(f *Field) Setter {
		return f.int32Mutator
	}
	result[reflect.Int16] = func(f *Field) Setter {
		return f.int16Mutator
	}
	result[reflect.Int8] = func(f *Field) Setter {
		return f.int8Mutator
	}

	result[reflect.Uint] = func(f *Field) Setter {
		return f.uintMutator
	}
	result[reflect.Uint64] = func(f *Field) Setter {
		return f.uint64Mutator
	}
	result[reflect.Uint32] = func(f *Field) Setter {
		return f.uint32Mutator
	}
	result[reflect.Uint16] = func(f *Field) Setter {
		return f.uint16Mutator
	}

	result[reflect.Uint8] = func(f *Field) Setter {
		return f.uint8Mutator
	}

	result[reflect.String] = func(f *Field) Setter {
		return f.stringMutator
	}
	result[reflect.Float64] = func(f *Field) Setter {
		return f.float64Mutator
	}
	result[reflect.Float32] = func(f *Field) Setter {
		return f.float32Mutator
	}
	result[reflect.Bool] = func(f *Field) Setter {
		return f.boolMutator
	}

	return result
}

func (f *Field) intMutator(addr unsafe.Pointer, val interface{}) {
	f.SetInt(addr, val.(int))
}

func (f *Field) int64Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetInt64(addr, val.(int64))
}

func (f *Field) int32Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetInt32(addr, val.(int32))
}

func (f *Field) int16Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetInt16(addr, val.(int16))
}

func (f *Field) int8Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetInt8(addr, val.(int8))
}

func (f *Field) uintMutator(addr unsafe.Pointer, val interface{}) {
	f.SetUint(addr, val.(uint))
}

func (f *Field) uint64Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetUint64(addr, val.(uint64))
}

func (f *Field) uint32Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetUint32(addr, val.(uint32))
}

func (f *Field) uint16Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetUint16(addr, val.(uint16))
}

func (f *Field) uint8Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetUint8(addr, val.(uint8))
}

func (f *Field) float64Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetFloat64(addr, val.(float64))
}

func (f *Field) float32Mutator(addr unsafe.Pointer, val interface{}) {
	f.SetFloat32(addr, val.(float32))
}

func (f *Field) stringMutator(addr unsafe.Pointer, val interface{}) {
	f.SetString(addr, val.(string))
}

func (f *Field) boolMutator(addr unsafe.Pointer, val interface{}) {
	f.SetBool(addr, val.(bool))
}

func (f *Field) interfaceMutator(addr unsafe.Pointer, val interface{}) {
	f.SetInterface(addr, val)
}

func FieldMutator(f *Field) Setter {
	if mutator := mutators[f.kind]; mutator != nil {
		return mutator(f)
	} else {
		if f.field.Type.ConvertibleTo(timeType) {
			return func(structAddr unsafe.Pointer, val interface{}) {
				f.SetTime(structAddr, val.(time.Time))
			}
		}
		if f.field.Type.ConvertibleTo(timeTypePtr) {
			return func(structAddr unsafe.Pointer, val interface{}) {
				f.SetTimePtr(structAddr, val.(*time.Time))
			}
		}
		return func(structAddr unsafe.Pointer, val interface{}) {
			f.SetValue(structAddr, val)
		}
	}
}
