package converter

import (
	"reflect"
	"unsafe"
)

func ToInt32(from reflect.Type) (func(pointer unsafe.Pointer) int32, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) int32 {
		return int32(inter(pointer))
	}, nil
}

func ToUint32(from reflect.Type) (func(pointer unsafe.Pointer) uint32, error) {
	inter, err := ToInt(from)
	if err != nil {
		return nil, err
	}

	return func(pointer unsafe.Pointer) uint32 {
		return uint32(inter(pointer))
	}, nil
}
