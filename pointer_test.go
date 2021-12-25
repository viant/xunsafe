package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TesAsPointer(t *testing.T) {
	k := 10
	l := 20
	i := int64(30)
	type Foo struct {
		ID   int
		Name string
	}
	var testCases = []struct {
		description string
		t           reflect.Type
		value       interface{}
		castBack    func(pointer unsafe.Pointer) interface{}
	}{

		{
			description: "int",
			t:           reflect.TypeOf(0),
			value:       10,
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*int)(pointer)
			},
		},
		{
			description: "*int64",
			t:           reflect.TypeOf(&i),
			value:       &i,
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(**int64)(pointer)
			},
		},
		{
			description: "[]*int",
			t:           reflect.TypeOf([]*int{}),
			value:       []*int{&k, &l},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*[]*int)(pointer)
			},
		},
		{
			description: "[]byte",
			t:           reflect.TypeOf([]byte{}),
			value:       []byte{1, 2, 3},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*[]byte)(pointer)
			},
		},
		{
			description: "struct ",
			t:           reflect.TypeOf(Foo{}),
			value:       Foo{ID: 1, Name: "abc"},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*Foo)(pointer)
			},
		},
		{
			description: "*struct ",
			t:           reflect.TypeOf(&Foo{}),
			value:       &Foo{ID: 1, Name: "abc"},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(**Foo)(pointer)
			},
		},
		{
			description: "[]struct ",
			t:           reflect.TypeOf([]Foo{}),
			value:       []Foo{{ID: 1, Name: "abc"}},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*[]Foo)(pointer)
			},
		},
		{
			description: "[]*struct ",
			t:           reflect.TypeOf([]*Foo{}),
			value:       []*Foo{{ID: 1, Name: "abc"}},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*[]*Foo)(pointer)
			},
		},
		{
			description: "[]*struct ",
			t:           reflect.TypeOf([]interface{}{}),
			value:       []interface{}{&Foo{ID: 1, Name: "abc"}},
			castBack: func(pointer unsafe.Pointer) interface{} {
				return *(*[]interface{})(pointer)
			},
		},
	}

	for _, testCase := range testCases {
		ptr := reflect.New(testCase.t)
		ptr.Elem().Set(reflect.ValueOf(testCase.value))
		v := AsPointer(ptr.Interface())
		actual := testCase.castBack(v)
		assert.EqualValues(t, testCase.value, actual, testCase.description)
	}
}

func TestEnsurePointer(t *testing.T) {

	var testCases = []struct {
		description string
		instance    interface{}
		Index       int
		expect      interface{}
	}{
		{
			description: "int type",
			instance: struct {
				A []int
				B int
				T time.Time
			}{
				nil, 101, time.Unix(0, 0),
			},
			Index:  1,
			expect: 101,
		},
		{
			description: "time type",
			instance: struct {
				A []int
				B int
				T time.Time
			}{
				nil, 101, time.Unix(0, 102),
			},
			Index:  2,
			expect: time.Unix(0, 102),
		},
		{
			description: "[]byte type",
			instance: struct {
				A []int
				B int
				C []byte
			}{
				nil, 102, []byte{'a', 'c', 'b'},
			},
			Index:  2,
			expect: []byte{'a', 'c', 'b'},
		},
		{
			description: "*bool type",
			instance: struct {
				A []int
				B int
				C []byte
				D *bool
			}{
				nil, 102, []byte{'a', 'c', 'b'}, pBool(true),
			},
			Index:  3,
			expect: true,
		},
		{
			description: "*[]string type",
			instance: struct {
				A []int
				B int
				C []byte
				D *bool
				E *[]string
			}{
				nil, 102, []byte{'a', 'c', 'b'}, pBool(true), pStrings([]string{"a", "1", "a"}),
			},
			Index:  4,
			expect: []string{"a", "1", "a"},
		},
	}

	for _, testCase := range testCases {
		testType := reflect.TypeOf(testCase.instance)
		instanceAddr := EnsurePointer(testCase.instance)
		field := FieldByIndex(testType, testCase.Index)
		actualPtr := field.Addr(instanceAddr)
		actual := reflect.ValueOf(actualPtr)
		for actual.Kind() == reflect.Ptr {
			actual = actual.Elem()
		}
		assert.EqualValues(t, testCase.expect, actual.Interface(), testCase.description)
	}

}

func pBool(b bool) *bool {
	return &b
}

func pStrings(s []string) *[]string {
	return &s
}
