package slices_test

import (
	"sort"
	"testing"

	"github.com/hamba/slices"
	"github.com/stretchr/testify/assert"
)

func TestGreaterOf(t *testing.T) {
	tests := []struct {
		name  string
		slice interface{}
		want  interface{}
	}{
		{
			name:  "string",
			slice: []string{"foo", "bar", "baz", "bat"},
			want:  []string{"foo", "baz", "bat", "bar"},
		},
		{
			name:  "int",
			slice: []int{4, 2, 1, 3},
			want:  []int{4, 3, 2, 1},
		},
		{
			name:  "int8",
			slice: []int8{4, 2, 1, 3},
			want:  []int8{4, 3, 2, 1},
		},
		{
			name:  "int16",
			slice: []int16{4, 2, 1, 3},
			want:  []int16{4, 3, 2, 1},
		},
		{
			name:  "int32",
			slice: []int32{4, 2, 1, 3},
			want:  []int32{4, 3, 2, 1},
		},
		{
			name:  "int64",
			slice: []int64{4, 2, 1, 3},
			want:  []int64{4, 3, 2, 1},
		},
		{
			name:  "uint",
			slice: []uint{4, 2, 1, 3},
			want:  []uint{4, 3, 2, 1},
		},
		{
			name:  "uint8",
			slice: []uint8{4, 2, 1, 3},
			want:  []uint8{4, 3, 2, 1},
		},
		{
			name:  "uint16",
			slice: []uint16{4, 2, 1, 3},
			want:  []uint16{4, 3, 2, 1},
		},
		{
			name:  "uint32",
			slice: []uint32{4, 2, 1, 3},
			want:  []uint32{4, 3, 2, 1},
		},
		{
			name:  "uint64",
			slice: []uint64{4, 2, 1, 3},
			want:  []uint64{4, 3, 2, 1},
		},
		{
			name:  "float32",
			slice: []float32{4, 2, 1, 3},
			want:  []float32{4, 3, 2, 1},
		},
		{
			name:  "float64",
			slice: []float64{4, 2, 1, 3},
			want:  []float64{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.slice, slices.GreaterOf(tt.slice))

			assert.Equal(t, tt.want, tt.slice)
		})
	}
}

func TestGreaterOf_ChecksSliceType(t *testing.T) {
	assert.Panics(t, func() {
		slices.GreaterOf("test")
	})
}

func BenchmarkGreaterOf(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(slice, slices.GreaterOf(slice))
	}
}

func BenchmarkGreaterOfNative(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	greaterOf := func(i, j int) bool {
		return slice[i] > slice[j]
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(slice, greaterOf)
	}
}
