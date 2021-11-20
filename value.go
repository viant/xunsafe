package xunsafe

import (
	"reflect"
	"unsafe"
)

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
	panic("unsupported type")
}
