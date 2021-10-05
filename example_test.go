package xunsafe_test

import (
	"fmt"
	"github.com/viant/xunsafe"
	"reflect"
)

func Example_FastReflection() {
	type Foo struct {
		ID   int
		Name string
	}
	fooType := reflect.TypeOf(Foo{})
	fooID := xunsafe.FieldByName(fooType, "ID")
	fooName := xunsafe.FieldByName(fooType, "Name")

	var foos = make([]Foo, 100)
	for i := range foos {
		fooAddr := xunsafe.Addr(&foos[i])
		fooID.SetInt(fooAddr, i)
		fooName.SetString(fooAddr, fmt.Sprintf("name %d", i))
	}

}

func ExampleAddr() {
	type Foo struct {
		ID   int
		Name string
	}
	fooType := reflect.TypeOf(Foo{})
	fooID := xunsafe.FieldByName(fooType, "ID")
	foo := &Foo{ID: 101, Name: "name 101"}

	fooAddr := xunsafe.Addr(foo)
	*(fooID.Addr(fooAddr).(*int)) = 201
	fmt.Printf("foo.ID: %v\n", foo.ID) //prints 201
}
