package converter

import (
	"reflect"
	"time"
)

type Unified struct {
	X     UnifyFn
	Y     UnifyFn
	RType reflect.Type
}

var timeType = reflect.TypeOf(time.Time{})
var intType = reflect.TypeOf(0)
var float64Type = reflect.TypeOf(0.0)

func NormalizeAndUnify(x, y reflect.Type) (*Unified, error) {
	if x == y {
		return &Unified{RType: x}, nil
	}

	toType := x
	if hasPrecedence(y, x) {
		toType = y
	}

	resultType := NormalizeType(toType)
	return newUnifier(x, y, resultType)
}

func Unify(x, y reflect.Type) (*Unified, error) {
	return newUnifier(x, y, x)
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
