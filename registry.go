package xunsafe

import (
	"reflect"
	"sync"
)

var _registry = sync.Map{}


//Register register custom type with pointer getter
func lookup(aType reflect.Type) Getter {
	getter, ok := _registry.Load(aType)
	if ! ok {
		return nil
	}
	return getter.(Getter)
}


//Register register custom type with pointer getter
func Register(aType reflect.Type, getter Getter) {
	_registry.Store(aType, getter)
}
