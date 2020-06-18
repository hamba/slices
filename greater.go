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

func stringGreater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]string)(ptr)
		return v[i] > v[j]
	}
}

func intGreater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int)(ptr)
		return v[i] > v[j]
	}
}

func int8Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int8)(ptr)
		return v[i] > v[j]
	}
}

func int16Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int16)(ptr)
		return v[i] > v[j]
	}
}

func int32Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int32)(ptr)
		return v[i] > v[j]
	}
}

func int64Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]int64)(ptr)
		return v[i] > v[j]
	}
}

func uintGreater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint)(ptr)
		return v[i] > v[j]
	}
}

func uint8Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint8)(ptr)
		return v[i] > v[j]
	}
}

func uint16Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint16)(ptr)
		return v[i] > v[j]
	}
}

func uint32Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint32)(ptr)
		return v[i] > v[j]
	}
}

func uint64Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]uint64)(ptr)
		return v[i] > v[j]
	}
}

func float32Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]float32)(ptr)
		return v[i] > v[j]
	}
}

func float64Greater(ptr unsafe.Pointer) func(i, j int) bool {
	return func(i, j int) bool {
		v := *(*[]float64)(ptr)
		return v[i] > v[j]
	}
}
