package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestFieldPointer(t *testing.T) {

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
