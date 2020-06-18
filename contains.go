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

func boolContains(sptr, vptr unsafe.Pointer) bool {
	v := *(*bool)(vptr)
	for _, vv := range *(*[]bool)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func stringContains(sptr, vptr unsafe.Pointer) bool {
	v := *(*string)(vptr)
	for _, vv := range *(*[]string)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func intContains(sptr, vptr unsafe.Pointer) bool {
	v := *(*int)(vptr)
	for _, vv := range *(*[]int)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func int8Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*int8)(vptr)
	for _, vv := range *(*[]int8)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func int16Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*int16)(vptr)
	for _, vv := range *(*[]int16)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func int32Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*int32)(vptr)
	for _, vv := range *(*[]int32)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func int64Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*int64)(vptr)
	for _, vv := range *(*[]int64)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func uintContains(sptr, vptr unsafe.Pointer) bool {
	v := *(*uint)(vptr)
	for _, vv := range *(*[]uint)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func uint8Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*uint8)(vptr)
	for _, vv := range *(*[]uint8)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func uint16Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*uint16)(vptr)
	for _, vv := range *(*[]uint16)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func uint32Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*uint32)(vptr)
	for _, vv := range *(*[]uint32)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func uint64Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*uint64)(vptr)
	for _, vv := range *(*[]uint64)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func float32Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*float32)(vptr)
	for _, vv := range *(*[]float32)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}

func float64Contains(sptr, vptr unsafe.Pointer) bool {
	v := *(*float64)(vptr)
	for _, vv := range *(*[]float64)(sptr) {
		if vv == v {
			return true
		}
	}
	return false
}
