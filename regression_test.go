package xunsafe

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func TestRefPointerStableAcrossGC(t *testing.T) {
	value := 42
	ptr := unsafe.Pointer(&value)

	ref := RefPointer(ptr)
	runtime.GC()

	actual := *(**int)(ref)
	if actual == nil {
		t.Fatalf("expected non-nil pointer")
	}
	if *actual != 42 {
		t.Fatalf("expected 42, got %d", *actual)
	}
}

func TestEnsureAddressPointerAllocatesElement(t *testing.T) {
	var actual *int

	cell := EnsureAddressPointer(unsafe.Pointer(&actual), reflect.TypeOf(actual))
	value := (*int)(*cell)
	*value = 99
	runtime.GC()

	if actual == nil {
		t.Fatalf("expected allocated pointer")
	}
	if *actual != 99 {
		t.Fatalf("expected 99, got %d", *actual)
	}
}

func TestFieldEnsurePointerAllocatesElement(t *testing.T) {
	type holder struct {
		Value *int
	}

	field := FieldByName(reflect.TypeOf(holder{}), "Value")
	instance := &holder{}

	ptr := field.EnsurePointer(unsafe.Pointer(instance))
	value := (*int)(ptr)
	*value = 123
	runtime.GC()

	if instance.Value == nil {
		t.Fatalf("expected field pointer to be initialized")
	}
	if *instance.Value != 123 {
		t.Fatalf("expected 123, got %d", *instance.Value)
	}
}
