package xunsafe_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/xunsafe"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type Struct1 struct {
	I   int
	I64 int64
	I32 int32
	I16 int16
	I8  int8

	UI   uint
	UI64 uint64
	UI32 uint32
	UI16 uint16
	UI8  uint8

	F64 float64
	F32 float32

	B    bool
	S    string
	Bs   []byte
	T    time.Time
	Bars []Bar
	Foo  *Foo
	F2   Foo
	F3   *Foo
	Fn   func()
}

type Bar struct {
	ID int
}

type Foo struct {
	ID int
}

func NewStruct1() *Struct1 {
	return &Struct1{
		I:   1,
		I64: 2,
		I32: 3,
		I16: 4,
		I8:  5,

		UI:   6,
		UI64: 7,
		UI32: 8,
		UI16: 9,
		UI8:  10,

		F64: 11,
		F32: 12,

		B:  true,
		S:  "test",
		Bs: []byte("abc"),
		T:  time.Now(),
		Bars: []Bar{
			{
				ID: 1,
			},
		},
		Foo: &Foo{ID: 12},
		F2:  Foo{ID: 30},
		Fn: func() {
			fmt.Println("123")
		},
	}
}

func TestField_Mutators(t *testing.T) {
	structArchetype := NewStruct1()
	structInstance1 := NewStruct1()
	structInstance2 := NewStruct1()

	var testCases = []struct {
		description string
		expect      interface{}
		actual      func() interface{}
		name        string
	}{

		{
			description: "func",
			expect:      structArchetype.Fn,
			name:        "Fn",
		},
		{
			description: "*Foo",
			expect:      structArchetype.Foo,
			name:        "Foo",
		},
		{
			description: "Foo",
			expect:      structArchetype.F2,
			name:        "F2",
		},
		{
			description: "nil *Foo",
			expect:      structArchetype.F3,
			name:        "F3",
		},
		{
			description: "int",
			expect:      100 + structArchetype.I,
			name:        "I",
		},
		{
			description: "int",
			expect:      100 + structArchetype.I,
			name:        "I",
		},
		{
			description: "int64",
			expect:      100 + structArchetype.I64,
			name:        "I64",
		},
		{
			description: "int32",
			expect:      100 + structArchetype.I32,
			name:        "I32",
		},
		{
			description: "int16",
			expect:      100 + structArchetype.I16,
			name:        "I16",
		},
		{
			description: "int8",
			expect:      100 + structArchetype.I8,
			name:        "I8",
		},
		{
			description: "uint",
			expect:      100 + structArchetype.UI,
			name:        "UI",
		},
		{
			description: "uint64",
			expect:      100 + structArchetype.UI64,
			name:        "UI64",
		},
		{
			description: "uint32",
			expect:      100 + structArchetype.UI32,
			name:        "UI32",
		},
		{
			description: "uint16",
			expect:      100 + structArchetype.UI16,
			name:        "UI16",
		},
		{
			description: "uint8",
			expect:      100 + structArchetype.UI8,
			name:        "UI8",
		},
		{
			description: "string",
			expect:      "pre" + structArchetype.S,
			name:        "S",
		},
		{
			description: "bool",
			expect:      structArchetype.B,
			name:        "B",
		},
		{
			description: "float64",
			expect:      100 + structArchetype.F64,
			name:        "F64",
		},
		{
			description: "float32",
			expect:      100 + structArchetype.F32,
			name:        "F32",
		},
		{
			description: "bytes",
			expect:      structArchetype.Bs,
			name:        "Bs",
		},
		{
			description: "time",
			expect:      structArchetype.T,
			name:        "T",
		},
	}

	aStruct1Type := reflect.TypeOf(Struct1{})
	aStructAddr1 := xunsafe.Addr(structInstance1)
	aStructAddr2 := xunsafe.Addr(structInstance2)

	for _, testCase := range testCases {
		checkInterfaceMutator(t, aStruct1Type, testCase, aStructAddr1, structInstance1)
		checkTypedMutator(t, aStruct1Type, testCase, aStructAddr2, structInstance2)
	}
}

func checkInterfaceMutator(t *testing.T, aStruct1Type reflect.Type, testCase struct {
	description string
	expect      interface{}
	actual      func() interface{}
	name        string
}, aStructAddr unsafe.Pointer, aStruct *Struct1) {
	field := xunsafe.FieldByName(aStruct1Type, testCase.name)
	structValue := reflect.ValueOf(aStruct)

	field.Set(aStructAddr, testCase.expect)
	checkIfSetProperly(t, structValue, aStructAddr, field, testCase)
}

func checkTypedMutator(t *testing.T, aStruct1Type reflect.Type, testCase struct {
	description string
	expect      interface{}
	actual      func() interface{}
	name        string
}, aStructAddr unsafe.Pointer, aStruct1 *Struct1) {
	field := xunsafe.FieldByName(aStruct1Type, testCase.name)
	switch val := testCase.expect.(type) {
	case int:
		field.SetInt(aStructAddr, val)
	case int64:
		field.SetInt64(aStructAddr, val)
	case int32:
		field.SetInt32(aStructAddr, val)
	case int16:
		field.SetInt16(aStructAddr, val)
	case int8:
		field.SetInt8(aStructAddr, val)
	case uint:
		field.SetUint(aStructAddr, val)
	case uint64:
		field.SetUint64(aStructAddr, val)
	case uint32:
		field.SetUint32(aStructAddr, val)
	case uint16:
		field.SetUint16(aStructAddr, val)
	case uint8:
		field.SetUint8(aStructAddr, val)
	case bool:
		field.SetBool(aStructAddr, val)
	case string:
		field.SetString(aStructAddr, val)
	case []byte:
		field.SetBytes(aStructAddr, val)
	case float64:
		field.SetFloat64(aStructAddr, val)
	case float32:
		field.SetFloat32(aStructAddr, val)
	case time.Time:
		field.SetTime(aStructAddr, val)
	case []Bar:
		field.SetValue(aStructAddr, val)
	case *int:
		field.SetIntPtr(aStructAddr, val)
	case *int64:
		field.SetInt64Ptr(aStructAddr, val)
	case *int32:
		field.SetInt32Ptr(aStructAddr, val)
	case *int16:
		field.SetInt16Ptr(aStructAddr, val)
	case *int8:
		field.SetInt8Ptr(aStructAddr, val)
	case *uint:
		field.SetUintPtr(aStructAddr, val)
	case *uint64:
		field.SetUint64Ptr(aStructAddr, val)
	case *uint32:
		field.SetUint32Ptr(aStructAddr, val)
	case *uint16:
		field.SetUint16Ptr(aStructAddr, val)
	case *uint8:
		field.SetUint8Ptr(aStructAddr, val)
	case *bool:
		field.SetBoolPtr(aStructAddr, val)
	case *string:
		field.SetStringPtr(aStructAddr, val)
	case *[]byte:
		field.SetBytesPtr(aStructAddr, val)
	case *float64:
		field.SetFloat64Ptr(aStructAddr, val)
	case *float32:
		field.SetFloat32Ptr(aStructAddr, val)
	case *time.Time:
		field.SetTimePtr(aStructAddr, val)
	case *Foo:
		field.SetValue(aStructAddr, val)
	case Foo:
		field.Set(aStructAddr, val)
	}

	structValue := reflect.ValueOf(aStruct1)
	checkIfSetProperly(t, structValue, aStructAddr, field, testCase)
}

func checkIfSetProperly(t *testing.T, structValue reflect.Value, structAddr unsafe.Pointer, field *xunsafe.Field, testCase struct {
	description string
	expect      interface{}
	actual      func() interface{}
	name        string
}) {
	actual := structValue.Elem().FieldByName(testCase.name).Interface()
	valuePtr, getValue := fieldValueHolder(field)
	field.GetInto(structAddr, valuePtr)
	if _, ok := actual.(func()); ok {
		assert.EqualValues(t, fmt.Sprintf("%v", testCase.expect), fmt.Sprintf("%v", actual), testCase.description)
	} else {
		assert.EqualValues(t, testCase.expect, actual, testCase.description)
		assert.EqualValues(t, testCase.expect, getValue(), testCase.description)
	}
}

func fieldValueHolder(field *xunsafe.Field) (interface{}, func() interface{}) {
	switch field.Type.Kind() {
	case reflect.Int:
		intPtr := new(int)
		return intPtr, func() interface{} {
			return *intPtr
		}
	case reflect.Int64:
		int64Ptr := new(int64)
		return int64Ptr, func() interface{} {
			return *int64Ptr
		}
	case reflect.Int32:
		int32Ptr := new(int32)
		return int32Ptr, func() interface{} {
			return *int32Ptr
		}
	case reflect.Int16:
		int16Ptr := new(int16)
		return int16Ptr, func() interface{} {
			return *int16Ptr
		}
	case reflect.Int8:
		int8Ptr := new(int8)
		return int8Ptr, func() interface{} {
			return *int8Ptr
		}
	case reflect.Uint:
		uintPtr := new(uint)
		return uintPtr, func() interface{} {
			return *uintPtr
		}
	case reflect.Uint64:
		uint64Ptr := new(uint64)
		return uint64Ptr, func() interface{} {
			return *uint64Ptr
		}
	case reflect.Uint32:
		uint32Ptr := new(uint32)
		return uint32Ptr, func() interface{} {
			return *uint32Ptr
		}
	case reflect.Uint16:
		uint16Ptr := new(uint16)
		return uint16Ptr, func() interface{} {
			return *uint16Ptr
		}
	case reflect.Uint8:
		uint8Ptr := new(uint8)
		return uint8Ptr, func() interface{} {
			return *uint8Ptr
		}
	case reflect.Bool:
		boolPtr := new(bool)
		return boolPtr, func() interface{} {
			return *boolPtr
		}
	case reflect.String:
		stringPtr := new(string)
		return stringPtr, func() interface{} {
			return *stringPtr
		}
	case reflect.Float64:
		float64Ptr := new(float64)
		return float64Ptr, func() interface{} {
			return *float64Ptr
		}
	case reflect.Float32:
		float32Ptr := new(float32)
		return float32Ptr, func() interface{} {
			return *float32Ptr
		}
	case reflect.Ptr:
		interfacePtrPtr := new(interface{})
		return interfacePtrPtr, func() interface{} {
			return *interfacePtrPtr
		}
	default:
		if field.Type.ConvertibleTo(reflect.TypeOf(time.Time{})) {
			timePtr := new(time.Time)
			return timePtr, func() interface{} {
				return *timePtr
			}
		}

		if field.Type.ConvertibleTo(reflect.TypeOf(&time.Time{})) {
			timePtr := new(*time.Time)
			return timePtr, func() interface{} {
				return *timePtr
			}
		}

		interfacePointer := new(interface{})
		return interfacePointer, func() interface{} {
			return *interfacePointer
		}
	}
}
