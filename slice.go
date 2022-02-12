package xunsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

//Slice represents a slice
type (
	Slice struct {
		reflect.Type
		isPointer   bool
		useItemAddr bool //useItemAddr flag instructs implementation to use item address as **T for []*T or *T for []T
		//otherwise item would use *T for a slice defined as []T or []*T,
		itemSize uintptr
		rtype    *rtype
		rtypePtr *rtype
	}
	//Appender represents a slice appender
	Appender struct {
		cap          int
		len          int
		slice        *Slice
		header       *reflect.SliceHeader
		reflectSlice reflect.Value
		itemType     reflect.Type
		ptr          unsafe.Pointer
	}
)

//PointerAt return slice item pointer for supplied index
func (s *Slice) PointerAt(sliceAddr unsafe.Pointer, index uintptr) unsafe.Pointer {
	header := (*reflect.SliceHeader)(sliceAddr)
	return unsafe.Add(unsafe.Pointer(header.Data), index*s.itemSize)
}

//Len return slice length
func (s *Slice) Len(slicePtr unsafe.Pointer) int {
	header := *(*reflect.SliceHeader)(slicePtr)
	return header.Len
}

//Cap return slice capacity
func (s *Slice) Cap(slicePtr unsafe.Pointer) int {
	header := *(*reflect.SliceHeader)(slicePtr)
	return header.Cap
}

//ValuePointerAt return value pointer *T, for both []T and []*T slice definition
func (s *Slice) ValuePointerAt(slicePtr unsafe.Pointer, index int) interface{} {
	p := s.PointerAt(slicePtr, uintptr(index))
	if s.isPointer {
		p = DerefPointer(p)
		return asInterface(p, s.rtype, false)
	}
	return asInterface(p, s.rtypePtr, false)
}

//ValueAt return slice item for supplied index
func (s *Slice) ValueAt(slicePtr unsafe.Pointer, index int) interface{} {
	p := s.PointerAt(slicePtr, uintptr(index))
	return asInterface(p, s.rtype, false)
}

//AddrAt return slice item addr for supplied index
func (s *Slice) AddrAt(slicePtr unsafe.Pointer, index int) interface{} {
	return asInterface(s.PointerAt(slicePtr, uintptr(index)), s.rtypePtr, false)
}

//Range call visit callback for each slice element , to terminate visit should return false
//use useItemAddr would use item pointer as *T for a slice defined as []T or []*T,
//otherwise for slice defined as []*T, item would get **T pointer
func (s *Slice) Range(slicePtr unsafe.Pointer, visit func(index int, item interface{}) bool) {
	header := *(*reflect.SliceHeader)(slicePtr)
	for i := 0; i < header.Len; i++ {
		val := s.ValueAt(slicePtr, i)
		if !visit(i, val) {
			return
		}
	}
}

//Appender returns a slice appender
func (s *Slice) Appender(slicePointer unsafe.Pointer) *Appender {
	header := (*reflect.SliceHeader)(slicePointer)
	return &Appender{slice: s,
		header:   header,
		ptr:      slicePointer,
		itemType: s.Type.Elem(),
		cap:      header.Cap,
		len:      header.Len,
	}
}

func (s *Slice) initTypes() {
	s.isPointer = s.Type.Elem().Kind() == reflect.Ptr
	ptrValue := reflect.New(s.Type.Elem())
	ptr := ptrValue.Interface()
	val := ptrValue.Elem().Interface()
	s.rtype = ((*emptyInterface)(unsafe.Pointer(&val))).typ
	s.rtypePtr = ((*emptyInterface)(unsafe.Pointer(&ptr))).typ
}

//UseItemAddrOpt option that instructs implementation to use item address as **T for []*T or *T for []T, otherwise *T would be used
type UseItemAddrOpt bool

//NewSlice creates  slice
func NewSlice(sliceType reflect.Type, options ...interface{}) *Slice {
	switch sliceType.Kind() {
	case reflect.Slice:
	case reflect.Array:
		panic(fmt.Sprintf("unsupported type: %v", sliceType.Name()))
	default:
		sliceType = reflect.SliceOf(sliceType)
	}
	itemType := sliceType.Elem()
	result := &Slice{
		Type:     sliceType,
		itemSize: itemType.Size(),
	}
	result.applyOptions(options)
	result.initTypes()
	return result
}

func (s *Slice) applyOptions(options []interface{}) {
	if len(options) == 0 {
		s.useItemAddr = s.Type.Elem().Kind() != reflect.Ptr
	}
	for _, opt := range options {
		if useItemAddr, ok := opt.(UseItemAddrOpt); ok {
			s.useItemAddr = bool(useItemAddr) || s.Type.Elem().Kind() != reflect.Ptr
		}
	}
}

//Append appends items to a slice
func (a *Appender) Append(items ...interface{}) {
	itemLen := len(items)
	if a.cap < a.len+itemLen {
		a.grow(itemLen)
	}
	i := 0

	if a.slice.useItemAddr {
	loop1:
		sourcePtr := AsPointer(items[i])
		ptr := a.slice.PointerAt(a.ptr, uintptr(a.len))
		*(*unsafe.Pointer)(ptr) = *(*unsafe.Pointer)(sourcePtr)
		a.len++
		i++
		if i < itemLen {
			goto loop1
		}
		a.header.Len = a.len
		return
	}
loop2:
	sourcePtr := AsPointer(items[i])
	ptr := a.slice.PointerAt(a.ptr, uintptr(a.len))
	*(*unsafe.Pointer)(ptr) = sourcePtr
	a.len++
	i++
	if i < itemLen {
		goto loop2
	}
	a.header.Len = a.len
}

//Add grows slice by 1 and returns item pointer (see UseItemAddrOpt)
func (a *Appender) Add() interface{} {
	if a.cap < a.len+1 {
		a.grow(1)
	}
	ptr := a.slice.PointerAt(a.ptr, uintptr(a.len))
	if a.slice.useItemAddr {
		a.len++
		a.header.Len = a.len
		return asInterface(ptr, a.slice.rtypePtr, false)
	}
	if a.slice.isPointer {
		nPtr := reflect.New(a.itemType.Elem())
		*(*unsafe.Pointer)(ptr) = unsafe.Pointer(nPtr.Pointer())
	}
	itemPtr := EnsureAddressPointer(ptr)
	a.len++
	a.header.Len = a.len
	return asInterface(*itemPtr, a.slice.rtype, false)
}

func (a *Appender) grow(by int) {
	cap := a.cap + by
	if a.cap > 0 {
		cap = (a.cap + by) * 2
	}
	newSlice := reflect.MakeSlice(a.slice.Type, cap, cap)
	if a.cap > 0 {
		reflect.Copy(newSlice, a.reflectSlice)
	}
	a.reflectSlice = newSlice
	a.header.Data = newSlice.Pointer()
	a.header.Len = cap
	a.header.Cap = cap
	a.cap = cap
}
