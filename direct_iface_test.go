package xunsafe

import (
	"reflect"
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

func TestType_DirectIfaceStruct(t *testing.T) {
	type singlePtr struct{ P *int }
	v := 0
	aType := NewType(reflect.TypeOf(singlePtr{}))
	for _, testCase := range []struct {
		name string
		p    *int
	}{
		{name: "non-nil", p: &v},
		{name: "nil"},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			value := singlePtr{P: testCase.p}
			check := func(t *testing.T, actual singlePtr) {
				t.Helper()
				if actual.P != value.P {
					t.Errorf("P = %p, want %p", actual.P, value.P)
				}
			}
			t.Run("DerefValue", func(t *testing.T) {
				check(t, aType.Deref(value).(singlePtr))
			})
			t.Run("DerefPointer", func(t *testing.T) {
				check(t, aType.Deref(&value).(singlePtr))
			})
			t.Run("Interface", func(t *testing.T) {
				check(t, aType.Interface(unsafe.Pointer(&value)).(singlePtr))
			})
			t.Run("Value", func(t *testing.T) {
				check(t, aType.Value(unsafe.Pointer(&value)).(singlePtr))
			})
			t.Run("Ref", func(t *testing.T) {
				ref := aType.Ref(value).(*singlePtr)
				if ref == nil {
					t.Fatal("Ref returned nil for a struct value")
				}
				check(t, *ref)
				check(t, aType.Deref(ref).(singlePtr))
			})
		})
	}
}

func TestField_InterfaceDirectIfaceStruct(t *testing.T) {
	type singlePtr struct{ P *int }
	type holder struct {
		Prefix int
		Value  singlePtr
	}
	v := 0
	field := FieldByName(reflect.TypeOf(holder{}), "Value")
	for _, testCase := range []struct {
		name string
		p    *int
	}{
		{name: "non-nil", p: &v},
		{name: "nil"},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			value := holder{Prefix: 1, Value: singlePtr{P: testCase.p}}
			actual := field.Interface(unsafe.Pointer(&value)).(singlePtr)
			if actual.P != value.Value.P {
				t.Errorf("P = %p, want %p", actual.P, value.Value.P)
			}
		})
	}
}

func TestSlice_ValueAtDirectIfaceStruct(t *testing.T) {
	type singlePtr struct{ P *int }
	v, other := 42, 7
	items := []singlePtr{{P: &v}, {}, {P: &other}}
	slice := NewSlice(reflect.TypeOf(items))
	field := FieldByName(reflect.TypeOf(singlePtr{}), "P")
	for i, expected := range items {
		actual := slice.ValueAt(unsafe.Pointer(&items), i).(singlePtr)
		if actual.P != expected.P {
			t.Errorf("ValueAt(%d).P = %p, want %p", i, actual.P, expected.P)
		}
		fieldValue := field.Interface(AsPointer(actual)).(*int)
		if fieldValue != expected.P {
			t.Errorf("ValueAt(%d) via AsPointer and Field.Interface = %p, want %p", i, fieldValue, expected.P)
		}
	}
}

func TestSlice_ValueAtPointerIdentity(t *testing.T) {
	v, other := 42, 7
	items := []*int{&v, nil, &other}
	slice := NewSlice(reflect.TypeOf(items))
	for i, expected := range items {
		actual := slice.ValueAt(unsafe.Pointer(&items), i).(*int)
		if actual != expected {
			t.Errorf("ValueAt(%d) = %p, want %p", i, actual, expected)
		}
	}
}
