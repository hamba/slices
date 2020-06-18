package slices_test

import (
	"fmt"
	"sort"

	"github.com/hamba/slices"
)

func ExampleContains() {
	slice := []string{"foo", "bar"}
	v := "bar"

	fmt.Println(slices.Contains(slice, v))
	// Output: true
}

func ExampleGreaterOf() {
	slice := []string{"foo", "bar", "baz"}
	sort.Slice(slice, slices.GreaterOf(slice))

	fmt.Println(slice)
	// Outputs: [bar baz foo]
}

func ExampleLesserOf() {
	slice := []string{"foo", "bar", "baz"}
	sort.Slice(slice, slices.LesserOf(slice))

	fmt.Println(slice)
	// Outputs: [foo baz bar]
}
