package xunsafe

import (
	"reflect"
	"unsafe"
)

//flag copied from reflect pacakge
type flag uintptr

const (
	flagKindWidth        = 5 // there are 27 kinds
	flagKindMask    flag = 1<<flagKindWidth - 1
	flagStickyRO    flag = 1 << 5
	flagEmbedRO     flag = 1 << 6
	flagIndir       flag = 1 << 7
	flagAddr        flag = 1 << 8
	flagMethod      flag = 1 << 9
	flagMethodShift      = 10
	flagRO          flag = flagStickyRO | flagEmbedRO
)

//rtype copied from reflect package
type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      uint8   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte // garbage collection data
	str       int32 // string form
	ptrToThis int32 // type for pointer to this type, may be zero
}

func (t *rtype) Size() uintptr { return t.size }

func (t *rtype) Align() int { return int(t.align) }

func (t *rtype) FieldAlign() int { return int(t.fieldAlign) }

func (t *rtype) Kind() reflect.Kind { return reflect.Kind(t.kind & kindMask) }

func (t *rtype) pointers() bool { return t.ptrdata != 0 }

func (t *rtype) common() *rtype { return t }

//emptyInterface copied from reflect package
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

//value is simplied reflect value
type value struct {
	typ  *rtype
	ptr  unsafe.Pointer
	flag flag
}

func valuePointer(v *reflect.Value) unsafe.Pointer {
	return (*value)(unsafe.Pointer(v)).ptr
}

const (
	kindDirectIface = 1 << 5
	kindGCProg      = 1 << 6 // Type.gc points to GC program
	kindMask        = (1 << 5) - 1
)

//AsPointer returns a  pointer for an interface (interface has to store a pointer)
func asInterface(ptr unsafe.Pointer, rtype *rtype, canDeref bool) (v interface{}) {
	empty := (*emptyInterface)(unsafe.Pointer(&v))
	empty.word = ptr
	if rtype.kind&kindDirectIface != 0 && canDeref {
		empty.word = *(*unsafe.Pointer)(ptr)
	}
	empty.typ = rtype
	return v
}
