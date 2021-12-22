package xunsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

//Index returns item pointer for supplied indexByAddr (see useItemAddr slice option)
type Index func(index uintptr) unsafe.Pointer

//Slice represents a slice
type (
	Slice struct {
		reflect.Type
		useItemAddr bool //useItemAddr flag instructs implementation to use item address as **T for []*T or *T for []T
		//otherwise item would use *T for a slice defined as []T or []*T,
		itemSize uintptr
	}
	//Appender represents a slice appender
	Appender struct {
		cap          int
		size         int
		slice        *Slice
		header       *reflect.SliceHeader
		reflectSlice reflect.Value
		itemType     reflect.Type
		ptr          unsafe.Pointer
		indexAddr    Index
	}
)

//Range call visit callback for each slice element , to terminate visit should return false
//use useItemAddr would use item pointer as *T for a slice defined as []T or []*T,
//otherwise for slice defined as []*T, item would get **T pointer
func (s *Slice) Range(sliceAddress unsafe.Pointer, visit func(index int, item unsafe.Pointer) bool) {
	header := *(*reflect.SliceHeader)(sliceAddress)
	fn := s.Index(sliceAddress)
	for i := 0; i < header.Len; i++ {
		if !visit(i, fn(uintptr(i))) {
			return
		}
	}
}

//Appender returns a slice appender
func (s *Slice) Appender(slicePointer unsafe.Pointer) *Appender {
	header := (*reflect.SliceHeader)(slicePointer)
	return &Appender{slice: s,
		header:    header,
		ptr:       slicePointer,
		indexAddr: s.IndexAddr(slicePointer),
		itemType:  s.Type.Elem(),
	}
}

//Index return slice item
func (s *Slice) Index(sliceAddr unsafe.Pointer) Index {
	header := (*reflect.SliceHeader)(sliceAddr)
	size := s.itemSize
	if s.useItemAddr {
		return func(index uintptr) unsafe.Pointer {
			offset := header.Data - uintptr(sliceAddr)
			return unsafe.Add(sliceAddr, offset+index*size)
		}
	}
	return func(index uintptr) unsafe.Pointer {
		offset := header.Data - uintptr(sliceAddr)
		return *(*unsafe.Pointer)(unsafe.Add(sliceAddr, offset+index*size))
	}
}

//IndexAddr return slice item addr pointer
func (s *Slice) IndexAddr(sliceAddr unsafe.Pointer) Index {
	header := (*reflect.SliceHeader)(sliceAddr)
	size := s.itemSize
	return func(index uintptr) unsafe.Pointer {
		offset := header.Data - uintptr(sliceAddr)
		return unsafe.Add(sliceAddr, offset+index*size)
	}
}

//UseItemAddrOpt option that instructs implementation to use item address as **T for []*T or *T for []T, otherwise *T would be used
type UseItemAddrOpt bool

//NewSlice creates  slice
func NewSlice(aType reflect.Type, options ...interface{}) *Slice {
	switch aType.Kind() {
	case reflect.Slice:
	case reflect.Array:
		panic(fmt.Sprintf("unsupported type: %v", aType.Name()))
	default:
		aType = reflect.SliceOf(aType)
	}
	itemType := aType.Elem()
	size := itemType.Size()

	result := &Slice{
		Type:     aType,
		itemSize: size,
	}
	result.applyOptions(options)
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
func (a *Appender) Append(items ...unsafe.Pointer) {
	itemLen := len(items)
	if a.cap < a.size+itemLen {
		a.grow(itemLen)
	}
	i := 0
	if a.slice.useItemAddr {
	loop1:
		*(*unsafe.Pointer)(a.indexAddr(uintptr(a.size))) = *(*unsafe.Pointer)(items[i])
		a.size++
		i++
		if i < itemLen {
			goto loop1
		}
		a.header.Len = a.size
		return
	}
loop2:
	*(*unsafe.Pointer)(a.indexAddr(uintptr(a.size))) = unsafe.Pointer(reflect.New(a.itemType).Elem().UnsafeAddr())
	*(*unsafe.Pointer)(DereferencePointer(a.indexAddr(uintptr(a.size)))) = *(*unsafe.Pointer)(items[i])
	a.size++
	i++
	if i < itemLen {
		goto loop2
	}
	a.header.Len = a.size
}

//Add grows slice by 1 and returns item pointer (see UseItemAddrOpt)
func (a *Appender) Add() unsafe.Pointer {
	if a.cap < a.size+1 {
		a.grow(1)
	}
	i := 0
	if a.slice.useItemAddr {
		result := (a.indexAddr(uintptr(a.size)))
		a.size++
		a.header.Len = a.size
		return result
	}
	newPtr := reflect.New(a.itemType)
	*(*unsafe.Pointer)(a.indexAddr(uintptr(a.size))) = unsafe.Pointer(newPtr.Pointer())
	result := (DereferencePointer(a.indexAddr(uintptr(a.size))))
	a.size++
	i++
	a.header.Len = a.size
	return result
}

func (a *Appender) grow(by int) {
	cap := (a.cap + by + 1) * 2
	newSlice := reflect.MakeSlice(a.slice.Type, cap, cap)
	if a.cap > 0 {
		reflect.Copy(newSlice, a.reflectSlice)
	}
	a.reflectSlice = newSlice
	a.header.Data = a.reflectSlice.Pointer()
	a.header.Len = cap
	a.header.Cap = cap
	a.cap = cap
}
