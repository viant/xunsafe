package xunsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

var accessors = initAccessors()

func initAccessors() []func(f *Field) Getter {
	var accesors = make([]func(f *Field) Getter, reflect.UnsafePointer)
	accesors[reflect.Int] = func(f *Field) Getter {
		return f.intAccessor
	}
	accesors[reflect.Int64] = func(f *Field) Getter {
		return f.int64Accessor
	}
	accesors[reflect.Int32] = func(f *Field) Getter {
		return f.int32Accessor
	}
	accesors[reflect.Int16] = func(f *Field) Getter {
		return f.int16Accessor
	}
	accesors[reflect.Int8] = func(f *Field) Getter {
		return f.int8Accessor
	}

	accesors[reflect.Uint] = func(f *Field) Getter {
		return f.uintAccessor
	}
	accesors[reflect.Uint64] = func(f *Field) Getter {
		return f.uint64Accessor
	}
	accesors[reflect.Uint32] = func(f *Field) Getter {
		return f.uint32Accessor
	}
	accesors[reflect.Uint16] = func(f *Field) Getter {
		return f.uint16Accessor
	}

	accesors[reflect.Uint8] = func(f *Field) Getter {
		return f.uint8Accessor
	}
	accesors[reflect.String] = func(f *Field) Getter {
		return f.stringAccessor
	}
	accesors[reflect.Float64] = func(f *Field) Getter {
		return f.float64Accessor
	}
	accesors[reflect.Float32] = func(f *Field) Getter {
		return f.float32Accessor
	}
	accesors[reflect.Bool] = func(f *Field) Getter {
		return f.boolAccessor
	}

	accesors[reflect.Func] = func(f *Field) Getter {
		return func(structPtr unsafe.Pointer) interface{} {
			fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structPtr)+f.field.Offset))
			return fieldValue.Elem().Interface()
		}
	}

	accesors[reflect.Slice] = func(f *Field) Getter {
		return f.getSliceAccessor()
	}

	accesors[reflect.Ptr] = func(f *Field) Getter {
		return f.getPointerAccessor()
	}

	accesors[reflect.Struct] = func(f *Field) Getter {
		return f.getStructAccessor()
	}
	return accesors
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

func (f *Field) getStructAccessor() Getter {
	offset := f.field.Offset
	if f.field.Type.ConvertibleTo(timeType) {
		return f.timeAccessor
	}
	if getter := lookup(f.field.Type); getter != nil {
		return func(structAddr unsafe.Pointer) interface{} {
			return getter(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	}
	if f.Field == nil {
		return func(structAddr unsafe.Pointer) interface{} {
			fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
			return fieldValue.Interface()
		}
	}
	fn := f.AddrGetter()
	return func(structAddr unsafe.Pointer) interface{} {
		fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
		return fn(unsafe.Pointer(fieldValue.Elem().UnsafeAddr()))
	}
}

func (f *Field) getSliceAccessor() Getter {
	elemKind := f.field.Type.Elem().Kind()
	switch elemKind {
	case reflect.Int:
		return f.intSliceAccessor
	case reflect.Uint:
		return f.uintSliceAccessor
	case reflect.Int64:
		return f.int64SliceAccessor
	case reflect.Int32:
		return f.int32SliceAccessor
	case reflect.Int16:
		return f.int16SliceAccessor
	case reflect.Int8:
		return f.int8SliceAccessor
	case reflect.Uint64:
		return f.uint64SliceAccessor
	case reflect.Uint32:
		return f.uint32SliceAccessor
	case reflect.Uint16:
		return f.uint16SliceAccessor
	case reflect.Uint8:
		return f.uint8SliceAccessor
	case reflect.String:
		return f.stringSliceAccessor
	case reflect.Float64:
		return f.float64SliceAccessor
	case reflect.Float32:
		return f.float32Accessor
	case reflect.Bool:
		return f.boolSliceAccessor
	case reflect.Struct:
		return f.interfaceSliceAccessor
	default:
		return func(structPtr unsafe.Pointer) interface{} {
			fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structPtr)+f.field.Offset))
			return fieldValue.Elem().Interface()
		}
	}
}

func (f *Field) getPointerAccessor() Getter {
	offset := f.field.Offset
	switch f.field.Type.Elem().Kind() {
	case reflect.Struct:
		return f.getPtrStructAccessor(offset)
	case reflect.Int:
		return f.intPtrAccessor
	case reflect.Uint:
		return f.uintPtrAccessor
	case reflect.Int64:
		return f.int64PtrAccessor
	case reflect.Int32:
		return f.int32PtrAccessor
	case reflect.Int16:
		return f.int16PtrAccessor
	case reflect.Int8:
		return f.int8PtrAccessor
	case reflect.Uint64:
		return f.uint64Accessor
	case reflect.Uint32:
		return f.uint32PtrAccessor
	case reflect.Uint16:
		return f.uint16PtrAccessor
	case reflect.Uint8:
		return f.uint8PtrAccessor
	case reflect.String:
		return f.stringAccessor
	case reflect.Float64:
		return f.float64PtrAccessor
	case reflect.Float32:
		return f.float32PtrAccessor
	case reflect.Bool:
		return f.boolPtrAccessor
	case reflect.Slice:
		switch f.field.Type.Elem().Elem().Kind() {
		case reflect.Int:
			return f.pointerSliceIntAccessor
		case reflect.Uint:
			return f.pointerSliceUintAccessor
		case reflect.Int64:
			return f.pointerSliceInt64Accessor
		case reflect.Int32:
			return f.pointerSliceInt32Accessor
		case reflect.Int16:
			return f.pointerSliceInt16Accessor
		case reflect.Int8:
			return f.pointerSliceInt8Accessor
		case reflect.Uint64:
			return f.pointerSliceUInt64Accessor
		case reflect.Uint32:
			return f.pointerSliceUInt32Accessor
		case reflect.Uint16:
			return f.pointerSliceUInt16Accessor
		case reflect.Uint8:
			return f.pointerSliceUInt8Accessor
		case reflect.String:
			return f.pointerSliceStringAccessor
		case reflect.Float64:
			return f.pointerSliceFloat64Accessor
		case reflect.Float32:
			return f.pointerSliceFloat32Accessor
		case reflect.Bool:
			return f.pointerSliceBoolAccessor
		}
	}
	panic(fmt.Sprintf("unsupported slice type: %v", f.field.Type.Elem().Elem().Kind()))
}
