package slices

import (
	"unsafe"
)

type lesserFn func(unsafe.Pointer) func(i, j int) bool

// LesserOf returns a lesser function for slice sorting.
func LesserOf(slice interface{}) func(i, j int) bool {
	fn := lesserOf(slice)
	if fn == nil {
		panic("slice is not supported slice type")
	}

	ptr := noescape(ptrOf(slice))
	return fn(ptr)
}

func lesserOf(slice interface{}) lesserFn {
	switch slice.(type) {
	case []string:
		return stringLesser
	case []int:
		return intLesser
	case []int8:
		return int8Lesser
	case []int16:
		return int16Lesser
	case []int32:
		return int32Lesser
	case []int64:
		return int64Lesser
	case []uint:
		return uintLesser
	case []uint8:
		return uint8Lesser
	case []uint16:
		return uint16Lesser
	case []uint32:
		return uint32Lesser
	case []uint64:
		return uint64Lesser
	case []float32:
		return float32Lesser
	case []float64:
		return float64Lesser
	}
	return nil
}
