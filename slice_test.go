package xunsafe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func TestSlice_Range(t *testing.T) {
	type Foo struct {
		ID   int
		Name string
	}

	ID := FieldByName(reflect.TypeOf(Foo{}), "ID")
	Name := FieldByName(reflect.TypeOf(Foo{}), "Name")

	var testCases = []struct {
		description string
		field       *Field
		source      interface{}
		expect      interface{}
	}{
		{
			description: "slice",
			field:       ID,
			source: []Foo{
				{
					ID:   1,
					Name: "abc",
				},
				{
					ID:   12,
					Name: "xyz",
				},
			},
			expect: []interface{}{1, 12},
		},
		{
			description: "slice item pointer",
			field:       Name,
			source: []*Foo{
				{
					ID:   1,
					Name: "abc",
				},
				{
					ID:   12,
					Name: "xyz",
				},
			},
			expect: []interface{}{"abc", "xyz"},
		},
		{
			description: "primitive slice",
			field:       nil,
			source: []string{
				"abc",
				"xyz",
				"zzz",
			},
			expect: []interface{}{"abc",
				"xyz",
				"zzz"},
		},
	}

	for _, testCase := range testCases {
		aSlice := NewSlice(reflect.TypeOf(testCase.source))
		holderAddr := EnsurePointer(testCase.source)
		actual := make([]interface{}, 0)
		aSlice.Range(holderAddr, func(index int, item interface{}) bool {
			if testCase.field == nil {
				actual = append(actual, item)
				return true
			}
			val := testCase.field.Interface(AsPointer(item))
			actual = append(actual, val)
			return true
		})

		assert.EqualValues(t, testCase.expect, actual)
	}

}

func TestSlice_ValueAt(t *testing.T) {
	type Foo struct {
		ID   int
		Name string
	}

	ID := FieldByName(reflect.TypeOf(Foo{}), "ID")
	Name := FieldByName(reflect.TypeOf(Foo{}), "Name")

	var testCases = []struct {
		description string
		field       *Field
		source      interface{}
		expect      interface{}
	}{
		{
			description: "slice",
			field:       ID,
			source: []Foo{
				{
					ID:   1,
					Name: "abc",
				},
				{
					ID:   12,
					Name: "xyz",
				},
			},
			expect: []interface{}{1, 12},
		},
		{
			description: "slice item pointer",
			field:       Name,
			source: []*Foo{
				{
					ID:   1,
					Name: "abc",
				},
				{
					ID:   12,
					Name: "xyz",
				},
			},
			expect: []interface{}{"abc", "xyz"},
		},
		{
			description: "primitive slice",
			field:       nil,
			source: []string{
				"abc",
				"xyz",
				"zzz",
			},
			expect: []interface{}{"abc",
				"xyz",
				"zzz"},
		},
	}

	for _, testCase := range testCases {
		aSlice := NewSlice(reflect.TypeOf(testCase.source))
		holderAddr := EnsurePointer(testCase.source)
		actual := make([]interface{}, 0)

		for i := 0; i < aSlice.Len(holderAddr); i++ {
			item := aSlice.ValueAt(holderAddr, i)
			if testCase.field == nil {
				actual = append(actual, item)
				continue
			}
			val := testCase.field.Interface(AsPointer(item))
			actual = append(actual, val)
		}
		assert.EqualValues(t, testCase.expect, actual)

	}

}

func TestSlice_ValuePointerAt(t *testing.T) {
	type Foo struct {
		ID   int
		Name string
	}

	ID := FieldByName(reflect.TypeOf(Foo{}), "ID")
	Name := FieldByName(reflect.TypeOf(Foo{}), "Name")

	var testCases = []struct {
		description string
		field       *Field
		source      interface{}
		expect      interface{}
	}{
		{
			description: "slice",
			field:       ID,
			source: []Foo{
				{
					ID:   1,
					Name: "abc",
				},
				{
					ID:   12,
					Name: "xyz",
				},
			},
			expect: []interface{}{1, 12},
		},
		{
			description: "slice item pointer",
			field:       Name,
			source: []*Foo{
				{
					ID:   1,
					Name: "abc",
				},
				{
					ID:   12,
					Name: "xyz",
				},
			},
			expect: []interface{}{"abc", "xyz"},
		},
	}

	for _, testCase := range testCases {
		aSlice := NewSlice(reflect.TypeOf(testCase.source))
		holderAddr := EnsurePointer(testCase.source)
		actual := make([]interface{}, 0)

		for i := 0; i < aSlice.Len(holderAddr); i++ {
			item := aSlice.ValuePointerAt(holderAddr, i)
			if testCase.field == nil {
				actual = append(actual, item)
				continue
			}
			val := testCase.field.Interface(AsPointer(item))
			actual = append(actual, val)
		}
		assert.EqualValues(t, testCase.expect, actual)

	}

}

func TestSlice_Appender(t *testing.T) {
	type Foo struct {
		ID   int
		Name string
	}
	aSlice := NewSlice(reflect.TypeOf([]*Foo{}))
	var foos []*Foo
	appender := aSlice.Appender(unsafe.Pointer(&foos))
	for i := 0; i < 20; i++ {
		fooPtr := &Foo{ID: i, Name: "foo"}
		appender.Append(fooPtr)
	}
	assert.EqualValues(t, 20, len(foos))
	for i := 0; i < 20; i++ {
		assert.EqualValues(t, i, foos[i].ID)
		assert.EqualValues(t, "foo", foos[i].Name)
	}
}

func TestSlice_AppenderAdd(t *testing.T) {
	type Foo struct {
		ID    int
		Price float64
		Name  string
	}

	aSlice := NewSlice(reflect.TypeOf([]*Foo{}))
	foosLen := 20
	var foos []*Foo
	appender := aSlice.Appender(unsafe.Pointer(&foos))

	for i := 0; i < foosLen; i++ {
		fooPtr, ok := appender.Add().(*Foo)
		assert.True(t, ok)
		//foo := Foo{ID: i, Price: float64(i), Name: "foo name"}
		//*fooPtr = foo
		fooPtr.ID = 1
		fooPtr.Price = float64(i)
		fooPtr.Name = "foo name"
	}

	fmt.Printf("%v\n", foos)
	assert.EqualValues(t, foosLen, len(foos))
	for i := 0; i < foosLen; i++ {
		assert.EqualValues(t, i, foos[i].ID)
		assert.EqualValues(t, float64(i), foos[i].Price)
		assert.EqualValues(t, "foo name", foos[i].Name)
	}
}

func TestAppender_Append(t *testing.T) {

	type Foo struct {
		ID   int
		Name string
	}

	var testCases = []struct {
		description string
		setter      func(interface{}, int)
		itemType    reflect.Type
		expect      []interface{}
	}{

		{
			description: "[]int",
			itemType:    reflect.TypeOf(0),
			setter: func(ptr interface{}, val int) {
				item := ptr.(*int)
				*item = val
			},
			expect: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			description: "[]*int",
			itemType:    reflect.PtrTo(reflect.TypeOf(0)),
			setter: func(ptr interface{}, val int) {
				item := ptr.(**int)
				*item = &val
			},
			expect: []interface{}{ptrTo(0), ptrTo(1), ptrTo(2), ptrTo(3), ptrTo(4), ptrTo(5), ptrTo(6), ptrTo(7), ptrTo(8), ptrTo(9)},
		},
		{
			description: "[]Foo",
			itemType:    reflect.TypeOf(Foo{}),
			setter: func(ptr interface{}, val int) {
				item := ptr.(*Foo)
				item.ID = val
			},
			expect: []interface{}{
				Foo{ID: 0},
				Foo{ID: 1},
				Foo{ID: 2},
				Foo{ID: 3},
				Foo{ID: 4},
				Foo{ID: 5},
			},
		},
		{
			description: "[]*Foo",
			itemType:    reflect.TypeOf(&Foo{}),
			setter: func(ptr interface{}, val int) {
				item := ptr.(**Foo)
				*item = &Foo{ID: val}
			},
			expect: []interface{}{
				&Foo{ID: 0},
				&Foo{ID: 1},
				&Foo{ID: 2},
				&Foo{ID: 3},
				&Foo{ID: 4},
				&Foo{ID: 5},
			},
		},
	}

	for _, testCase := range testCases {

		sliceType := reflect.SliceOf(testCase.itemType)
		actualSlice := reflect.New(sliceType)
		aSlice := NewSlice(sliceType, UseItemAddrOpt(true))
		appender := aSlice.Appender(unsafe.Pointer(actualSlice.Elem().UnsafeAddr()))
		for i := 0; i < len(testCase.expect); i++ {
			item := reflect.New(testCase.itemType)
			testCase.setter(item.Interface(), i)
			appender.Append(item.Interface())
		}

		for i, expect := range testCase.expect {
			actual := actualSlice.Elem().Index(i).Interface()
			assert.EqualValues(t, expect, actual, fmt.Sprintf("[%v]: ", i)+testCase.description)
		}

	}

}

func BenchmarkSlice_Index_Native(b *testing.B) {
	var sliceSize = 10
	var ints = make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		ints[i] = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for i := 0; i < sliceSize; i++ {
			v := ints[i]
			if v != i {
				b.Fail()
			}
		}
	}
}

func BenchmarkSlice_Index_Xunsafe(b *testing.B) {
	var sliceSize = 10
	var ints = make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		ints[i] = i
	}
	ptr := unsafe.Pointer(&ints)
	aSlice := NewSlice(reflect.TypeOf(ints))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for j := 0; j < sliceSize; j++ {
			val := aSlice.AddrAt(ptr, j).(*int)
			if *val != j {
				b.Fail()
			}
		}
	}
}

func BenchmarkSlice_Index_Reflect(b *testing.B) {
	var sliceSize = 10
	var ints = make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		ints[i] = i
	}
	aSlice := reflect.ValueOf(ints)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for i := 0; i < sliceSize; i++ {
			v := aSlice.Index(i).Interface()
			if v != i {
				b.Fail()
			}
		}
	}
}

func BenchmarkAppender_Append_Xunsafe(b *testing.B) {
	aSlice := NewSlice(reflect.TypeOf([]int{}))
	b.ReportAllocs()
	for k := 0; k < b.N; k++ {
		var ints []int
		appender := aSlice.Appender(unsafe.Pointer(&ints))
		for i := 0; i < 100; i++ {
			z := i
			appender.Append(unsafe.Pointer(&z))
		}
		if len(ints) != 100 {
			b.Fail()
		}
	}
}

func BenchmarkAppender_Append_Reflect(b *testing.B) {
	b.ReportAllocs()
	for k := 0; k < b.N; k++ {
		var ints []int
		slice := reflect.ValueOf(&ints)
		for i := 0; i < 100; i++ {
			z := i
			slice.Elem().Set(reflect.Append(slice.Elem(), reflect.ValueOf(z)))
		}
		if len(ints) != 100 {
			b.Fail()
		}
	}
}

func BenchmarkAppender_Append_Native(b *testing.B) {
	b.ReportAllocs()
	for k := 0; k < b.N; k++ {
		var ints []int
		for i := 0; i < 100; i++ {
			z := i
			ints = append(ints, z)
		}
		if len(ints) != 100 {
			b.Fail()
		}
	}
}

func ptrTo(i interface{}) interface{} {
	ptr := reflect.New(reflect.TypeOf(i))
	ptr.Elem().Set(reflect.ValueOf(i))
	return ptr.Interface()
}
