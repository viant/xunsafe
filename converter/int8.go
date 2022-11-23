package converter

import (
	"reflect"
	"unsafe"
)

func ToInt8(from reflect.Type) (func(pointer unsafe.Pointer) int8, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) int8 {
		return int8(inter(pointer))
	}, nil
}

func ToUint8(from reflect.Type) (func(pointer unsafe.Pointer) uint8, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) uint8 {
		return uint8(inter(pointer))
	}, nil
}
