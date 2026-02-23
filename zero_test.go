package xunsafe

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestIsZero_Primitives(t *testing.T) {
	var s string
	if !IsZero(unsafe.Pointer(&s), reflect.String, reflect.Invalid) {
		t.Fatalf("expected empty string to be zero")
	}
	s = "x"
	if IsZero(unsafe.Pointer(&s), reflect.String, reflect.Invalid) {
		t.Fatalf("expected non-empty string to be non-zero")
	}

	var i int64
	if !IsZero(unsafe.Pointer(&i), reflect.Int64, reflect.Invalid) {
		t.Fatalf("expected zero int64 to be zero")
	}
	i = 9
	if IsZero(unsafe.Pointer(&i), reflect.Int64, reflect.Invalid) {
		t.Fatalf("expected non-zero int64 to be non-zero")
	}

	var f float32
	if !IsZero(unsafe.Pointer(&f), reflect.Float32, reflect.Invalid) {
		t.Fatalf("expected zero float32 to be zero")
	}
	f = 1.25
	if IsZero(unsafe.Pointer(&f), reflect.Float32, reflect.Invalid) {
		t.Fatalf("expected non-zero float32 to be non-zero")
	}
}

func TestIsZero_Pointer(t *testing.T) {
	var p *int
	if !IsZero(unsafe.Pointer(&p), reflect.Ptr, reflect.Int) {
		t.Fatalf("expected nil pointer to be zero")
	}

	v := 0
	p = &v
	if !IsZero(unsafe.Pointer(&p), reflect.Ptr, reflect.Int) {
		t.Fatalf("expected pointer-to-zero int to be zero")
	}

	v = 7
	if IsZero(unsafe.Pointer(&p), reflect.Ptr, reflect.Int) {
		t.Fatalf("expected pointer-to-non-zero int to be non-zero")
	}
}
