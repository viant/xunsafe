package xunsafe

import (
	"reflect"
	"sync"
)

type registry struct {
	reg map[reflect.Type]Getter
	mux sync.RWMutex
}

func (r *registry) Register(aType reflect.Type, getter Getter) {
	r.mux.Lock()
	r.mux.Unlock()
	r.reg[aType] = getter
}

func (r *registry) Lookup(aType reflect.Type) Getter {
	return r.reg[aType]
}

var _registry = registry{reg: make(map[reflect.Type]Getter)}

//Register register custom type with pointer getter
func Register(aType reflect.Type, getter Getter) {
	_registry.Register(aType, getter)
}
