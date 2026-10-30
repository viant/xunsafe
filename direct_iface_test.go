package xunsafe

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestAsPointer_DirectIface(t *testing.T) {
	type singlePtr struct{ P *int }
	type nestedSinglePtr struct{ Inner singlePtr }
	type node struct{ Next *node }
	type valueOnly struct{ ID int }
	type ptrAndValue struct {
		P *int
		X int
	}

	v := 42
	other := 7
	next := &node{}
	singlePtrValue := &singlePtr{P: &v}

	var testCases = []struct {
		description string
		value       interface{}
		validate    func(ptr unsafe.Pointer) bool
	}{
		{
			description: "single pointer struct",
			value:       singlePtr{P: &v},
			validate: func(ptr unsafe.Pointer) bool {
				actual := (*singlePtr)(ptr)
				return actual.P == &v && *actual.P == 42
			},
		},
		{
			description: "single pointer struct with nil field",
			value:       singlePtr{},
			validate: func(ptr unsafe.Pointer) bool {
				return ptr != nil && (*singlePtr)(ptr).P == nil
			},
		},
		{
			description: "nested single pointer struct",
			value:       nestedSinglePtr{Inner: singlePtr{P: &v}},
			validate: func(ptr unsafe.Pointer) bool {
				actual := (*nestedSinglePtr)(ptr)
				return actual.Inner.P == &v
			},
		},
		{
			description: "self referencing single pointer struct",
			value:       node{Next: next},
			validate: func(ptr unsafe.Pointer) bool {
				return (*node)(ptr).Next == next
			},
		},
		{
			description: "pointer to single pointer struct",
			value:       singlePtrValue,
			validate: func(ptr unsafe.Pointer) bool {
				return ptr == unsafe.Pointer(singlePtrValue)
			},
		},
		{
			description: "non pointer single field struct",
			value:       valueOnly{ID: 3},
			validate: func(ptr unsafe.Pointer) bool {
				return (*valueOnly)(ptr).ID == 3
			},
		},
		{
			description: "pointer and value struct",
			value:       ptrAndValue{P: &other, X: 5},
			validate: func(ptr unsafe.Pointer) bool {
				actual := (*ptrAndValue)(ptr)
				return actual.P == &other && actual.X == 5
			},
		},
		{
			description: "*int",
			value:       &v,
			validate: func(ptr unsafe.Pointer) bool {
				return ptr == unsafe.Pointer(&v)
			},
		},
		{
			description: "nil interface",
			value:       nil,
			validate: func(ptr unsafe.Pointer) bool {
				return ptr == nil
			},
		},
	}

	for _, testCase := range testCases {
		actual := AsPointer(testCase.value)
		assert.True(t, testCase.validate(actual), testCase.description)
	}
}
