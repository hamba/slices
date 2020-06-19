package slices

import (
	"unsafe"
)

type containsFn func(sptr, vptr unsafe.Pointer) bool

// Contains determines if the slice contains val.
func Contains(slice, val interface{}) bool {
	fn, ok := containsOf(slice, val)
	if fn == nil {
		panic("slice is not supported slice type")
	}
	if !ok {
		panic("val is not the same type as slice")
	}

	sptr := noescape(ptrOf(slice))
	vptr := noescape(ptrOf(val))
	return fn(sptr, vptr)
}

func containsOf(slice, val interface{}) (containsFn, bool) {
	switch slice.(type) {
	case []bool:
		_, ok := val.(bool)
		return boolContains, ok
	case []string:
		_, ok := val.(string)
		return stringContains, ok
	case []int:
		_, ok := val.(int)
		return intContains, ok
	case []int8:
		_, ok := val.(int8)
		return int8Contains, ok
	case []int16:
		_, ok := val.(int16)
		return int16Contains, ok
	case []int32:
		_, ok := val.(int32)
		return int32Contains, ok
	case []int64:
		_, ok := val.(int64)
		return int64Contains, ok
	case []uint:
		_, ok := val.(uint)
		return uintContains, ok
	case []uint8:
		_, ok := val.(uint8)
		return uint8Contains, ok
	case []uint16:
		_, ok := val.(uint16)
		return uint16Contains, ok
	case []uint32:
		_, ok := val.(uint32)
		return uint32Contains, ok
	case []uint64:
		_, ok := val.(uint64)
		return uint64Contains, ok
	case []float32:
		_, ok := val.(float32)
		return float32Contains, ok
	case []float64:
		_, ok := val.(float64)
		return float64Contains, ok
	}
	return nil, false
}
