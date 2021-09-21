package eval

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/xunsafe"
	"reflect"
	"testing"
)

func TestSelector_IntAddr(t *testing.T) {

	type Foo struct {
		ID   int
		Name string
	}

	var testCases = []struct {
		description string
		source      interface{}
		selector    string
		expect      interface{}
		expectNil   bool
	}{

		{
			description: "top level string",
			source: struct {
				Name string
			}{Name: "test"},
			selector: "Name",
			expect:   "test",
		},
		{
			description: "child node int",
			source: struct {
				ID    int
				Child struct {
					N int
				}
			}{
				ID: 12,
				Child: struct {
					N int
				}{N: 23},
			},
			selector: "Child.N",
			expect:   23,
		},
		{
			description: "array item bool",
			source: struct {
				Items []struct {
					Active bool
					N      int
				}
			}{
				Items: []struct {
					Active bool
					N      int
				}{
					{N: 10},
					{Active: true, N: 20},
				},
			},
			selector: "Items[1].Active",
			expect:   true,
		},

		{
			description: "array item string",
			source: struct {
				Items []struct {
					ID   int
					Name string
				}
			}{
				Items: []struct {
					ID   int
					Name string
				}{
					{ID: 1, Name: "Adam"},
					{ID: 2, Name: "Ben"},
					{ID: 3, Name: "Chris"},
				},
			},
			selector: "Items[1].Name",
			expect:   "Ben",
		},
		{
			description: "array  item ptr string",
			source: struct {
				Items []*Foo
			}{
				Items: []*Foo{
					{ID: 1, Name: "Adam"},
					{ID: 2, Name: "Ben"},
					{ID: 3, Name: "Chris"},
				},
			},
			selector: "Items[1].Name",
			expect:   "Ben",
		},
		{
			description: "child ptr node string",
			source: struct {
				ID    int
				Child *Foo
			}{
				ID: 12,
				Child: &Foo{
					Name: "Ted",
				},
			},
			selector: "Child.Name",
			expect:   "Ted",
		},
		{
			description: "child ptr nil node int",
			source: struct {
				ID    int
				Child *Foo
			}{
				ID: 12,
			},
			selector:  "Child.ID",
			expect:    0,
			expectNil: true,
		},
	}

	for _, testCase := range testCases {
		val := reflect.ValueOf(testCase.source)
		ptr := reflect.New(val.Type())
		ptr.Elem().Set(val)
		instance := ptr.Interface()
		selector, err := NewSelector(val.Type(), testCase.selector)
		assert.Nil(t, err, testCase.description)
		switch expect := testCase.expect.(type) {
		case int:
			valPtr := selector.IntAddr(xunsafe.Addr(instance))
			if testCase.expectNil {
				assert.Nil(t, valPtr, testCase.description)
				continue
			}
			assert.EqualValues(t, expect, *valPtr)
		case string:
			valPtr := selector.StringAddr(xunsafe.Addr(instance))
			assert.EqualValues(t, expect, *valPtr)
		case bool:
			valPtr := selector.BoolAddr(xunsafe.Addr(instance))
			if expect {
				assert.True(t, *valPtr, testCase.description)
			} else {
				assert.False(t, *valPtr, testCase.description)
			}

		default:

			assert.Fail(t, testCase.description)
		}

	}

}
