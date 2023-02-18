package converter

import (
	"github.com/viant/xunsafe"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

type UnifyFn func(pointer unsafe.Pointer) (unsafe.Pointer, error)

func New(to, from reflect.Type) (*Unified, error) {
	return newUnifier(to, from, to)
}

func newUnifier(to reflect.Type, from reflect.Type, resultType reflect.Type) (*Unified, error) {
	toUnifyFn, err := newUnifyFn(to, resultType)
	if err != nil {
		return nil, err
	}

	fromUnifyFn, err := newUnifyFn(from, resultType)
	if err != nil {
		return nil, err
	}

	return &Unified{
		X:     toUnifyFn,
		Y:     fromUnifyFn,
		RType: resultType,
	}, nil
}

func newUnifyFn(x reflect.Type, to reflect.Type) (UnifyFn, error) {
	if x == to {
		return nil, nil
	}

	originalFrom := x
	originalResult := to
	x, fromPtrCounter := deref(x)
	to, resultTypeCounter := deref(to)

	switch x.Kind() {
	case reflect.Uint, reflect.Uint64:
		switch to.Kind() {
		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint, reflect.Uint64:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*uint)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint)(pointer))
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Int, reflect.Int64:
		switch to.Kind() {
		case reflect.Int, reflect.Int64:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*int)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*int)(pointer)
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*int)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Uint8:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint8:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)
		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*uint8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint8)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*uint8)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Int8:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int8:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*int8)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int8)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*int8)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Uint16:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*uint16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint16)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*uint16)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Int16:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*int16)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int16)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*int16)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Uint32:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*uint32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*uint32)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*uint32)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Int32:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*int32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*int32)(pointer))
				aString := strconv.Itoa(anInt)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*int32)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Float64:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float32(*(*float64)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float64:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aFloat := *(*float64)(pointer)
				aString := strconv.FormatFloat(aFloat, 'f', -1, 32)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*float64)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Float32:
		switch to.Kind() {
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint8(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int8(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint16(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int16(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := uint32(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := int32(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Float32:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.Float64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := float64(*(*float32)(pointer))
				return unsafe.Pointer(&anInt), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aFloat := float64(*(*float32)(pointer))
				aString := strconv.FormatFloat(aFloat, 'f', -1, 32)
				return unsafe.Pointer(&aString), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				anInt := *(*float32)(pointer)
				aBool := anInt != 0
				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.String:
		switch to.Kind() {
		case reflect.String:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)
		case reflect.Uint, reflect.Uint64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				anUint := uint(anInt)

				return unsafe.Pointer(&anUint), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int, reflect.Int64:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				return unsafe.Pointer(&anInt), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				anUint := uint8(anInt)

				return unsafe.Pointer(&anUint), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				casted := int8(anInt)
				return unsafe.Pointer(&casted), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				anUint := uint16(anInt)

				return unsafe.Pointer(&anUint), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				casted := int16(anInt)
				return unsafe.Pointer(&casted), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				anUint := uint32(anInt)

				return unsafe.Pointer(&anUint), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				anInt, err := strconv.Atoi(aString)
				if err != nil {
					return nil, err
				}

				casted := int32(anInt)
				return unsafe.Pointer(&casted), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aString := *(*string)(pointer)
				aBool, err := strconv.ParseBool(aString)
				if err != nil {
					return nil, err
				}

				return unsafe.Pointer(&aBool), nil
			}

			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Struct:
			if to == timeType {
				var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
					aString := *(*string)(pointer)
					aTime, err := time.Parse(time.RFC3339, aString)
					if err != nil {
						return nil, err
					}

					return unsafe.Pointer(&aTime), nil
				}

				return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
			}

		}

	case reflect.Bool:
		switch to.Kind() {
		case reflect.Int64, reflect.Int:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return intOnePtr, nil
				}

				return intZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint64, reflect.Uint:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return uintOnePtr, nil
				}

				return uintZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return int32OnePtr, nil
				}

				return int32ZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint32:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return uint32OnePtr, nil
				}

				return uint32ZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return uint16OnePtr, nil
				}

				return uint16ZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int16:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return int16OnePtr, nil
				}

				return int16ZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Int8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return int8OnePtr, nil
				}

				return int8ZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Uint8:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				if *(*bool)(pointer) {
					return uint8OnePtr, nil
				}

				return uint8ZeroPtr, nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)

		case reflect.Bool:
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, nil)

		case reflect.String:
			var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
				aBool := *(*bool)(pointer)
				asString := strconv.FormatBool(aBool)
				return unsafe.Pointer(&asString), nil
			}
			return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
		}

	case reflect.Struct:
		switch to.Kind() {
		case reflect.String:
			if x == timeType {
				var fn UnifyFn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
					aTime := *(*time.Time)(pointer)
					aString := aTime.Format(time.RFC3339)
					return unsafe.Pointer(&aString), nil
				}

				return wrapWithDeref(x, fromPtrCounter, resultTypeCounter, fn)
			}
		}
	}

	if to.Kind() == reflect.Bool {
		return func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
			if pointer == nil {
				return falsePtr, nil
			}

			if xunsafe.DerefPointer(pointer) == nil {
				return falsePtr, nil
			}

			return truePtr, nil
		}, nil
	}

	return nil, UnsupportedConversion(originalFrom, originalResult)
}

func wrapWithDeref(from reflect.Type, fromCounter int, resultTypeCounter int, fn UnifyFn) (UnifyFn, error) {
	if fn == nil {
		fn = func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
			return pointer, nil
		}
	}

	if fromCounter == 0 && resultTypeCounter == 0 {
		return fn, nil
	}

	//if IsPrimitive(from) {
	//	fromCounter--
	//}

	//if fromCounter <= 0 && resultTypeCounter <= 0 {
	//	return fn, nil
	//}

	return func(pointer unsafe.Pointer) (unsafe.Pointer, error) {
		for i := 0; i < fromCounter; i++ {
			pointer = xunsafe.DerefPointer(pointer)
			if pointer == nil {
				return nil, nil
			}
		}

		pointer, err := fn(pointer)
		if err != nil {
			return nil, err
		}

		for i := 0; i < resultTypeCounter; i++ {
			pointer = xunsafe.RefPointer(pointer)
		}

		return pointer, nil
	}, nil
}

func IsPrimitive(from reflect.Type) bool {
	switch from.Kind() {
	case reflect.Ptr, reflect.Struct, reflect.Slice, reflect.Map, reflect.Interface:
		return false
	default:
		return true
	}
}

func deref(rType reflect.Type) (reflect.Type, int) {
	i := 0
	for rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
		i++
	}

	return rType, i
}
