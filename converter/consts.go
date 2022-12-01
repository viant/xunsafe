package converter

import "unsafe"

var truePtr = unsafe.Pointer(boolPtr(true))
var falsePtr = unsafe.Pointer(boolPtr(false))

var intZeroPtr = unsafe.Pointer(intPtr(0))
var intOnePtr = unsafe.Pointer(intPtr(1))
var uintZeroPtr = unsafe.Pointer(uintPtr(0))
var uintOnePtr = unsafe.Pointer(uintPtr(1))

var int8ZeroPtr = unsafe.Pointer(int8Ptr(0))
var int8OnePtr = unsafe.Pointer(int8Ptr(1))
var uint8ZeroPtr = unsafe.Pointer(uint8Ptr(0))
var uint8OnePtr = unsafe.Pointer(uint8Ptr(1))

var int16ZeroPtr = unsafe.Pointer(int16Ptr(0))
var int16OnePtr = unsafe.Pointer(int16Ptr(1))
var uint16ZeroPtr = unsafe.Pointer(uint16Ptr(0))
var uint16OnePtr = unsafe.Pointer(uint16Ptr(1))

var int32ZeroPtr = unsafe.Pointer(int32Ptr(0))
var int32OnePtr = unsafe.Pointer(int32Ptr(1))
var uint32ZeroPtr = unsafe.Pointer(uint32Ptr(0))
var uint32OnePtr = unsafe.Pointer(uint32Ptr(1))

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func uintPtr(i uint) *uint {
	return &i
}

func int8Ptr(i int8) *int8 {
	return &i
}

func uint8Ptr(i uint8) *uint8 {
	return &i
}

func int16Ptr(i int16) *int16 {
	return &i
}

func uint16Ptr(i uint16) *uint16 {
	return &i
}

func int32Ptr(i int32) *int32 {
	return &i
}

func uint32Ptr(i uint32) *uint32 {
	return &i
}
