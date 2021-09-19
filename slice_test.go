package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func TestSlice_Range(t *testing.T) {
	type Foo struct {
		ID int
		Name string
	}
	ID := FieldByName(reflect.TypeOf(Foo{}), "ID")

	var testCases = []struct{
		description string
		source interface{}
		expect interface{}
	}{
		{
			description: "slice",
			source: []Foo{
				{
					ID: 1,
					Name: "abc",
				},
				{
					ID: 12,
					Name: "xyz",
				},
			},
			expect: []interface{}{1, 12},
		},
		{
			description: "slice item pointer",
			source: []*Foo{
				{
					ID: 1,
					Name: "abc",
				},
				{
					ID: 12,
					Name: "xyz",
				},
			},
			expect: []interface{}{1, 12},
		},
	}


	for _, testCase := range testCases {
		aSlice := NewSlice(reflect.TypeOf(testCase.source))
		holderAddr := Addr(testCase.source)
		actual := make([]interface{}, 0)
		aSlice.Range(holderAddr, func(index int, addr unsafe.Pointer) bool {
			actual = append(actual, ID.Int(addr))
			return true
		})
		assert.EqualValues(t, testCase.expect, actual)
	}

}
