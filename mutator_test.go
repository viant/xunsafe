package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type MutBenchStruct struct {
	ID   int
	Name string
	Val  float32
}

var _mutBenchInstance = &MutBenchStruct{
	ID:   102,
	Name: "test",
	Val:  232.2,
}

var _mutIDField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "ID")
var _mutNameField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "Name")
var _mutValField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "Val")

func BenchmarkField_Mutator_Native(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_mutBenchInstance.ID = id
		_mutBenchInstance.Name = name
		_mutBenchInstance.Val = val
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func Benchmark_Mutator_Fast(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	addr := Addr(_mutBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_mutIDField.SetInt(addr, id)
		_mutNameField.SetString(addr, name)
		_mutValField.SetFloat32(addr, val)
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func Benchmark_GenericMutator_Fast(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	addr := Addr(_mutBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_mutIDField.Set(addr, id)
		_mutNameField.Set(addr, name)
		_mutValField.Set(addr, val)
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func Benchmark_Mutator_Fast_Ptr(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	holderPtr := Addr(_mutBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		*(_mutIDField.Addr(holderPtr).(*int)) = id
		*(_mutNameField.Addr(holderPtr).(*string)) = name
		*(_mutValField.Addr(holderPtr).(*float32)) = val
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func BenchmarkField_Mutator_Reflect(b *testing.B) {
	aType := reflect.TypeOf(AccBenchStruct{})
	var idFieldIdx, nameFiledIdx, valFieldIdx int

	if field, ok := aType.FieldByName("ID"); ok {
		idFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Name"); ok {
		nameFiledIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Val"); ok {
		valFieldIdx = field.Index[0]
	}

	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	instanceVal := reflect.ValueOf(_mutBenchInstance).Elem()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		instanceVal.Field(idFieldIdx).Set(reflect.ValueOf(id))
		instanceVal.Field(nameFiledIdx).Set(reflect.ValueOf(name))
		instanceVal.Field(valFieldIdx).Set(reflect.ValueOf(val))
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func BenchmarkField_Mutator_Reflect_Ptr(b *testing.B) {
	aType := reflect.TypeOf(AccBenchStruct{})
	var idFieldIdx, nameFiledIdx, valFieldIdx int

	if field, ok := aType.FieldByName("ID"); ok {
		idFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Name"); ok {
		nameFiledIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Val"); ok {
		valFieldIdx = field.Index[0]
	}

	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	instanceVal := reflect.ValueOf(_mutBenchInstance).Elem()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		*(instanceVal.Field(idFieldIdx).Addr().Interface().(*int)) = id
		*(instanceVal.Field(nameFiledIdx).Addr().Interface().(*string)) = name
		*(instanceVal.Field(valFieldIdx).Addr().Interface().(*float32)) = val
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}
