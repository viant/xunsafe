package converter

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/xunsafe"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnifier(t *testing.T) {
	testCases := []struct {
		description string
		fromType    reflect.Type
		toType      reflect.Type
		actual      interface{}
		expected    interface{}
		interfacer  func(pointer unsafe.Pointer) interface{}
	}{
		{
			description: "ref pointer value",
			fromType:    reflect.PtrTo(reflect.TypeOf(0)),
			toType:      reflect.TypeOf(0),
			actual:      intPtr(10),
			expected:    10,
			interfacer: func(pointer unsafe.Pointer) interface{} {
				return *(*int)(pointer)
			},
		},
		{
			description: "deref pointer value",
			fromType:    reflect.TypeOf(0),
			toType:      reflect.PtrTo(reflect.TypeOf(0)),
			actual:      10,
			expected:    intPtr(10),
			interfacer: func(pointer unsafe.Pointer) interface{} {
				return *(**int)(pointer)
			},
		},
		{
			description: "string -> string ptr",
			fromType:    reflect.TypeOf(""),
			toType:      reflect.PtrTo(reflect.TypeOf("")),
			actual:      "12345",
			expected:    stringPtr("12345"),
			interfacer: func(pointer unsafe.Pointer) interface{} {
				return *(**string)(pointer)
			},
		},
		{
			description: "int64 -> *uint64",
			fromType:    reflect.TypeOf(int64(0)),
			toType:      reflect.PtrTo(reflect.TypeOf(uint64(10))),
			actual:      int64(64),
			expected:    uint64Ptr(64),
			interfacer: func(pointer unsafe.Pointer) interface{} {
				return *(**uint64)(pointer)
			},
		},
	}

	for _, testCase := range testCases[len(testCases)-1:] {
		//for _, testCase := range testCases {
		unified, err := Unify(testCase.toType, testCase.fromType)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}

		actPtr := xunsafe.AsPointer(testCase.actual)
		if unified.Y != nil {
			actPtr, err = unified.Y(actPtr)
			if !assert.Nil(t, err, testCase.description) {
				continue
			}
		}

		actualValue := testCase.interfacer(actPtr)
		assert.EqualValues(t, testCase.expected, actualValue, testCase.description)
	}
}

func uint64Ptr(i int) *uint64 {
	asUint := uint64(i)
	return &asUint
}

func intPtr(i int) *int {
	return &i
}

func stringPtr(i string) *string {
	return &i
}
