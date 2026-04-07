package converter

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func TestUnifyPointerResultStableAcrossGC(t *testing.T) {
	unified, err := Unify(reflect.PtrTo(reflect.TypeOf(0)), reflect.TypeOf(0))
	if err != nil {
		t.Fatalf("unexpected unify error: %v", err)
	}

	value := 7
	ptr, err := unified.Y(unsafe.Pointer(&value))
	if err != nil {
		t.Fatalf("unexpected unify fn error: %v", err)
	}

	runtime.GC()

	actual := *(**int)(ptr)
	if actual == nil {
		t.Fatalf("expected non-nil pointer")
	}
	if *actual != 7 {
		t.Fatalf("expected 7, got %d", *actual)
	}
}
