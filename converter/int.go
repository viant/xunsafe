package converter

import (
	"github.com/viant/xunsafe"
	"reflect"
	"strconv"
	"unsafe"
)

func ToInt(from reflect.Type) (func(pointer unsafe.Pointer) int, error) {
	var isPtr bool
	if from.Kind() == reflect.Ptr {
		isPtr = true
		from = from.Elem()
	}

	inter, ok := toInt(from, isPtr)
	if !ok {
		return nil, UnsupportedConversion(from, reflect.TypeOf(0))
	}

	return func(pointer unsafe.Pointer) int {
		if pointer == nil {
			return 0
		}

		return inter(pointer)
	}, nil
}

func ToUint(from reflect.Type) (func(unsafe.Pointer) uint, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) uint {
		return uint(inter(pointer))
	}, nil
}

func toInt(from reflect.Type, isPtr bool) (func(pointer unsafe.Pointer) int, bool) {
	switch from.Kind() {
	case reflect.Int:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*int)(pointer)
				if intPtr == nil {
					return 0
				}

				return *intPtr
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			value := *(*int)(pointer)
			return value
		}, true
	case reflect.Int8:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*int8)(pointer)
				if intPtr == nil {
					return 0
				}
				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*int8)(pointer))
		}, true
	case reflect.Int16:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*int16)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*int16)(pointer))
		}, true
	case reflect.Int32:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*int32)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*int32)(pointer))
		}, true
	case reflect.Int64:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*int64)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*int64)(pointer))
		}, true

	case reflect.Uint:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*uint)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*uint)(pointer))
		}, true
	case reflect.Uint8:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*uint8)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*uint8)(pointer))
		}, true
	case reflect.Uint16:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*uint16)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true

		}
		return func(pointer unsafe.Pointer) int {
			return int(*(*uint16)(pointer))
		}, true
	case reflect.Uint32:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*uint32)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*uint32)(pointer))
		}, true
	case reflect.Uint64:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				intPtr := (*uint64)(pointer)
				if intPtr == nil {
					return 0
				}

				return int(*intPtr)
			}, true

		}
		return func(pointer unsafe.Pointer) int {
			return int(*(*uint64)(pointer))
		}, true

	case reflect.Float64:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				floatPtr := (*float64)(pointer)
				if floatPtr == nil {
					return 0
				}

				return int(*floatPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*float64)(pointer))
		}, true

	case reflect.Float32:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				floatPtr := (*float32)(pointer)
				if floatPtr == nil {
					return 0
				}

				return int(*floatPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			return int(*(*float32)(pointer))
		}, true

	case reflect.String:
		if isPtr {
			return func(pointer unsafe.Pointer) int {
				strPtr := (*string)(pointer)
				if strPtr == nil {
					return 0
				}

				asInt, _ := strconv.Atoi(*strPtr)
				return asInt
			}, true
		}

		return func(pointer unsafe.Pointer) int {
			asInt, _ := strconv.Atoi(*(*string)(pointer))
			return asInt

		}, true

	default:
		return nil, false
	}
}

func ToIntPtr(from reflect.Type) (func(pointer unsafe.Pointer) unsafe.Pointer, error) {
	var isPtr bool
	if from.Kind() == reflect.Ptr {
		isPtr = true
		from = from.Elem()
	}

	inter, ok := toIntPtr(from, isPtr)
	if !ok {
		return nil, UnsupportedConversion(from, reflect.TypeOf(0))
	}

	return func(pointer unsafe.Pointer) unsafe.Pointer {
		if pointer == nil {
			return nil
		}

		return inter(pointer)
	}, nil
}

func toIntPtr(from reflect.Type, isPtr bool) (func(pointer unsafe.Pointer) unsafe.Pointer, bool) {
	switch from.Kind() {
	case reflect.Int:
		if isPtr {
			return func(pointer unsafe.Pointer) unsafe.Pointer {
				intPtr := (*int)(pointer)
				if intPtr == nil {
					return nil
				}

				return xunsafe.DerefPointer(pointer)
			}, true
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			return pointer
		}, true
	}

	inter, ok := toInt(from, isPtr)
	if !ok {
		return nil, false
	}

	return func(pointer unsafe.Pointer) unsafe.Pointer {
		anInt := inter(pointer)
		return unsafe.Pointer(&anInt)
	}, true
}
