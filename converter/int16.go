package converter

import (
	"reflect"
	"unsafe"
)

func ToInt16(from reflect.Type) (func(pointer unsafe.Pointer) int16, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) int16 {
		return int16(inter(pointer))
	}, nil
}

func ToUint16(from reflect.Type) (func(pointer unsafe.Pointer) uint16, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) uint16 {
		return uint16(inter(pointer))
	}, nil
}
