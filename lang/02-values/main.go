// So supports different value types like strings,
// integers, floats, booleans, and more.
// Here are a few simple examples.
package main

func main() {
	// Integers and floats.
	println("1+1 =", 1+1)
	println("5.0/2.0 =", 5.0/2.0)

	// Booleans and boolean operators.
	println(true && false)
	println(true || false)
	println(!true)

	// Strings.
	println("so is awesome")

	// You can concatenate strings with the + operator.
	// Don't do that for large strings,
	// because the result is stack-allocated.
	println("so " + "is " + "awesome")
}
