package slices

import (
	"unsafe"
)

type exceptFn func(sptr, optr unsafe.Pointer) interface{}

// Except returns a slice with the elements of slice that are not in other.
// When the slices different, slice is returned.
func Except(slice, other interface{}) interface{} {
	fn, ok := exceptOf(slice, other)
	if fn == nil {
		panic("slice is not a supported slice type")
	}
	if !ok {
		panic("other is not the same type as slice")
	}

	sptr := noescape(ptrOf(slice))
	optr := noescape(ptrOf(other))
	return fn(sptr, optr)
}

func exceptOf(slice, other interface{}) (exceptFn, bool) {
	switch slice.(type) {
	case []string:
		_, ok := other.([]string)
		return stringExcept, ok
	case []int:
		_, ok := other.([]int)
		return intExcept, ok
	case []int8:
		_, ok := other.([]int8)
		return int8Except, ok
	case []int16:
		_, ok := other.([]int16)
		return int16Except, ok
	case []int32:
		_, ok := other.([]int32)
		return int32Except, ok
	case []int64:
		_, ok := other.([]int64)
		return int64Except, ok
	case []uint:
		_, ok := other.([]uint)
		return uintExcept, ok
	case []uint8:
		_, ok := other.([]uint8)
		return uint8Except, ok
	case []uint16:
		_, ok := other.([]uint16)
		return uint16Except, ok
	case []uint32:
		_, ok := other.([]uint32)
		return uint32Except, ok
	case []uint64:
		_, ok := other.([]uint64)
		return uint64Except, ok
	case []float32:
		_, ok := other.([]float32)
		return float32Except, ok
	case []float64:
		_, ok := other.([]float64)
		return float64Except, ok
	}
	return nil, false
}
