package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestDerefSafePointer(t *testing.T) {

	type Bar struct {
		Id   int
		Name string
		Flag *int
	}
	type Foo struct {
		ID  int
		Bar *Bar
	}
	foo := &Foo{}
	foo2 := &Foo{Bar: &Bar{}}

	var testCases = []struct {
		description string
		init        func() (unsafe.Pointer, reflect.Type)
		validate    func(ptr unsafe.Pointer) bool
	}{
		{
			description: "**Bar",
			init: func() (unsafe.Pointer, reflect.Type) {
				ptr := unsafe.Pointer(&foo.Bar)
				return ptr, reflect.TypeOf(foo.Bar)
			},
			validate: func(ptr unsafe.Pointer) bool {
				barRef := (*Bar)(ptr)
				barRef.Name = "test"
				return assert.EqualValues(t, foo.Bar.Name, barRef.Name)
			},
		},
		{
			description: "**int",
			init: func() (unsafe.Pointer, reflect.Type) {
				ptr := unsafe.Pointer(&foo2.Bar.Flag)
				return ptr, reflect.TypeOf(foo2.Bar.Flag)
			},
			validate: func(ptr unsafe.Pointer) bool {
				barRef := (*int)(ptr)
				*barRef = 123
				return assert.EqualValues(t, *foo2.Bar.Flag, *barRef)
			},
		},
	}

	for _, testCase := range testCases {
		p, pType := testCase.init()
		actual := SafeDerefPointer(p, pType)
		assert.True(t, testCase.validate(actual), testCase.description)
	}

}

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
