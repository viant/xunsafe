package converter

import (
	"github.com/viant/xunsafe"
	"reflect"
	"unsafe"
)

func ToFloat64(from reflect.Type) (func(pointer unsafe.Pointer) float64, error) {
	var wasPtr bool
	if from.Kind() == reflect.Ptr {
		wasPtr = true
		from = from.Elem()
	}

	result, ok := toFloat64(from, wasPtr)
	if !ok {
		return nil, UnsupportedConversion(from, reflect.TypeOf(0.0))
	}

	return result, nil
}

func toFloat64(from reflect.Type, wasPtr bool) (func(pointer unsafe.Pointer) float64, bool) {
	switch from.Kind() {
	case reflect.Float64:
		if wasPtr {
			return func(pointer unsafe.Pointer) float64 {
				floatPtr := (*float64)(pointer)
				if floatPtr == nil {
					return 0
				}
				return *floatPtr
			}, true
		}

		return func(pointer unsafe.Pointer) float64 {
			return *(*float64)(pointer)
		}, true

	case reflect.Float32:
		if wasPtr {
			return func(pointer unsafe.Pointer) float64 {
				floatPtr := (*float32)(pointer)
				if floatPtr == nil {
					return 0
				}
				return float64(*floatPtr)
			}, true
		}

		return func(pointer unsafe.Pointer) float64 {
			return float64(*(*float32)(pointer))
		}, true

	default:
		toInter, ok := toInt(from, wasPtr)
		if !ok {
			return nil, false
		}

		return func(pointer unsafe.Pointer) float64 {
			return float64(toInter(pointer))
		}, true
	}
}

func ToFloat64Ptr(from reflect.Type) (func(pointer unsafe.Pointer) unsafe.Pointer, error) {
	wasPtr := false
	if from.Kind() == reflect.Ptr {
		wasPtr = true
		from = from.Elem()
	}

	floater, ok := toFloat64Ptr(from, wasPtr)
	if !ok {
		return nil, UnsupportedConversion(from, float64Type)
	}

	return func(pointer unsafe.Pointer) unsafe.Pointer {
		if pointer == nil {
			return nil
		}

		return floater(pointer)
	}, nil
}

func toFloat64Ptr(from reflect.Type, wasPtr bool) (func(pointer unsafe.Pointer) unsafe.Pointer, bool) {
	switch from.Kind() {
	case reflect.Float64:
		if wasPtr {
			return func(pointer unsafe.Pointer) unsafe.Pointer {
				floatPtr := (*float64)(pointer)
				if floatPtr == nil {
					return nil
				}

				return xunsafe.DerefPointer(pointer)
			}, true
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			return pointer
		}, true

	case reflect.Float32:
		if wasPtr {
			return func(pointer unsafe.Pointer) unsafe.Pointer {
				floatPtr := (*float32)(pointer)
				if floatPtr == nil {
					return nil
				}
				aFloat := float64(*floatPtr)

				return unsafe.Pointer(&aFloat)
			}, true
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			aFloat := float64(*(*float32)(pointer))
			return unsafe.Pointer(&aFloat)
		}, true

	default:
		toInter, ok := toInt(from, wasPtr)
		if !ok {
			return nil, false
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			inter := toInter(pointer)
			v := float64(inter)
			return unsafe.Pointer(&v)
		}, true
	}
}

func ToFloat32(from reflect.Type) (func(pointer unsafe.Pointer) float32, error) {
	toFloat, err := ToFloat64(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) float32 {
		return float32(toFloat(pointer))
	}, nil
}
