package xunsafe

import (
	"reflect"
	"unsafe"
)

// IsZero reports whether the value at ptr is a zero value for kind.
// For reflect.Ptr, elem controls pointed-value kind checks; if elem is
// reflect.Invalid, a non-nil pointer is considered non-zero.
func IsZero(ptr unsafe.Pointer, kind reflect.Kind, elem reflect.Kind) bool {
	switch kind {
	case reflect.String:
		return AsString(ptr) == ""
	case reflect.Bool:
		return !AsBool(ptr)
	case reflect.Int:
		return AsInt(ptr) == 0
	case reflect.Int8:
		return AsInt8(ptr) == 0
	case reflect.Int16:
		return AsInt16(ptr) == 0
	case reflect.Int32:
		return AsInt32(ptr) == 0
	case reflect.Int64:
		return AsInt64(ptr) == 0
	case reflect.Uint:
		return AsUint(ptr) == 0
	case reflect.Uint8:
		return AsUint8(ptr) == 0
	case reflect.Uint16:
		return AsUint16(ptr) == 0
	case reflect.Uint32:
		return AsUint32(ptr) == 0
	case reflect.Uint64:
		return AsUint64(ptr) == 0
	case reflect.Uintptr:
		return AsUintptr(ptr) == 0
	case reflect.Float32:
		return AsFloat32(ptr) == 0
	case reflect.Float64:
		return AsFloat64(ptr) == 0
	case reflect.Ptr:
		p := *(*unsafe.Pointer)(ptr)
		if p == nil {
			return true
		}
		if elem == reflect.Invalid {
			return false
		}
		return IsZero(p, elem, reflect.Invalid)
	}
	return false
}
