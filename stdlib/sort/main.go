// The `so/slices` package implements sorting for builtins and
// user-defined types. We'll look at sorting for builtins first.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/slices"
)

func main() {
	// Sorting functions are generic, and work for any
	// ordered built-in type. For a list of ordered
	// types, see `cmp.Ordered`.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Printf("Strings: [%s %s %s]\n", strs[0], strs[1], strs[2])

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Printf("Ints: [%d %d %d]\n", ints[0], ints[1], ints[2])

	// We can also use the `slices` package to check
	// if a slice is already in sorted order.
	if slices.IsSorted(ints) {
		fmt.Println("Sorted!")
	} else {
		fmt.Println("Not sorted :(")
	}
}
