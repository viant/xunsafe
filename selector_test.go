package xunsafe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSelector_IntAddr(t *testing.T) {

	type Foo struct {
		ID   int
		Name string
		Fn   func()
		Foos []Foo
	}

	var testCases = []struct {
		description string
		source      interface{}
		selector    string
		indexes     []int
		expect      interface{}
		expectNil   bool
	}{

		{
			description: "nexted slice",
			source: struct {
				Z []Foo
			}{Z: []Foo{
				{
					ID: 1,
					Foos: []Foo{
						{
							Name: "Tester 1",
						},
						{
							Name: "Tester 2",
						},
					},
				},
			}},
			selector: "Z[].Foos[].Name",
			indexes:  []int{0, 1},
			expect:   "Tester 2",
		},

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
			description: "child fn node string",
			source: struct {
				ID    int
				Child *Foo
			}{
				ID: 12,
				Child: &Foo{
					Name: "Ted",
					Fn: func() {
						fmt.Println("123")
					},
				},
			},
			selector: "Child.Fn",
			expect: func() {
				fmt.Println("123")
			},
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
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		switch expect := testCase.expect.(type) {
		case int:
			valPtr := selector.IntAddr(Addr(instance))
			if testCase.expectNil {
				assert.Nil(t, valPtr, testCase.description)
				continue
			}
			assert.EqualValues(t, expect, *valPtr)
		case string:
			if len(testCase.indexes) > 0 {
				actual := selector.IValue(Addr(instance), testCase.indexes)
				assert.EqualValues(t, expect, actual)
			} else {
				valPtr := selector.StringAddr(Addr(instance))
				assert.EqualValues(t, expect, *valPtr)
			}
		case bool:
			valPtr := selector.BoolAddr(Addr(instance))
			if expect {
				assert.True(t, *valPtr, testCase.description)
			} else {
				assert.False(t, *valPtr, testCase.description)
			}
		case func():
			val := selector.Value(Addr(instance))
			selector.Set(Addr(instance), val)

		default:

			assert.Fail(t, testCase.description)
		}

	}

}
