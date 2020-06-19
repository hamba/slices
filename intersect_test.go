package slices_test

import (
	"testing"

	"github.com/hamba/slices"
	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name  string
		slice interface{}
		other interface{}
		want  interface{}
	}{
		{
			name:  "string contains",
			slice: []string{"foo", "bar", "baz", "bat"},
			other: []string{"bar", "baz", "test"},
			want:  []string{"bar", "baz"},
		},
		{
			name:  "int contains",
			slice: []int{1, 2, 3, 4},
			other: []int{2, 3, 5},
			want:  []int{2, 3},
		},
		{
			name:  "int8 contains",
			slice: []int8{1, 2, 3, 4},
			other: []int8{2, 3, 5},
			want:  []int8{2, 3},
		},
		{
			name:  "int16 contains",
			slice: []int16{1, 2, 3, 4},
			other: []int16{2, 3, 5},
			want:  []int16{2, 3},
		},
		{
			name:  "int32 contains",
			slice: []int32{1, 2, 3, 4},
			other: []int32{2, 3, 5},
			want:  []int32{2, 3},
		},
		{
			name:  "int64 contains",
			slice: []int64{1, 2, 3, 4},
			other: []int64{2, 3, 5},
			want:  []int64{2, 3},
		},
		{
			name:  "uint contains",
			slice: []uint{1, 2, 3, 4},
			other: []uint{2, 3, 5},
			want:  []uint{2, 3},
		},
		{
			name:  "uint8 contains",
			slice: []uint8{1, 2, 3, 4},
			other: []uint8{2, 3, 5},
			want:  []uint8{2, 3},
		},
		{
			name:  "uint16 contains",
			slice: []uint16{1, 2, 3, 4},
			other: []uint16{2, 3, 5},
			want:  []uint16{2, 3},
		},
		{
			name:  "uint32 contains",
			slice: []uint32{1, 2, 3, 4},
			other: []uint32{2, 3, 5},
			want:  []uint32{2, 3},
		},
		{
			name:  "uint64 contains",
			slice: []uint64{1, 2, 3, 4},
			other: []uint64{2, 3, 5},
			want:  []uint64{2, 3},
		},
		{
			name:  "float32 contains",
			slice: []float32{1, 2, 3, 4},
			other: []float32{2, 3, 5},
			want:  []float32{2, 3},
		},
		{
			name:  "float64 contains",
			slice: []float64{1, 2, 3, 4},
			other: []float64{2, 3, 5},
			want:  []float64{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Intersect(tt.slice, tt.other)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntersect_ChecksSliceType(t *testing.T) {
	assert.Panics(t, func() {
		slices.Intersect("test", "test")
	})
}

func TestIntersect_ChecksValType(t *testing.T) {
	assert.Panics(t, func() {
		slices.Intersect([]string{"test"}, 1)
	})
}

func BenchmarkIntersect(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	other := []string{"bar", "baz", "test"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.Intersect(slice, other)
	}
}

func intersect(slice, other []string) interface{} {
	s := make([]string, len(slice))
	copy(s, slice)
	for i := 0; i < len(s); i++ {
		found := false
		for _, v := range other {
			if v == s[i] {
				found = true
				break
			}
		}
		if !found {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func BenchmarkIntersectNative(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	other := []string{"bar", "baz", "test"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersect(slice, other)
	}
}
