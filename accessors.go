package xunsafe

import (
	"reflect"
	"unsafe"
)

var accessors = initAccessors()

func initAccessors() []func(f *Field) Getter {
	newAccessors := make([]func(f *Field) Getter, reflect.UnsafePointer)
	newAccessors[reflect.Int] = func(f *Field) Getter {
		return f.intAccessor
	}
	newAccessors[reflect.Int64] = func(f *Field) Getter {
		return f.int64Accessor
	}
	newAccessors[reflect.Int32] = func(f *Field) Getter {
		return f.int32Accessor
	}
	newAccessors[reflect.Int16] = func(f *Field) Getter {
		return f.int16Accessor
	}
	newAccessors[reflect.Int8] = func(f *Field) Getter {
		return f.int8Accessor
	}

	newAccessors[reflect.Uint] = func(f *Field) Getter {
		return f.uintAccessor
	}
	newAccessors[reflect.Uint64] = func(f *Field) Getter {
		return f.uint64Accessor
	}
	newAccessors[reflect.Uint32] = func(f *Field) Getter {
		return f.uint32Accessor
	}
	newAccessors[reflect.Uint16] = func(f *Field) Getter {
		return f.uint16Accessor
	}

	newAccessors[reflect.Uint8] = func(f *Field) Getter {
		return f.uint8Accessor
	}
	newAccessors[reflect.String] = func(f *Field) Getter {
		return f.stringAccessor
	}
	newAccessors[reflect.Float64] = func(f *Field) Getter {
		return f.float64Accessor
	}
	newAccessors[reflect.Float32] = func(f *Field) Getter {
		return f.float32Accessor
	}
	newAccessors[reflect.Bool] = func(f *Field) Getter {
		return f.boolAccessor
	}

	newAccessors[reflect.Func] = func(f *Field) Getter {
		return func(structPtr unsafe.Pointer) interface{} {
			fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structPtr)+f.field.Offset))
			return fieldValue.Elem().Interface()
		}
	}

	newAccessors[reflect.Slice] = func(f *Field) Getter {
		return f.getSliceAccessor()
	}

	newAccessors[reflect.Ptr] = func(f *Field) Getter {
		return f.getPointerAccessor()
	}

	newAccessors[reflect.Struct] = func(f *Field) Getter {
		return f.getStructAccessor()
	}
	return newAccessors
}


func FieldAccessor(f *Field) Getter {
	return accessors[f.kind](f)
}

func (f *Field) intAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Int(structAddr)
}

func (f *Field) int64Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Int64(structAddr)
}

func (f *Field) int32Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Int32(structAddr)
}

func (f *Field) int16Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Int16(structAddr)
}

func (f *Field) int8Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Int8(structAddr)
}

func (f *Field) uintAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint(structAddr)
}

func (f *Field) uint64Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint64(structAddr)
}

func (f *Field) uint32Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint32(structAddr)
}

func (f *Field) uint16Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint16(structAddr)
}

func (f *Field) uint8Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint8(structAddr)
}

func (f *Field) stringAccessor(structAddr unsafe.Pointer) interface{} {
	return f.String(structAddr)
}

func (f *Field) boolAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Bool(structAddr)
}

func (f *Field) float64Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Float64(structAddr)
}

func (f *Field) float32Accessor(structAddr unsafe.Pointer) interface{} {
	return f.Float32(structAddr)
}

func (f *Field) timeAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Time(structAddr)
}

func (f *Field) intSliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceInt(structAddr)
}

func (f *Field) uintSliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceUint(structAddr)
}

func (f *Field) int64SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceInt64(structAddr)
}

func (f *Field) int32SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceInt32(structAddr)
}

func (f *Field) int16SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceInt16(structAddr)
}

func (f *Field) int8SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceInt8(structAddr)
}

func (f *Field) uint64SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceUint64(structAddr)
}

func (f *Field) uint32SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceUint32(structAddr)
}

func (f *Field) uint16SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceUint16(structAddr)
}

func (f *Field) uint8SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceUint8(structAddr)
}

func (f *Field) stringSliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceString(structAddr)
}

func (f *Field) float64SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceFloat64(structAddr)
}

func (f *Field) float32SliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceFloat32(structAddr)
}

func (f *Field) boolSliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.SliceBool(structAddr)
}

func (f *Field) interfaceSliceAccessor(structAddr unsafe.Pointer) interface{} {
	return f.InterfaceSlice(structAddr)
}

func (f *Field) timePtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.TimePtr(structAddr)
}

func (f *Field) getPtrStructAccessor(offset uintptr) Getter {
	if f.field.Type.ConvertibleTo(timeTypePtr) {
		return f.timePtrAccessor
	}
	if getter := lookup(f.field.Type); getter != nil {
		return func(structAddr unsafe.Pointer) interface{} {
			return getter(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	}
	if f.Field == nil {
		return func(structAddr unsafe.Pointer) interface{} {
			newValue := reflect.New(f.field.Type)
			actualPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(structAddr) + offset))
			if actualPtr == nil {
				return nil
			}
			newPtr := (*unsafe.Pointer)(unsafe.Pointer(newValue.Elem().UnsafeAddr()))
			*newPtr = *actualPtr
			elem := newValue.Elem()
			return elem.Interface()
		}
	}
	fn := f.Field.AddrGetter()
	return func(structAddr unsafe.Pointer) interface{} {
		fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
		if fieldValue.Elem().IsNil() {
			ptr := reflect.New(fieldValue.Type().Elem().Elem())
			fieldValue.Elem().Set(ptr)
		}
		return fn(unsafe.Pointer(fieldValue.Elem().Elem().UnsafeAddr()))
	}
}

func (f *Field) intPtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.IntPtr(structAddr)
}

func (f *Field) uintPtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.UintPtr(structAddr)
}

func (f *Field) int64PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Int64Ptr(structAddr)
}

func (f *Field) int32PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Int32Ptr(structAddr)
}

func (f *Field) int16PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Int16Ptr(structAddr)
}

func (f *Field) int8PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Int8Ptr(structAddr)
}

func (f *Field) uint64PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint64Ptr(structAddr)
}

func (f *Field) uint32PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint32Ptr(structAddr)
}

func (f *Field) uint16PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint16Ptr(structAddr)
}

func (f *Field) uint8PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Uint8Ptr(structAddr)
}

func (f *Field) stringPtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.StringPtr(structAddr)
}

func (f *Field) float64PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Float64Ptr(structAddr)
}

func (f *Field) float32PtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.Float32Ptr(structAddr)
}

func (f *Field) boolPtrAccessor(structAddr unsafe.Pointer) interface{} {
	return f.BoolPtr(structAddr)
}

func (f *Field) pointerSliceIntAccessor(structAddr unsafe.Pointer) interface{} {
	offset := f.field.Offset
	return func(structAddr unsafe.Pointer) interface{} {
		result := (**[]int)(unsafe.Pointer(uintptr(structAddr) + offset))
		if result == nil {
			return nil
		}
		return *result
	}
}

func (f *Field) pointerSliceUintAccessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceUint(structAddr)
}

func (f *Field) pointerSliceInt64Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceInt64(structAddr)
}

func (f *Field) pointerSliceInt32Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceInt32(structAddr)
}

func (f *Field) pointerSliceInt16Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceInt16(structAddr)
}

func (f *Field) pointerSliceInt8Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceInt8(structAddr)
}

func (f *Field) pointerSliceUInt64Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceUInt64(structAddr)
}

func (f *Field) pointerSliceUInt32Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceUInt32(structAddr)
}

func (f *Field) pointerSliceUInt16Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceUInt16(structAddr)
}

func (f *Field) pointerSliceUInt8Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceUInt8(structAddr)
}

func (f *Field) pointerSliceStringAccessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceString(structAddr)
}

func (f *Field) pointerSliceFloat64Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceFloat64(structAddr)
}

func (f *Field) pointerSliceFloat32Accessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceFloat32(structAddr)
}

func (f *Field) pointerSliceBoolAccessor(structAddr unsafe.Pointer) interface{} {
	return f.PtrSliceBool(structAddr)
}