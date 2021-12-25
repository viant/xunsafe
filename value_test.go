package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func TestValuePointerForType(t *testing.T) {
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
		fn := ValuePointerForType(testCase.t)
		if !assert.NotNil(t, fn, testCase.description) {
			continue
		}
		ptr := reflect.New(testCase.t)
		ptr.Elem().Set(reflect.ValueOf(testCase.value))
		v := AsPointer(ptr.Interface())
		actual := testCase.castBack(v)
		assert.EqualValues(t, testCase.value, actual, testCase.description)
	}
}
