package slices

import (
	"unsafe"
)

type intersectFn func(sptr, optr unsafe.Pointer) interface{}

// Intersect returns a slice with the intersection of slice and other.
// When the slices are the same, slice is returned.
func Intersect(slice, other interface{}) interface{} {
	fn, ok := intersectOf(slice, other)
	if fn == nil {
		panic("slice is not supported slice type")
	}
	if !ok {
		panic("other is not the same type as slice")
	}

	sptr := noescape(ptrOf(slice))
	optr := noescape(ptrOf(other))
	return fn(sptr, optr)
}

func intersectOf(slice, other interface{}) (intersectFn, bool) {
	switch slice.(type) {
	case []string:
		_, ok := other.([]string)
		return stringIntersect, ok
	case []int:
		_, ok := other.([]int)
		return intIntersect, ok
	case []int8:
		_, ok := other.([]int8)
		return int8Intersect, ok
	case []int16:
		_, ok := other.([]int16)
		return int16Intersect, ok
	case []int32:
		_, ok := other.([]int32)
		return int32Intersect, ok
	case []int64:
		_, ok := other.([]int64)
		return int64Intersect, ok
	case []uint:
		_, ok := other.([]uint)
		return uintIntersect, ok
	case []uint8:
		_, ok := other.([]uint8)
		return uint8Intersect, ok
	case []uint16:
		_, ok := other.([]uint16)
		return uint16Intersect, ok
	case []uint32:
		_, ok := other.([]uint32)
		return uint32Intersect, ok
	case []uint64:
		_, ok := other.([]uint64)
		return uint64Intersect, ok
	case []float32:
		_, ok := other.([]float32)
		return float32Intersect, ok
	case []float64:
		_, ok := other.([]float64)
		return float64Intersect, ok
	}
	return nil, false
}
