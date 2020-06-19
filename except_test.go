package slices_test

import (
	"testing"

	"github.com/hamba/slices"
	"github.com/stretchr/testify/assert"
)

func TestExcept(t *testing.T) {
	tests := []struct {
		name string
		a    interface{}
		b    interface{}
		want interface{}
	}{
		{
			name: "string contains",
			a:    []string{"foo", "bar", "baz", "bat"},
			b:    []string{"bar", "baz", "test"},
			want: []string{"foo", "bat"},
		},
		{
			name: "int contains",
			a:    []int{1, 2, 3, 4},
			b:    []int{2, 3, 5},
			want: []int{1, 4},
		},
		{
			name: "int8 contains",
			a:    []int8{1, 2, 3, 4},
			b:    []int8{2, 3, 5},
			want: []int8{1, 4},
		},
		{
			name: "int16 contains",
			a:    []int16{1, 2, 3, 4},
			b:    []int16{2, 3, 5},
			want: []int16{1, 4},
		},
		{
			name: "int32 contains",
			a:    []int32{1, 2, 3, 4},
			b:    []int32{2, 3, 5},
			want: []int32{1, 4},
		},
		{
			name: "int64 contains",
			a:    []int64{1, 2, 3, 4},
			b:    []int64{2, 3, 5},
			want: []int64{1, 4},
		},
		{
			name: "uint contains",
			a:    []uint{1, 2, 3, 4},
			b:    []uint{2, 3, 5},
			want: []uint{1, 4},
		},
		{
			name: "uint8 contains",
			a:    []uint8{1, 2, 3, 4},
			b:    []uint8{2, 3, 5},
			want: []uint8{1, 4},
		},
		{
			name: "uint16 contains",
			a:    []uint16{1, 2, 3, 4},
			b:    []uint16{2, 3, 5},
			want: []uint16{1, 4},
		},
		{
			name: "uint32 contains",
			a:    []uint32{1, 2, 3, 4},
			b:    []uint32{2, 3, 5},
			want: []uint32{1, 4},
		},
		{
			name: "uint64 contains",
			a:    []uint64{1, 2, 3, 4},
			b:    []uint64{2, 3, 5},
			want: []uint64{1, 4},
		},
		{
			name: "float32 contains",
			a:    []float32{1, 2, 3, 4},
			b:    []float32{2, 3, 5},
			want: []float32{1, 4},
		},
		{
			name: "float64 contains",
			a:    []float64{1, 2, 3, 4},
			b:    []float64{2, 3, 5},
			want: []float64{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Except(tt.a, tt.b)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExcept_ChecksSliceType(t *testing.T) {
	assert.Panics(t, func() {
		slices.Except("test", "test")
	})
}

func TestExcept_ChecksValType(t *testing.T) {
	assert.Panics(t, func() {
		slices.Except([]string{"test"}, 1)
	})
}

func BenchmarkExcept(b *testing.B) {
	a := []string{"foo", "bar", "baz", "bat"}
	c := []string{"bar", "baz", "test"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.Except(a, c)
	}
}

func except(slice, other []string) interface{} {
	s := make([]string, len(slice))
	copy(s, slice)
	for i := 0; i < len(s); i++ {
		for _, v := range other {
			if v == s[i] {
				s = append(s[:i], s[i+1:]...)
				i--
				break
			}
		}
	}
	return s
}

func BenchmarkExceptNative(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "bat"}
	other := []string{"bar", "baz", "test"}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		except(slice, other)
	}
}
