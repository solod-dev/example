// In So, an array is a numbered sequence of elements
// of a specific, fixed length.
package main

import "solod.dev/so/fmt"

func main() {
	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var a [3]int32
	printArray("empty:", a)

	// You can access an array element using square brackets.
	a[2] = 100
	printArray("set[2]:", a)
	println("get[2]:", a[2])

	// The `len` builtin returns the length of an array.
	println("len:", len(a))

	// You can declare and initialize an array in one line.
	b := [3]int32{1, 2, 3}
	printArray("dcl:", b)

	// You can also have the compiler count the number of
	// elements for you with `...`
	b = [...]int32{1, 2, 3}
	printArray("dcl:", b)

	// If you specify the index with `:`, the elements in
	// between will be zeroed.
	b = [...]int32{2: 300}
	printArray("idx:", b)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int32
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = int32(i + j)
		}
	}
	fmt.Printf("2d: [%d...%d]\n", twoD[0][0], twoD[1][2])

	// You can create and initialize multi-dimensional
	// arrays at once too.
	twoD = [2][3]int32{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Printf("2d: [%d...%d]\n", twoD[0][0], twoD[1][2])
}

func printArray(msg string, arr [3]int32) {
	fmt.Printf("%s [%d %d %d]\n", msg, arr[0], arr[1], arr[2])
}
