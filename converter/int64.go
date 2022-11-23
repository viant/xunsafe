package converter

import (
	"reflect"
	"unsafe"
)

func ToInt64(from reflect.Type) (func(pointer unsafe.Pointer) int64, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) int64 {
		return int64(inter(pointer))
	}, nil
}

func ToUint64(from reflect.Type) (func(pointer unsafe.Pointer) uint64, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) uint64 {
		return uint64(inter(pointer))
	}, nil
}
