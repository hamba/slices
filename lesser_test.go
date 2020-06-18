package slices_test

import (
	"sort"
	"testing"

	"github.com/hamba/slices"
	"github.com/stretchr/testify/assert"
)

func TestLesserOf(t *testing.T) {
	tests := []struct {
		name  string
		slice interface{}
		want  interface{}
	}{
		{
			name:  "string",
			slice: []string{"foo", "bar", "baz", "bat"},
			want:  []string{"bar", "bat", "baz", "foo"},
		},
		{
			name:  "int",
			slice: []int{4, 2, 1, 3},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "int8",
			slice: []int8{4, 2, 1, 3},
			want:  []int8{1, 2, 3, 4},
		},
		{
			name:  "int16",
			slice: []int16{4, 2, 1, 3},
			want:  []int16{1, 2, 3, 4},
		},
		{
			name:  "int32",
			slice: []int32{4, 2, 1, 3},
			want:  []int32{1, 2, 3, 4},
		},
		{
			name:  "int64",
			slice: []int64{4, 2, 1, 3},
			want:  []int64{1, 2, 3, 4},
		},
		{
			name:  "uint",
			slice: []uint{4, 2, 1, 3},
			want:  []uint{1, 2, 3, 4},
		},
		{
			name:  "uint8",
			slice: []uint8{4, 2, 1, 3},
			want:  []uint8{1, 2, 3, 4},
		},
		{
			name:  "uint16",
			slice: []uint16{4, 2, 1, 3},
			want:  []uint16{1, 2, 3, 4},
		},
		{
			name:  "uint32",
			slice: []uint32{4, 2, 1, 3},
			want:  []uint32{1, 2, 3, 4},
		},
		{
			name:  "uint64",
			slice: []uint64{4, 2, 1, 3},
			want:  []uint64{1, 2, 3, 4},
		},
		{
			name:  "float32",
			slice: []float32{4, 2, 1, 3},
			want:  []float32{1, 2, 3, 4},
		},
		{
			name:  "float64",
			slice: []float64{4, 2, 1, 3},
			want:  []float64{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.slice, slices.LesserOf(tt.slice))

			assert.Equal(t, tt.want, tt.slice)
		})
	}
}

func TestLesserOf_ChecksSliceType(t *testing.T) {
	assert.Panics(t, func() {
		slices.LesserOf("test")
	})
}

func BenchmarkLesserOf(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(slice, slices.LesserOf(slice))
	}
}

func BenchmarkLesserOfNative(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	lesserOf := func(i, j int) bool {
		return slice[i] < slice[j]
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(slice, lesserOf)
	}
}
