// So is statically typed. Variables are declared explicitly,
// and their types are known to the compiler.
package main

func main() {
	// `var` declares a variable, and `=` assigns it a specific value:
	var b bool = true
	println("b =", b)

	var s string = "hello"
	println("s =", s)

	var i int = 42
	println("i =", i)

	var f float64 = 12.34
	println("f =", f)

	// You can declare multiple variables at once.
	var one, two int = 1, 2
	println(one, two)

	// So will infer the type of initialized variables.
	var sunny = true // inferred as `bool`
	println("sunny =", sunny)

	// Если не инициализировать переменную при объявлении, она получит нулевое значение (zero value). У каждого типа оно свое: int — 0, string — "", bool — false.

	// If you don't initialize a variable when you declare it,
	// it gets a zero value. Each type has its own: int is 0,
	// string is "", and bool is false.
	var num int
	var str string
	var ok bool
	println("num =", num, "str =", str, "ok =", ok)

	// The := operator declares and initializes a variable.
	// You don't need to specify `var` or the type.
	food := "apple" // var food string = "apple"
	println("food =", food)
}
