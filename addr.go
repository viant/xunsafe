package xunsafe

import (
	"reflect"
	"time"
	"unsafe"
)

//Addr returns src unsafe addr
func Addr(src interface{}) uintptr {
	value := reflect.ValueOf(src)
	if value.Kind() != reflect.Ptr { //convert to a pointer
		vp := reflect.New(value.Type())
		vp.Elem().Set(value)
		value = vp
	}
	holderPtr := value.Elem().UnsafeAddr()
	return holderPtr
}

//AddrGetter creates a Getter function returning filed pointer or error
func (f *Field) AddrGetter() Getter {
	offset := f.field.Offset
	if f.Address != nil {
		return func(structAddr uintptr) interface{} {
			return f.Address(structAddr)
		}
	}
	switch f.field.Type.Kind() {
	case reflect.Int:
		return func(structAddr uintptr) interface{} {
			return (*int)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Uint:
		return func(structAddr uintptr) interface{} {
			return (*uint)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Int64:
		return func(structAddr uintptr) interface{} {
			return (*int64)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Int32:
		return func(structAddr uintptr) interface{} {
			return (*int32)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Int16:
		return func(structAddr uintptr) interface{} {
			return (*int16)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Int8:
		return func(structAddr uintptr) interface{} {
			return (*int8)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Uint64:
		return func(structAddr uintptr) interface{} {
			return (*uint64)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Uint32:
		return func(structAddr uintptr) interface{} {
			return (*uint32)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Uint16:
		return func(structAddr uintptr) interface{} {
			return (*uint16)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Uint8:
		return func(structAddr uintptr) interface{} {
			return (*uint8)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.String:
		return func(structAddr uintptr) interface{} {
			return (*string)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Float64:
		return func(structAddr uintptr) interface{} {
			return (*float64)(unsafe.Pointer(structAddr + offset))
		}

	case reflect.Float32:
		return func(structAddr uintptr) interface{} {
			return (*float32)(unsafe.Pointer(structAddr + offset))
		}
	case reflect.Bool:
		return func(structAddr uintptr) interface{} {
			return (*bool)(unsafe.Pointer(structAddr + offset))
		}

	case reflect.Struct:
		if f.field.Type.ConvertibleTo(timeType) {
			return func(structAddr uintptr) interface{} {
				return (*time.Time)(unsafe.Pointer(structAddr + offset))
			}
		}
		if f.Field == nil {
			if getter := _registry.Lookup(f.field.Type); getter != nil {
				return func(structAddr uintptr) interface{} {
					return getter(structAddr + offset)
				}
			}
			return func(structAddr uintptr) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
				return fieldValue.Interface()
			}
		}

		fn := f.Field.AddrGetter()
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
					return (**time.Time)(unsafe.Pointer(structAddr + offset))
				}
			}
			if f.Field == nil {
				if getter := _registry.Lookup(f.field.Type); getter != nil {
					return func(structAddr uintptr) interface{} {
						return getter(structAddr + offset)
					}
				}
				return interfacePtrGetter(f)
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
				return (**int)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint:
			return func(structAddr uintptr) interface{} {
				return (**uint)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int64:
			return func(structAddr uintptr) interface{} {
				return (**int64)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int32:
			return func(structAddr uintptr) interface{} {
				return (**int32)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int16:
			return func(structAddr uintptr) interface{} {
				return (**int16)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int8:
			return func(structAddr uintptr) interface{} {
				return (**int8)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint64:
			return func(structAddr uintptr) interface{} {
				return (**uint64)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint32:
			return func(structAddr uintptr) interface{} {
				return (**uint32)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint16:
			return func(structAddr uintptr) interface{} {
				return (**uint16)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint8:
			return func(structAddr uintptr) interface{} {
				return (**uint8)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.String:
			return func(structAddr uintptr) interface{} {
				return (**string)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Float64:
			return func(structAddr uintptr) interface{} {
				return (**float64)(unsafe.Pointer(structAddr + offset))
			}

		case reflect.Float32:
			return func(structAddr uintptr) interface{} {
				return (**float32)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Bool:
			return func(structAddr uintptr) interface{} {
				return (**bool)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Slice:
			switch f.field.Type.Elem().Elem().Kind() {
			case reflect.Int:
				return func(structAddr uintptr) interface{} {
					return (**[]int)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Uint:
				return func(structAddr uintptr) interface{} {
					return (**[]uint)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Int64:
				return func(structAddr uintptr) interface{} {
					return (**[]int64)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Int32:
				return func(structAddr uintptr) interface{} {
					return (**[]int32)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Int16:
				return func(structAddr uintptr) interface{} {
					return (**[]int16)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Int8:
				return func(structAddr uintptr) interface{} {
					return (**[]int8)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Uint64:
				return func(structAddr uintptr) interface{} {
					return (**[]uint64)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Uint32:
				return func(structAddr uintptr) interface{} {
					return (**[]uint32)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Uint16:
				return func(structAddr uintptr) interface{} {
					return (**[]uint16)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Uint8:
				return func(structAddr uintptr) interface{} {
					return (**[]uint8)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.String:
				return func(structAddr uintptr) interface{} {
					return (**[]string)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Float64:
				return func(structAddr uintptr) interface{} {
					return (**[]float64)(unsafe.Pointer(structAddr + offset))
				}

			case reflect.Float32:
				return func(structAddr uintptr) interface{} {
					return (**[]float32)(unsafe.Pointer(structAddr + offset))
				}
			case reflect.Bool:
				return func(structAddr uintptr) interface{} {
					return (**[]bool)(unsafe.Pointer(structAddr + offset))
				}
			}
		}
	case reflect.Slice:
		switch f.field.Type.Elem().Kind() {
		case reflect.Int:
			return func(structAddr uintptr) interface{} {
				return (*[]int)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint:
			return func(structAddr uintptr) interface{} {
				return (*[]uint)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int64:
			return func(structAddr uintptr) interface{} {
				return (*[]int64)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int32:
			return func(structAddr uintptr) interface{} {
				return (*[]int32)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int16:
			return func(structAddr uintptr) interface{} {
				return (*[]int16)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Int8:
			return func(structAddr uintptr) interface{} {
				return (*[]int8)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint64:
			return func(structAddr uintptr) interface{} {
				return (*[]uint64)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint32:
			return func(structAddr uintptr) interface{} {
				return (*[]uint32)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint16:
			return func(structAddr uintptr) interface{} {
				return (*[]uint16)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Uint8:
			return func(structAddr uintptr) interface{} {
				return (*[]uint8)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.String:
			return func(structAddr uintptr) interface{} {
				return (*[]string)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Float64:
			return func(structAddr uintptr) interface{} {
				return (*[]float64)(unsafe.Pointer(structAddr + offset))
			}

		case reflect.Float32:
			return func(structAddr uintptr) interface{} {
				return (*[]float32)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Bool:
			return func(structAddr uintptr) interface{} {
				return (*[]bool)(unsafe.Pointer(structAddr + offset))
			}
		case reflect.Ptr:
			return func(structAddr uintptr) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
				return fieldValue.Interface()
			}
		case reflect.Struct:
			return func(structAddr uintptr) interface{} {
				fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+offset))
				return fieldValue.Interface()
			}
		}
	}
	return func(structAddr uintptr) interface{} {
		fieldValue := reflect.NewAt(f.field.Type, unsafe.Pointer(structAddr+f.field.Offset))
		return fieldValue.Elem().Elem().Interface()
	}
}

func interfacePtrGetter(field *Field) func(structAddr uintptr) interface{} {
	return func(structAddr uintptr) interface{} {
		fieldValue := reflect.NewAt(field.field.Type, unsafe.Pointer(structAddr+field.field.Offset))
		return fieldValue.Interface()
	}
}
