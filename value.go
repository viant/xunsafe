package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//fieldValue creates Getter function for a field value or error
func (f *Field) fieldValue() Getter {
	offset := f.field.Offset
	if f.Address != nil {
		return func(structAddr uintptr) interface{} {
			return f.Address(structAddr)
		}
	}
	switch f.field.Type.Kind() {
	case reflect.Int:
		return func(structAddr uintptr) interface{} {
			result := (*int)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint:
		return func(structAddr uintptr) interface{} {
			result := (*uint)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int64:
		return func(structAddr uintptr) interface{} {
			result := (*int64)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int32:
		return func(structAddr uintptr) interface{} {
			result := (*int32)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int16:
		return func(structAddr uintptr) interface{} {
			result := (*int16)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Int8:
		return func(structAddr uintptr) interface{} {
			result := (*int8)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint64:
		return func(structAddr uintptr) interface{} {
			result := (*uint64)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint32:
		return func(structAddr uintptr) interface{} {
			result := (*uint32)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint16:
		return func(structAddr uintptr) interface{} {
			result := (*uint16)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Uint8:
		return func(structAddr uintptr) interface{} {
			result := (*uint8)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.String:
		return func(structAddr uintptr) interface{} {
			result := (*string)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Float64:
		return func(structAddr uintptr) interface{} {
			result := (*float64)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}

	case reflect.Float32:
		return func(structAddr uintptr) interface{} {
			result := (*float32)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}
	case reflect.Bool:
		return func(structAddr uintptr) interface{} {
			result := (*bool)(unsafe.Pointer(structAddr + offset))
			if result == nil {
				return nil
			}
			return *result
		}

	case reflect.Struct:

		if f.field.Type.ConvertibleTo(timeType) {
			return func(structAddr uintptr) interface{} {
				result := (*time.Time)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		}
		if getter := _registry.Lookup(f.field.Type); getter != nil {
			return func(structAddr uintptr) interface{} {
				return getter(structAddr + offset)
			}
		}
		if f.Field == nil {
			return func(structAddr uintptr) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
				return fieldValue.Interface()
			}
		}
		fn := f.AddrGetter()
		return func(structAddr uintptr) interface{} {
			fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
			addr := fieldValue.Elem().UnsafeAddr()
			return fn(addr)
		}

	case reflect.Ptr:
		switch f.field.Type.Elem().Kind() {
		case reflect.Struct:
			if f.field.Type.ConvertibleTo(timeTypePtr) {
				return func(structAddr uintptr) interface{} {
					result := (**time.Time)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			}
			if getter := _registry.Lookup(f.field.Type); getter != nil {
				return func(structAddr uintptr) interface{} {
					return getter(structAddr + offset)
				}
			}
			if f.Field == nil {
				return func(structAddr uintptr) interface{} {
					fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
					if fieldValue.Elem().IsNil() {
						ptr := reflect.New(fieldValue.Type().Elem().Elem())
						fieldValue.Elem().Set(ptr)
					}
					return fieldValue.Interface()
				}
			}
			fn := f.Field.AddrGetter()
			return func(structAddr uintptr) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
				if fieldValue.Elem().IsNil() {
					ptr := reflect.New(fieldValue.Type().Elem().Elem())
					fieldValue.Elem().Set(ptr)
				}
				return fn(fieldValue.Elem().Elem().UnsafeAddr())
			}

		case reflect.Int:
			return func(structAddr uintptr) interface{} {
				result := (**int)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint:
			return func(structAddr uintptr) interface{} {
				result := (**uint)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int64:
			return func(structAddr uintptr) interface{} {
				result := (**int64)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int32:
			return func(structAddr uintptr) interface{} {
				result := (**int32)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int16:
			return func(structAddr uintptr) interface{} {
				result := (**int16)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int8:
			return func(structAddr uintptr) interface{} {
				result := (**int8)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint64:
			return func(structAddr uintptr) interface{} {
				result := (**uint64)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint32:
			return func(structAddr uintptr) interface{} {
				result := (**uint32)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint16:
			return func(structAddr uintptr) interface{} {
				result := (**uint16)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint8:
			return func(structAddr uintptr) interface{} {
				result := (**uint8)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.String:
			return func(structAddr uintptr) interface{} {
				result := (**string)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Float64:
			return func(structAddr uintptr) interface{} {
				result := (**float64)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}

		case reflect.Float32:
			return func(structAddr uintptr) interface{} {
				result := (**float32)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Bool:
			return func(structAddr uintptr) interface{} {
				result := (**bool)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Slice:
			switch f.field.Type.Elem().Elem().Kind() {
			case reflect.Int:
				return func(structAddr uintptr) interface{} {
					result := (**[]int)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint:
				return func(structAddr uintptr) interface{} {
					result := (**[]uint)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int64:
				return func(structAddr uintptr) interface{} {
					result := (**[]int64)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int32:
				return func(structAddr uintptr) interface{} {
					result := (**[]int32)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int16:
				return func(structAddr uintptr) interface{} {
					result := (**[]int16)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Int8:
				return func(structAddr uintptr) interface{} {
					result := (**[]int8)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint64:
				return func(structAddr uintptr) interface{} {
					result := (**[]uint64)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint32:
				return func(structAddr uintptr) interface{} {
					result := (**[]uint32)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint16:
				return func(structAddr uintptr) interface{} {
					result := (**[]uint16)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Uint8:
				return func(structAddr uintptr) interface{} {
					result := (**[]uint8)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.String:
				return func(structAddr uintptr) interface{} {
					result := (**[]string)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Float64:
				return func(structAddr uintptr) interface{} {
					result := (**[]float64)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}

			case reflect.Float32:
				return func(structAddr uintptr) interface{} {
					result := (**[]float32)(unsafe.Pointer(structAddr + offset))
					if result == nil {
						return nil
					}
					return *result
				}
			case reflect.Bool:
				return func(structAddr uintptr) interface{} {
					result := (**[]bool)(unsafe.Pointer(structAddr + offset))
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
			return func(structAddr uintptr) interface{} {
				result := (*[]int)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint:
			return func(structAddr uintptr) interface{} {
				result := (*[]uint)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int64:
			return func(structAddr uintptr) interface{} {
				result := (*[]int64)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int32:
			return func(structAddr uintptr) interface{} {
				result := (*[]int32)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int16:
			return func(structAddr uintptr) interface{} {
				result := (*[]int16)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Int8:
			return func(structAddr uintptr) interface{} {
				result := (*[]int8)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint64:
			return func(structAddr uintptr) interface{} {
				result := (*[]uint64)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint32:
			return func(structAddr uintptr) interface{} {
				result := (*[]uint32)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint16:
			return func(structAddr uintptr) interface{} {
				result := (*[]uint16)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Uint8:
			return func(structAddr uintptr) interface{} {
				result := (*[]uint8)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.String:
			return func(structAddr uintptr) interface{} {
				result := (*[]string)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Float64:
			return func(structAddr uintptr) interface{} {
				result := (*[]float64)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}

		case reflect.Float32:
			return func(structAddr uintptr) interface{} {
				result := (*[]float32)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		case reflect.Bool:
			return func(structAddr uintptr) interface{} {
				result := (*[]bool)(unsafe.Pointer(structAddr + offset))
				if result == nil {
					return nil
				}
				return *result
			}
		}
	}
	return func(structPtr uintptr) interface{} {
		fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structPtr+offset))
		return fieldValue.Elem().Interface()
	}
}
