package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//fieldValue creates Getter function for a field value or error
func (f *Field) fieldValue() Getter {
	offset := f.field.Offset
	if f.value != nil {
		return func(structAddr unsafe.Pointer) interface{} {
			return f.value(structAddr)
		}
	}

	switch f.field.Type.Kind() {
	case reflect.Int:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*int)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*uint)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int64:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*int64)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int32:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*int32)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int16:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*int16)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int8:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*int8)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint64:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint32:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint16:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint8:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.String:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*string)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Float64:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*float64)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}

	case reflect.Float32:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*float32)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Bool:
		return func(structAddr unsafe.Pointer) interface{} {
			result := (*bool)(unsafe.Pointer(uintptr(structAddr) + offset))
			if result == nil {
				return nil
			}
			return *result
		}

	case reflect.Struct:

		if f.field.Type.ConvertibleTo(timeType) {
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*time.Time)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
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

	case reflect.Ptr:
		switch f.field.Type.Elem().Kind() {
		case reflect.Struct:
			if f.field.Type.ConvertibleTo(timeTypePtr) {
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**time.Time)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
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
					return newValue.Elem().Interface()
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

		case reflect.Int:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**int)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**uint)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int64:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**int64)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int32:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**int32)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int16:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**int16)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int8:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**int8)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint64:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint32:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint16:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint8:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.String:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**string)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Float64:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**float64)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}

		case reflect.Float32:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**float32)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Bool:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (**bool)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Slice:
			switch f.field.Type.Elem().Elem().Kind() {
			case reflect.Int:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]int)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]uint)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int64:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]int64)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int32:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]int32)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int16:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]int16)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int8:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]int8)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint64:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint32:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint16:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint8:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.String:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]string)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Float64:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]float64)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}

			case reflect.Float32:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]float32)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Bool:
				return func(structAddr unsafe.Pointer) interface{} {
					result := (**[]bool)(unsafe.Pointer(uintptr(structAddr) + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			}
		}
	case reflect.Slice:
		switch f.field.Type.Elem().Kind() {
		case reflect.Int:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]int)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]uint)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int64:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]int64)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int32:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]int32)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int16:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]int16)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int8:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]int8)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint64:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint32:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint16:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint8:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.String:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]string)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Float64:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]float64)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}

		case reflect.Float32:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]float32)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Bool:
			return func(structAddr unsafe.Pointer) interface{} {
				result := (*[]bool)(unsafe.Pointer(uintptr(structAddr) + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		}
	}
	return func(structPtr unsafe.Pointer) interface{} {
		fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structPtr)+offset))
		return fieldValue.Elem().Interface()
	}
}
