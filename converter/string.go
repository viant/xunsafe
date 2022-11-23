package converter

import (
	"github.com/viant/xunsafe"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

var (
	timeType    = reflect.TypeOf(time.Time{})
	intType     = reflect.TypeOf(1)
	stringType  = reflect.TypeOf("1")
	float64Type = reflect.TypeOf(1.1)
)

func ToString(from reflect.Type) (func(pointer unsafe.Pointer) string, error) {
	var wasPtr bool
	if from.Kind() == reflect.Ptr {
		wasPtr = true
		from = from.Elem()
	}

	result, ok := toString(from, wasPtr)
	if !ok {
		return nil, UnsupportedConversion(from, reflect.TypeOf(""))
	}

	return result, nil
}

func toString(from reflect.Type, wasPtr bool) (func(pointer unsafe.Pointer) string, bool) {
	switch from.Kind() {
	case reflect.Float64, reflect.Float32:
		floater, ok := toFloat64(from, wasPtr)
		if !ok {
			return nil, false
		}

		return func(pointer unsafe.Pointer) string {
			return strconv.FormatFloat(floater(pointer), 'f', -1, 64)
		}, true

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		inter, ok := toInt(from, wasPtr)
		if !ok {
			return nil, false
		}

		return func(pointer unsafe.Pointer) string {
			return strconv.Itoa(inter(pointer))
		}, true

	case reflect.String:
		if wasPtr {
			return func(pointer unsafe.Pointer) string {
				strPtr := (*string)(pointer)
				if strPtr == nil {
					return ""
				}

				return *strPtr
			}, true
		}

		return func(pointer unsafe.Pointer) string {
			return *(*string)(pointer)
		}, true

	case reflect.Struct:
		if from == timeType {
			timer, ok := toTime(from, wasPtr)
			if !ok {
				return nil, false
			}

			return func(pointer unsafe.Pointer) string {
				return timer(pointer).Format(time.RFC3339)
			}, true
		}
	}

	return nil, false
}

func ToStringPtr(from reflect.Type) (func(pointer unsafe.Pointer) unsafe.Pointer, error) {
	wasPtr := false
	if from.Kind() == reflect.Ptr {
		wasPtr = true
		from = from.Elem()
	}

	stringifierPtr, ok := toStringPtr(from, wasPtr)
	if !ok {
		return nil, UnsupportedConversion(from, stringType)
	}

	return stringifierPtr, nil
}

func toStringPtr(from reflect.Type, wasPtr bool) (func(pointer unsafe.Pointer) unsafe.Pointer, bool) {
	switch from.Kind() {
	case reflect.Float64, reflect.Float32:
		floater, ok := toFloat64(from, wasPtr)
		if !ok {
			return nil, false
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			aFloat := strconv.FormatFloat(floater(pointer), 'f', -1, 64)
			return unsafe.Pointer(&aFloat)
		}, true

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		inter, ok := toInt(from, wasPtr)
		if !ok {
			return nil, false
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			aString := strconv.Itoa(inter(pointer))
			return unsafe.Pointer(&aString)
		}, true

	case reflect.String:
		if wasPtr {
			return func(pointer unsafe.Pointer) unsafe.Pointer {
				strPtr := (*string)(pointer)
				if strPtr == nil {
					return nil
				}

				return xunsafe.DerefPointer(pointer)
			}, true
		}

		return func(pointer unsafe.Pointer) unsafe.Pointer {
			return pointer
		}, true
	}

	stringifier, ok := toString(from, wasPtr)
	if !ok {
		return nil, false
	}

	return func(pointer unsafe.Pointer) unsafe.Pointer {
		return xunsafe.AsPointer(stringifier(pointer))
	}, true
}
