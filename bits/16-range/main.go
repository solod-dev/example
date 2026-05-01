// `range` iterates over elements in a variety of built-in
// data structures. Let's see how to use `range` with some
// of the data structures we've already learned.
package main

import "solod.dev/so/fmt"

func main() {
	// Iterate over slice and sum the values.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Printf("sum: %d\n", int32(sum))

	// `range` on arrays and slices provides both the index and
	// the value for each item. Earlier, we didn't need the index,
	// so we used the blank identifier `_` to ignore it. But sometimes,
	// we do want to use the indexes.
	for i, num := range nums {
		if num == 3 {
			fmt.Printf("index: %d\n", int32(i))
		}
	}

	// `range` on strings iterates over Unicode code points. The
	// first value is the starting byte index of the `rune` and
	// the second the `rune` itself.
	for i, c := range "go" {
		fmt.Printf("%d %c\n", int32(i), c)
	}
}
