package xunsafe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestIf(t *testing.T) {
	type Bar struct {
		ID   int
		Name string
	}
	type Foo struct {
		Bar *Bar
	}
	f := &Foo{}
	x := NewStruct(reflect.TypeOf(f))
	bt := reflect.TypeOf(&Bar{})
	v := reflect.New(bt.Elem()).Interface()
	ptr := unsafe.Pointer(f)
	x.Fields[0].SetValue(ptr, v)
	f.Bar.ID = 123

}

func TestField_Accessor(t *testing.T) {

	type Bar struct {
		ID int
	}

	type Foo struct {
		ID int
	}

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

		B         bool
		S         string
		Bs        []byte
		T         time.Time
		Bars      []Bar
		Foo       *Foo
		F2        *Foo
		Fn        func()
		Interface interface{}
	}

	type Struct2 struct {
		I   *int
		I64 *int64
		I32 *int32
		I16 *int16
		I8  *int8

		UI   *uint
		UI64 *uint64
		UI32 *uint32
		UI16 *uint16
		UI8  *uint8

		F64 *float64
		F32 *float32

		B  *bool
		S  *string
		Bs *[]byte
		T  *time.Time
	}

	aStruct1 := &Struct1{
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
		Foo: &Foo{ID: 1},
		Fn: func() {

		},
		Interface: &Foo{ID: 1},
	}

	aStruct2 := &Struct2{
		I:   &aStruct1.I,
		I64: &aStruct1.I64,
		I32: &aStruct1.I32,
		I16: &aStruct1.I16,
		I8:  &aStruct1.I8,

		UI:   &aStruct1.UI,
		UI64: &aStruct1.UI64,
		UI32: &aStruct1.UI32,
		UI16: &aStruct1.UI16,
		UI8:  &aStruct1.UI8,

		F64: &aStruct1.F64,
		F32: &aStruct1.F32,

		B:  &aStruct1.B,
		S:  &aStruct1.S,
		Bs: &aStruct1.Bs,
		T:  &aStruct1.T,
	}

	var testCases = []struct {
		description string
		expect      interface{}
		name        string
	}{
		{
			description: "interface",
			expect:      aStruct1.Interface,
			name:        "Interface",
		},
		{
			description: "function",
			expect:      aStruct1.Fn,
			name:        "Fn",
		},
		{
			description: "int",
			expect:      aStruct1.I,
			name:        "I",
		},

		{
			description: "Foo",
			expect:      aStruct1.Foo,
			name:        "Foo",
		},
		{
			description: "F2",
			expect:      aStruct1.F2,
			name:        "F2",
		},
		{
			description: "int64",
			expect:      aStruct1.I64,
			name:        "I64",
		},
		{
			description: "int32",
			expect:      aStruct1.I32,
			name:        "I32",
		},
		{
			description: "int16",
			expect:      aStruct1.I16,
			name:        "I16",
		},
		{
			description: "int8",
			expect:      aStruct1.I8,
			name:        "I8",
		},
		{
			description: "uint",
			expect:      aStruct1.UI,
			name:        "UI",
		},
		{
			description: "uint64",
			expect:      aStruct1.UI64,
			name:        "UI64",
		},
		{
			description: "uint32",
			expect:      aStruct1.UI32,
			name:        "UI32",
		},
		{
			description: "uint16",
			expect:      aStruct1.UI16,
			name:        "UI16",
		},
		{
			description: "uint8",
			expect:      aStruct1.UI8,
			name:        "UI8",
		},
		{
			description: "string",
			expect:      aStruct1.S,
			name:        "S",
		},
		{
			description: "bool",
			expect:      aStruct1.B,
			name:        "B",
		},
		{
			description: "float64",
			expect:      aStruct1.F64,
			name:        "F64",
		},
		{
			description: "float32",
			expect:      aStruct1.F32,
			name:        "F32",
		},
		{
			description: "bytes",
			expect:      aStruct1.Bs,
			name:        "Bs",
		},
		{
			description: "time",
			expect:      aStruct1.T,
			name:        "T",
		},

		{
			description: "[]Bar",
			expect:      aStruct1.Bars,
			name:        "Bars",
		},

		{
			description: "*int",
			expect:      aStruct2.I,
			name:        "I",
		},
		{
			description: "*int64",
			expect:      aStruct2.I64,
			name:        "I64",
		},
		{
			description: "*int32",
			expect:      aStruct2.I32,
			name:        "I32",
		},
		{
			description: "*int16",
			expect:      aStruct2.I16,
			name:        "I16",
		},
		{
			description: "*int8",
			expect:      aStruct2.I8,
			name:        "I8",
		},
		{
			description: "*uint",
			expect:      aStruct2.UI,
			name:        "UI",
		},
		{
			description: "*uint64",
			expect:      aStruct2.UI64,
			name:        "UI64",
		},
		{
			description: "*uint32",
			expect:      aStruct2.UI32,
			name:        "UI32",
		},
		{
			description: "*uint16",
			expect:      aStruct2.UI16,
			name:        "UI16",
		},
		{
			description: "*uint8",
			expect:      aStruct2.UI8,
			name:        "UI8",
		},
		{
			description: "*string",
			expect:      aStruct2.S,
			name:        "S",
		},
		{
			description: "*bool",
			expect:      aStruct2.B,
			name:        "B",
		},
		{
			description: "*float64",
			expect:      aStruct2.F64,
			name:        "F64",
		},
		{
			description: "*float32",
			expect:      aStruct2.F32,
			name:        "F32",
		},
		{
			description: "*bytes",
			expect:      aStruct2.Bs,
			name:        "Bs",
		},
		{
			description: "*time",
			expect:      aStruct2.T,
			name:        "T",
		},
	}

	aStruct1Type := reflect.TypeOf(Struct1{})
	aStruct1Addr := EnsurePointer(aStruct1)

	aStruct2Type := reflect.TypeOf(Struct2{})
	aStruct2Addr := EnsurePointer(aStruct2)

	assert.Nil(t, aStruct1.F2)
	for _, testCase := range testCases {

		var field *Field
		var aStructAddr unsafe.Pointer

		if strings.Contains(testCase.description, "*") {
			field = FieldByName(aStruct2Type, testCase.name)
			aStructAddr = aStruct2Addr

		} else {
			field = FieldByName(aStruct1Type, testCase.name)
			aStructAddr = aStruct1Addr
		}

		var actual interface{}
		switch testCase.expect.(type) {
		case int:
			actual = field.Int(aStructAddr)
		case int64:
			actual = field.Int64(aStructAddr)
		case int32:
			actual = field.Int32(aStructAddr)
		case int16:
			actual = field.Int16(aStructAddr)
		case int8:
			actual = field.Int8(aStructAddr)
		case uint:
			actual = field.Uint(aStructAddr)
		case uint64:
			actual = field.Uint64(aStructAddr)
		case uint32:
			actual = field.Uint32(aStructAddr)
		case uint16:
			actual = field.Uint16(aStructAddr)
		case uint8:
			actual = field.Uint8(aStructAddr)
		case bool:
			actual = field.Bool(aStructAddr)
		case string:
			actual = field.String(aStructAddr)
		case []byte:
			actual = field.Bytes(aStructAddr)
		case float64:
			actual = field.Float64(aStructAddr)
		case float32:
			actual = field.Float32(aStructAddr)
		case time.Time:
			actual = field.Time(aStructAddr)
		case []Bar:
			actual = field.Interface(aStructAddr)
		case *int:
			actual = field.IntPtr(aStructAddr)
		case *int64:
			actual = field.Int64Ptr(aStructAddr)
		case *int32:
			actual = field.Int32Ptr(aStructAddr)
		case *int16:
			actual = field.Int16Ptr(aStructAddr)
		case *int8:
			actual = field.Int8Ptr(aStructAddr)
		case *uint:
			actual = field.UintPtr(aStructAddr)
		case *uint64:
			actual = field.Uint64Ptr(aStructAddr)
		case *uint32:
			actual = field.Uint32Ptr(aStructAddr)
		case *uint16:
			actual = field.Uint16Ptr(aStructAddr)
		case *uint8:
			actual = field.Uint8Ptr(aStructAddr)
		case *bool:
			actual = field.BoolPtr(aStructAddr)
		case *string:
			actual = field.StringPtr(aStructAddr)
		case *[]byte:
			actual = field.BytesPtr(aStructAddr)
		case *float64:
			actual = field.Float64Ptr(aStructAddr)
		case *float32:
			actual = field.Float32Ptr(aStructAddr)
		case *time.Time:
			actual = field.TimePtr(aStructAddr)
		case *Foo:
			actual = field.Interface(aStructAddr)
		case func():
			actual = field.Interface(aStructAddr)
		default:
			actual = field.Interface(aStructAddr)
		}
		if _, ok := actual.(func()); ok {
			assert.EqualValues(t, fmt.Sprintf("%v", testCase.expect), fmt.Sprintf("%v", actual), testCase.description)
			continue
		}
		if !assert.EqualValues(t, testCase.expect, actual, testCase.description) {
			break
		}
	}

}

type AccBenchStruct struct {
	ID   int
	Name string
	Val  float32
	Time time.Time
}

var _accBenchInstance = &AccBenchStruct{
	ID:   102,
	Name: "test",
	Val:  232.2,
	Time: time.Now(),
}

var _AcciDField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "ID")
var _AccNameField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "Name")
var _AccValField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "Val")
var _TimeValField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "Time")

func BenchmarkField_Accessor_Native(b *testing.B) {
	var id int
	var name string
	var val float32
	var ts time.Time
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id = _accBenchInstance.ID
		name = _accBenchInstance.Name
		val = _accBenchInstance.Val
		ts = _accBenchInstance.Time
	}
	assert.EqualValues(b, _accBenchInstance.ID, id)
	assert.EqualValues(b, _accBenchInstance.Name, name)
	assert.EqualValues(b, _accBenchInstance.Val, val)
	assert.EqualValues(b, _accBenchInstance.Time, ts)

}

func BenchmarkField_Accessor_Direct_Xunsafe(b *testing.B) {
	var id int
	var name string
	var val float32
	var ts time.Time
	ptr := EnsurePointer(_accBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id = _AcciDField.Int(ptr)
		name = _AccNameField.String(ptr)
		val = _AccValField.Float32(ptr)
		ts = _TimeValField.Time(ptr)
	}
	assert.EqualValues(b, _accBenchInstance.ID, id)
	assert.EqualValues(b, _accBenchInstance.Name, name)
	assert.EqualValues(b, _accBenchInstance.Val, val)
	assert.EqualValues(b, _accBenchInstance.Time, ts)

}

func BenchmarkField_Accessor_Interface_Xunsafe(b *testing.B) {
	var id int
	var name string
	var val float32
	var ts time.Time
	ptr := EnsurePointer(_accBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id = _AcciDField.Interface(ptr).(int)
		name = _AccNameField.Interface(ptr).(string)
		val = _AccValField.Interface(ptr).(float32)
		ts = _TimeValField.Interface(ptr).(time.Time)
	}
	assert.EqualValues(b, _accBenchInstance.ID, id)
	assert.EqualValues(b, _accBenchInstance.Name, name)
	assert.EqualValues(b, _accBenchInstance.Val, val)
	assert.EqualValues(b, _accBenchInstance.Time, ts)
}

func BenchmarkField_Accessor_Interface_Reflect(b *testing.B) {
	aType := reflect.TypeOf(AccBenchStruct{})
	var idFieldIdx, nameFiledIdx, valFieldIdx, tsFieldIdx int

	if field, ok := aType.FieldByName("ID"); ok {
		idFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Name"); ok {
		nameFiledIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Val"); ok {
		valFieldIdx = field.Index[0]
	}

	if field, ok := aType.FieldByName("Time"); ok {
		tsFieldIdx = field.Index[0]
	}
	var id int
	var name string
	var val float32
	var ts time.Time
	instanceVal := reflect.ValueOf(_accBenchInstance).Elem()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id = instanceVal.Field(idFieldIdx).Interface().(int)
		name = instanceVal.Field(nameFiledIdx).Interface().(string)
		val = instanceVal.Field(valFieldIdx).Interface().(float32)
		ts = instanceVal.Field(tsFieldIdx).Interface().(time.Time)

	}
	assert.EqualValues(b, _accBenchInstance.ID, id)
	assert.EqualValues(b, _accBenchInstance.Name, name)
	assert.EqualValues(b, _accBenchInstance.Val, val)
	assert.EqualValues(b, _accBenchInstance.Time, ts)

}

func BenchmarkField_Accessor_Addr_Xunsafe(b *testing.B) {
	var id *int
	var name *string
	var val *float32
	var ts *time.Time
	ptr := EnsurePointer(_accBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id = _AcciDField.Addr(ptr).(*int)
		name = _AccNameField.Addr(ptr).(*string)
		val = _AccValField.Addr(ptr).(*float32)
		ts = _TimeValField.Addr(ptr).(*time.Time)
	}
	assert.EqualValues(b, _accBenchInstance.ID, *id)
	assert.EqualValues(b, _accBenchInstance.Name, *name)
	assert.EqualValues(b, _accBenchInstance.Val, *val)
	assert.EqualValues(b, _accBenchInstance.Time, *ts)
}

func BenchmarkField_Accessor_Addr_Reflect(b *testing.B) {
	aType := reflect.TypeOf(AccBenchStruct{})
	var idFieldIdx, nameFiledIdx, valFieldIdx, tsFieldIdx int

	if field, ok := aType.FieldByName("ID"); ok {
		idFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Name"); ok {
		nameFiledIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Val"); ok {
		valFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Time"); ok {
		tsFieldIdx = field.Index[0]
	}

	var id *int
	var name *string
	var val *float32
	var ts *time.Time
	instanceVal := reflect.ValueOf(_accBenchInstance).Elem()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id = instanceVal.Field(idFieldIdx).Addr().Interface().(*int)
		name = instanceVal.Field(nameFiledIdx).Addr().Interface().(*string)
		val = instanceVal.Field(valFieldIdx).Addr().Interface().(*float32)
		ts = instanceVal.Field(tsFieldIdx).Addr().Interface().(*time.Time)

	}
	assert.EqualValues(b, _accBenchInstance.ID, *id)
	assert.EqualValues(b, _accBenchInstance.Name, *name)
	assert.EqualValues(b, _accBenchInstance.Val, *val)
	assert.EqualValues(b, _accBenchInstance.Time, *ts)

}
