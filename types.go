package xunsafe

import (
	"reflect"
	"sync"
	"unsafe"
)

//NOTE that is API is not stable and may change in the future.

//go:linkname typelinks reflect.typelinks
func typelinks() ([]unsafe.Pointer, [][]int32)

//go:linkname rtypeOff reflect.rtypeOff
func rtypeOff(base unsafe.Pointer, offset int32) unsafe.Pointer

var packageTypes = make(map[string][]reflect.Type)
var indexedTypes = make(map[string]reflect.Type)
var once sync.Once

// LookupType returns linked in reflect.Type for the given name.
func LookupType(name string) reflect.Type {
	loadLinkedinTypes()
	return indexedTypes[name]
}

func loadLinkedinTypes() {
	once.Do(func() {
		bases, offs := typelinks()
		var result []reflect.Type
		for i, base := range bases {
			for _, off := range offs[i] {
				ptr := rtypeOff(base, off)
				typ := reflect.TypeOf(*(*interface{})(unsafe.Pointer(&struct {
					t unsafe.Pointer
					d uintptr
				}{ptr, 0})))

				result = append(result, typ)
				if typ.Kind() == reflect.Ptr {
					typ = typ.Elem()
				}
				if typ.Name() == "" {
					continue
				}
				packageTypes[typ.PkgPath()] = append(packageTypes[typ.Name()], typ)
				var typeName string
				if typ.PkgPath() == "" {
					typeName = typ.Name()
				} else {
					typeName = typ.PkgPath() + "." + typ.Name()
				}
				indexedTypes[typeName] = typ
			}
		}
	})
}

// PackageTypes returns all types for the given package name.
func PackageTypes(name string) []reflect.Type {
	loadLinkedinTypes()
	return packageTypes[name]
}
