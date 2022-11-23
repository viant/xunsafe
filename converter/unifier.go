package converter

import (
	"reflect"
	"unsafe"
)

type Unified struct {
	X     func(pointer unsafe.Pointer) unsafe.Pointer
	Y     func(pointer unsafe.Pointer) unsafe.Pointer
	RType reflect.Type
}

func NormalizeAndUnify(x, y reflect.Type) (*Unified, error) {
	if x == y {
		return &Unified{RType: x}, nil
	}

	toType := x
	if hasPrecedence(y, x) {
		toType = y
	}

	resultType := NormalizeType(toType)
	return unify(x, y, resultType)
}

func Unify(x, y reflect.Type) (*Unified, error) {
	return unify(x, y, x)
}

func unify(x reflect.Type, y reflect.Type, resultType reflect.Type) (*Unified, error) {
	xNormalizer, err := normalizeTo(x, resultType)
	if err != nil {
		return nil, err
	}

	yNormalizer, err := normalizeTo(y, resultType)
	if err != nil {
		return nil, err
	}

	return &Unified{
		X:     xNormalizer,
		Y:     yNormalizer,
		RType: resultType,
	}, nil
}

func normalizeTo(from reflect.Type, to reflect.Type) (func(pointer unsafe.Pointer) unsafe.Pointer, error) {
	if from == to {
		return func(pointer unsafe.Pointer) unsafe.Pointer {
			return pointer
		}, nil
	}

	switch to.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return ToIntPtr(from)
	case reflect.Float64, reflect.Float32:
		return ToFloat64Ptr(from)
	case reflect.String:
		return ToStringPtr(from)
	default:
		if to == timeType {
			return ToTimePtr(to)
		}

		return nil, nil
	}
}

func NormalizeType(from reflect.Type) reflect.Type {
	if from == nil {
		return nil
	}

	if from.Kind() == reflect.Ptr {
		from = from.Elem()
	}

	switch from.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return intType
	case reflect.Int:
		return intType
	case reflect.Float64:
		return float64Type
	case reflect.Float32:
		return float64Type
	}
	return from
}

func hasPrecedence(over reflect.Type, rType reflect.Type) bool {
	switch over.Kind() {
	case reflect.Float64, reflect.Float32:
		switch rType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true
		case reflect.Float32:
			return true
		case reflect.Float64:
			return false
		}

	case reflect.String:
		return true
	}

	return false
}
