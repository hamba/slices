package slices

import (
	"unsafe"
)

type greaterFn func(unsafe.Pointer) func(i, j int) bool

// GreaterOf returns a greater function for slice sorting.
func GreaterOf(slice interface{}) func(i, j int) bool {
	fn := greaterOf(slice)
	if fn == nil {
		panic("slice is not supported slice type")
	}

	ptr := noescape(ptrOf(slice))
	return fn(ptr)
}

func greaterOf(slice interface{}) greaterFn {
	switch slice.(type) {
	case []string:
		return stringGreater
	case []int:
		return intGreater
	case []int8:
		return int8Greater
	case []int16:
		return int16Greater
	case []int32:
		return int32Greater
	case []int64:
		return int64Greater
	case []uint:
		return uintGreater
	case []uint8:
		return uint8Greater
	case []uint16:
		return uint16Greater
	case []uint32:
		return uint32Greater
	case []uint64:
		return uint64Greater
	case []float32:
		return float32Greater
	case []float64:
		return float64Greater
	}
	return nil
}
