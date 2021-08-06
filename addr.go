package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//Addr returns src unsafe addr
func Addr(src interface{}) unsafe.Pointer {
	value := reflect.ValueOf(src)
	if value.Kind() != reflect.Ptr { //convert to a pointer
		vp := reflect.New(value.Type())
		vp.Elem().Set(value)
		value = vp
	}
	return unsafe.Pointer(value.Elem().UnsafeAddr())
}

//AddrGetter creates a Getter function returning filed pointer or error
func (f *Field) AddrGetter() Getter {
	offset := f.field.Offset
	if f.address != nil {
		return func(structAddr unsafe.Pointer) interface{} {
			return f.address(structAddr)
		}
	}
	switch f.field.Type.Kind() {
	case reflect.Int:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*int)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Uint:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*uint)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Int64:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*int64)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Int32:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*int32)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Int16:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*int16)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Int8:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*int8)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Uint64:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Uint32:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Uint16:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Uint8:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.String:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*string)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Float64:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*float64)(unsafe.Pointer(uintptr(structAddr) + offset))
		}

	case reflect.Float32:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*float32)(unsafe.Pointer(uintptr(structAddr) + offset))
		}
	case reflect.Bool:
		return func(structAddr unsafe.Pointer) interface{} {
			return (*bool)(unsafe.Pointer(uintptr(structAddr) + offset))
		}

	case reflect.Struct:
		if f.field.Type.ConvertibleTo(timeType) {
			return func(structAddr unsafe.Pointer) interface{} {
				return (*time.Time)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		}
		if f.Field == nil {
			if getter := lookup(f.field.Type); getter != nil {
				return func(structAddr unsafe.Pointer) interface{} {
					return getter(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			}
			return func(structAddr unsafe.Pointer) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
				return fieldValue.Interface()
			}
		}

		fn := f.Field.AddrGetter()
		return func(structAddr unsafe.Pointer) interface{} {
			fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
			return fn(unsafe.Pointer(fieldValue.Elem().UnsafeAddr()))
		}

	case reflect.Ptr:
		switch f.field.Type.Elem().Kind() {
		case reflect.Struct:
			if f.field.Type.ConvertibleTo(timeTypePtr) {
				return func(structAddr unsafe.Pointer) interface{} {
					return (**time.Time)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			}
			if f.Field == nil {
				if getter := lookup(f.field.Type); getter != nil {
					return func(structAddr unsafe.Pointer) interface{} {
						return getter(unsafe.Pointer(uintptr(structAddr) + offset))
					}
				}
				return interfacePtrGetter(f)
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
				return (**int)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**uint)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int64:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**int64)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int32:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**int32)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int16:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**int16)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int8:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**int8)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint64:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint32:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint16:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint8:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.String:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**string)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Float64:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**float64)(unsafe.Pointer(uintptr(structAddr) + offset))
			}

		case reflect.Float32:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**float32)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Bool:
			return func(structAddr unsafe.Pointer) interface{} {
				return (**bool)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Slice:
			switch f.field.Type.Elem().Elem().Kind() {
			case reflect.Int:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]int)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Uint:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]uint)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Int64:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]int64)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Int32:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]int32)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Int16:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]int16)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Int8:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]int8)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Uint64:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Uint32:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Uint16:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Uint8:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.String:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]string)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Float64:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]float64)(unsafe.Pointer(uintptr(structAddr) + offset))
				}

			case reflect.Float32:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]float32)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			case reflect.Bool:
				return func(structAddr unsafe.Pointer) interface{} {
					return (**[]bool)(unsafe.Pointer(uintptr(structAddr) + offset))
				}
			}
		}
	case reflect.Slice:
		switch f.field.Type.Elem().Kind() {
		case reflect.Int:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]int)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]uint)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int64:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]int64)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int32:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]int32)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int16:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]int16)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Int8:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]int8)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint64:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]uint64)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint32:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]uint32)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint16:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]uint16)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Uint8:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]uint8)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.String:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]string)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Float64:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]float64)(unsafe.Pointer(uintptr(structAddr) + offset))
			}

		case reflect.Float32:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]float32)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Bool:
			return func(structAddr unsafe.Pointer) interface{} {
				return (*[]bool)(unsafe.Pointer(uintptr(structAddr) + offset))
			}
		case reflect.Ptr:
			return func(structAddr unsafe.Pointer) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
				return fieldValue.Interface()
			}
		case reflect.Struct:
			return func(structAddr unsafe.Pointer) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+offset))
				return fieldValue.Interface()
			}
		}
	}
	return func(structAddr unsafe.Pointer) interface{} {
		fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(uintptr(structAddr)+f.field.Offset))
		return fieldValue.Elem().Elem().Interface()
	}
}

func interfacePtrGetter(field *Field) func(structAddr unsafe.Pointer) interface{} {
	return func(structAddr unsafe.Pointer) interface{} {
		fieldValue := reflect.NewAt(field.field.Type, unsafe.Pointer(uintptr(structAddr)+field.field.Offset))
		return fieldValue.Interface()
	}
}
