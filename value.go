package xunsafe

import (
	"reflect"
	"unsafe"
)

//ValuePointer cast generic value ptr to unsafe pointer
type ValuePointer func(interface{}) unsafe.Pointer

//ValuePointerForType returns function casting interface to unsafe.Pointer for supplied type
func ValuePointerForType(t reflect.Type) ValuePointer {
	switch t.Kind() {
	case reflect.Int:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*int))
		}
	case reflect.Uint:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*uint))
		}
	case reflect.Int64:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*int64))
		}
	case reflect.Int32:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*int32))
		}
	case reflect.Int16:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*int16))
		}
	case reflect.Int8:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*int8))
		}
	case reflect.Uint64:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*uint64))
		}
	case reflect.Uint32:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*uint32))
		}
	case reflect.Uint16:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*uint16))
		}
	case reflect.Uint8:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*uint8))
		}
	case reflect.String:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*string))
		}
	case reflect.Float64:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*float64))
		}
	case reflect.Float32:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*float32))
		}
	case reflect.Bool:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*bool))
		}
	case reflect.Interface:
		return func(ptr interface{}) unsafe.Pointer {
			return unsafe.Pointer(ptr.(*interface{}))
		}
	case reflect.Struct:
		return func(ptr interface{}) unsafe.Pointer {
			fieldValue := reflect.ValueOf(ptr)
			if fieldValue.IsNil() {
				aValue := reflect.New(t)
				fieldValue.Elem().Set(aValue)
			}
			return unsafe.Pointer(fieldValue.Elem().UnsafeAddr())
		}
	case reflect.Ptr:
		switch t.Elem().Kind() {
		case reflect.Int:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**int))
			}
		case reflect.Uint:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**uint))
			}
		case reflect.Int64:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**int64))
			}
		case reflect.Int32:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**int32))
			}
		case reflect.Int16:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**int16))
			}
		case reflect.Int8:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**int8))
			}
		case reflect.Uint64:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**uint64))
			}
		case reflect.Uint32:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**uint32))
			}
		case reflect.Uint16:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**uint16))
			}
		case reflect.Uint8:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**uint8))
			}
		case reflect.String:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**string))
			}
		case reflect.Float64:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**float64))
			}
		case reflect.Float32:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**float32))
			}
		case reflect.Bool:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**bool))
			}
		case reflect.Struct:
			return func(ptr interface{}) unsafe.Pointer {
				fieldValue := reflect.ValueOf(ptr)
				if fieldValue.IsNil() {
					aValue := reflect.New(t)
					fieldValue.Elem().Set(aValue)
				}
				return unsafe.Pointer(fieldValue.Elem().UnsafeAddr())
			}
		case reflect.Interface:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(**interface{}))
			}
		}
	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Int:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]int))
			}
		case reflect.Uint:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]uint))
			}
		case reflect.Int64:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]int64))
			}
		case reflect.Int32:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]int32))
			}
		case reflect.Int16:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]int16))
			}
		case reflect.Int8:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]int8))
			}
		case reflect.Uint64:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]uint64))
			}
		case reflect.Uint32:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]uint32))
			}
		case reflect.Uint16:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]uint16))
			}
		case reflect.Uint8:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]uint8))
			}
		case reflect.String:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]string))
			}
		case reflect.Float64:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]float64))
			}
		case reflect.Float32:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]float32))
			}
		case reflect.Bool:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]bool))
			}
		case reflect.Interface:
			return func(ptr interface{}) unsafe.Pointer {
				return unsafe.Pointer(ptr.(*[]interface{}))
			}
		case reflect.Struct:
			return func(ptr interface{}) unsafe.Pointer {
				fieldValue := reflect.ValueOf(ptr)
				if fieldValue.IsNil() {
					aValue := reflect.MakeSlice(t, 0, 0)
					fieldValue.Elem().Set(aValue)
				}
				return unsafe.Pointer(fieldValue.Elem().UnsafeAddr())
			}
		case reflect.Ptr:
			switch t.Elem().Elem().Kind() {
			case reflect.Int:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*int))
				}
			case reflect.Uint:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*uint))
				}
			case reflect.Int64:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*int64))
				}
			case reflect.Int32:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*int32))
				}
			case reflect.Int16:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*int16))
				}
			case reflect.Int8:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*int8))
				}
			case reflect.Uint64:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*uint64))
				}
			case reflect.Uint32:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*uint32))
				}
			case reflect.Uint16:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*uint16))
				}
			case reflect.Uint8:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*uint8))
				}
			case reflect.String:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*string))
				}
			case reflect.Float64:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*float64))
				}
			case reflect.Float32:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*float32))
				}
			case reflect.Bool:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*bool))
				}
			case reflect.Interface:
				return func(ptr interface{}) unsafe.Pointer {
					return unsafe.Pointer(ptr.(*[]*interface{}))
				}
			case reflect.Struct:
				return func(ptr interface{}) unsafe.Pointer {
					fieldValue := reflect.ValueOf(ptr)
					if fieldValue.IsNil() {
						aValue := reflect.New(t)
						fieldValue.Elem().Set(aValue)
					}
					return unsafe.Pointer(fieldValue.Elem().UnsafeAddr())
				}
			}
		}
	}
	return nil
}
