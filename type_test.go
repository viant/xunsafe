package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"io"
	"reflect"
	"testing"
)

func TestType_Type(t *testing.T) {
	var e error
	e = io.EOF
	var z = 0
	type Foo struct {
		Name string
		ID   int
	}

	var foo = &Foo{Name: "abc"}
	var testCases = []struct {
		description string
		value       interface{}
	}{
		{
			description: "int",
			value:       100,
		},
		{
			description: "**Foo",
			value:       &foo,
		},
		{
			description: "string",
			value:       "abc",
		},
		{
			description: "struct",
			value: struct {
				ID   int
				Name string
			}{1, "abc"},
		},
		{
			description: "slice",
			value: []struct {
				ID   int
				Name string
			}{{1, "abc"}},
		},
		{
			description: "*struct",
			value: &Foo{
				ID:   12,
				Name: "123123",
			},
		},
		{
			description: "*int",
			value:       &z,
		},
		{
			description: "error",
			value:       e,
		},
	}

	for _, testCase := range testCases {
		if testCase.description != "*struct" {
			continue
		}
		aType := NewType(reflect.TypeOf(testCase.value))

		ptr := aType.Ref(testCase.value)
		value := reflect.ValueOf(ptr)

		assert.EqualValues(t, reflect.Ptr, value.Kind(), testCase.description)
		assert.EqualValues(t, testCase.value, value.Elem().Interface(), testCase.description)
		deref := aType.Deref(ptr)
		assert.EqualValues(t, testCase.value, deref, testCase.description)
		iFace := aType.Interface(AsPointer(deref))
		assert.EqualValues(t, iFace, deref, testCase.description)
	}

}
