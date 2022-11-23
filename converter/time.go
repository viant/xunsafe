package converter

import (
	"github.com/viant/xunsafe"
	"reflect"
	"time"
	"unsafe"
)

func ToTime(from reflect.Type) (func(pointer unsafe.Pointer) time.Time, error) {
	var wasPtr bool
	if from.Kind() == reflect.Ptr {
		wasPtr = true
		from = from.Elem()
	}

	result, ok := toTime(from, wasPtr)
	if !ok {
		return nil, UnsupportedConversion(from, timeType)
	}

	return result, nil
}

func toTime(from reflect.Type, wasPtr bool) (func(pointer unsafe.Pointer) time.Time, bool) {
	switch from.Kind() {
	case reflect.Struct:
		if from == timeType {
			if wasPtr {
				return func(pointer unsafe.Pointer) time.Time {
					timePtr := (*time.Time)(pointer)
					if timePtr == nil {
						return time.Time{}
					}

					return *timePtr
				}, true
			}

			return func(pointer unsafe.Pointer) time.Time {
				return *(*time.Time)(pointer)
			}, true
		}

	case reflect.String:
		if wasPtr {
			return func(pointer unsafe.Pointer) time.Time {
				strPtr := (*string)(pointer)
				if strPtr == nil {
					return time.Time{}
				}

				result, _ := time.Parse(time.RFC3339, *strPtr)
				return result
			}, true
		}
	}

	return nil, false
}

func ToTimePtr(from reflect.Type) (func(pointer unsafe.Pointer) unsafe.Pointer, error) {
	var wasPtr bool
	if from.Kind() == reflect.Ptr {
		wasPtr = true
		from = from.Elem()
	}

	result, ok := toTimePtr(from, wasPtr)
	if !ok {
		return nil, UnsupportedConversion(from, timeType)
	}

	return result, nil
}

func toTimePtr(from reflect.Type, wasPtr bool) (func(pointer unsafe.Pointer) unsafe.Pointer, bool) {
	switch from.Kind() {
	case reflect.Struct:
		if from == timeType {
			if wasPtr {
				return func(pointer unsafe.Pointer) unsafe.Pointer {
					timePtr := (*time.Time)(pointer)

					if timePtr == nil {
						return nil
					}

					return xunsafe.DerefPointer(pointer)
				}, true
			}

			return func(pointer unsafe.Pointer) unsafe.Pointer {
				return pointer
			}, true
		}
	}

	timer, ok := toTime(from, wasPtr)
	if !ok {
		return nil, false
	}

	return func(pointer unsafe.Pointer) unsafe.Pointer {
		aTime := timer(pointer)
		return unsafe.Pointer(&aTime)
	}, true
}
