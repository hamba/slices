package slices_test

import (
	"testing"

	"github.com/hamba/slices"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice interface{}
		val   interface{}
		want  bool
	}{
		{
			name:  "bool contains",
			slice: []bool{false, false, true},
			val:   true,
			want:  true,
		},
		{
			name:  "bool not contains",
			slice: []bool{false, false, false},
			val:   true,
			want:  false,
		},
		{
			name:  "string contains",
			slice: []string{"foo", "bar", "baz", "bat"},
			val:   "bat",
			want:  true,
		},
		{
			name:  "string not contains",
			slice: []string{"foo", "bar", "baz", "bat"},
			val:   "test",
			want:  false,
		},
		{
			name:  "int contains",
			slice: []int{1, 2, 3, 4},
			val:   4,
			want:  true,
		},
		{
			name:  "int not contains",
			slice: []int{1, 2, 3, 4},
			val:   5,
			want:  false,
		},
		{
			name:  "int8 contains",
			slice: []int8{1, 2, 3, 4},
			val:   int8(4),
			want:  true,
		},
		{
			name:  "int8 not contains",
			slice: []int8{1, 2, 3, 4},
			val:   int8(5),
			want:  false,
		},
		{
			name:  "int16 contains",
			slice: []int16{1, 2, 3, 4},
			val:   int16(4),
			want:  true,
		},
		{
			name:  "int16 not contains",
			slice: []int16{1, 2, 3, 4},
			val:   int16(5),
			want:  false,
		},
		{
			name:  "int32 contains",
			slice: []int32{1, 2, 3, 4},
			val:   int32(4),
			want:  true,
		},
		{
			name:  "int32 not contains",
			slice: []int32{1, 2, 3, 4},
			val:   int32(5),
			want:  false,
		},
		{
			name:  "int64 contains",
			slice: []int64{1, 2, 3, 4},
			val:   int64(4),
			want:  true,
		},
		{
			name:  "int64 not contains",
			slice: []int64{1, 2, 3, 4},
			val:   int64(5),
			want:  false,
		},
		{
			name:  "uint contains",
			slice: []uint{1, 2, 3, 4},
			val:   uint(4),
			want:  true,
		},
		{
			name:  "uint not contains",
			slice: []uint{1, 2, 3, 4},
			val:   uint(5),
			want:  false,
		},
		{
			name:  "uint8 contains",
			slice: []uint8{1, 2, 3, 4},
			val:   uint8(4),
			want:  true,
		},
		{
			name:  "uint8 not contains",
			slice: []uint8{1, 2, 3, 4},
			val:   uint8(5),
			want:  false,
		},
		{
			name:  "uint16 contains",
			slice: []uint16{1, 2, 3, 4},
			val:   uint16(4),
			want:  true,
		},
		{
			name:  "uint16 not contains",
			slice: []uint16{1, 2, 3, 4},
			val:   uint16(5),
			want:  false,
		},
		{
			name:  "uint32 contains",
			slice: []uint32{1, 2, 3, 4},
			val:   uint32(4),
			want:  true,
		},
		{
			name:  "uint32 not contains",
			slice: []uint32{1, 2, 3, 4},
			val:   uint32(5),
			want:  false,
		},
		{
			name:  "uint64 contains",
			slice: []uint64{1, 2, 3, 4},
			val:   uint64(4),
			want:  true,
		},
		{
			name:  "uint64 not contains",
			slice: []uint64{1, 2, 3, 4},
			val:   uint64(5),
			want:  false,
		},
		{
			name:  "float32 contains",
			slice: []float32{1, 2, 3, 4},
			val:   float32(4),
			want:  true,
		},
		{
			name:  "float32 not contains",
			slice: []float32{1, 2, 3, 4},
			val:   float32(5),
			want:  false,
		},
		{
			name:  "float64 contains",
			slice: []float64{1, 2, 3, 4},
			val:   float64(4),
			want:  true,
		},
		{
			name:  "float64 not contains",
			slice: []float64{1, 2, 3, 4},
			val:   float64(5),
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Contains(tt.slice, tt.val)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestContains_ChecksSliceType(t *testing.T) {
	assert.Panics(t, func() {
		slices.Contains("test", "test")
	})
}

func TestContains_ChecksValType(t *testing.T) {
	assert.Panics(t, func() {
		slices.Contains([]string{"test"}, 1)
	})
}

func BenchmarkContains(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	val := "bat"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.Contains(slice, val)
	}
}

func contains(slice []string, val string) bool {
	for _, vv := range slice {
		if vv == val {
			return true
		}
	}
	return false
}

func BenchmarkContainsNative(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	val := "bat"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contains(slice, val)
	}
}
