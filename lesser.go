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

func stringLesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]string)(ptr)
		return v[i] < v[j]
	}
}

func intLesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int)(ptr)
		return v[i] < v[j]
	}
}

func int8Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int8)(ptr)
		return v[i] < v[j]
	}
}

func int16Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int16)(ptr)
		return v[i] < v[j]
	}
}

func int32Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int32)(ptr)
		return v[i] < v[j]
	}
}

func int64Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int64)(ptr)
		return v[i] < v[j]
	}
}

func uintLesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint)(ptr)
		return v[i] < v[j]
	}
}

func uint8Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint8)(ptr)
		return v[i] < v[j]
	}
}

func uint16Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint16)(ptr)
		return v[i] < v[j]
	}
}

func uint32Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint32)(ptr)
		return v[i] < v[j]
	}
}

func uint64Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint64)(ptr)
		return v[i] < v[j]
	}
}

func float32Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]float32)(ptr)
		return v[i] < v[j]
	}
}

func float64Lesser(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]float64)(ptr)
		return v[i] < v[j]
	}
}
