// Functions are central in So. We'll learn about
// functions with a few different examples.
package main

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func sum2(a int, b int) int {
	return a + b
}

// When you have several parameters in a row with the same type,
// you can leave out the type name for all but the last one.
// Just include the type with the final parameter.
func sum3(a, b, c int) int {
	return a + b + c
}

func main() {
	res := sum2(1, 2)
	println("1+2 =", res)

	res = sum3(1, 2, 3)
	println("1+2+3 =", res)
}
