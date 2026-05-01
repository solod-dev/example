// So has a limited support for multiple return values.
// This feature can be used to return two results (T1, T1)
// or a result and an error (T, error) from a function.
package main

// The `(int, error)` in this function signature shows that
// the function returns an `int` and an `error`.
func vals() (int, error) {
	return 3, nil
}

func main() {
	// Here we use the 2 different return values from the
	// call with multiple assignment.
	v, err := vals()
	println("v =", v, "err =", err)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	v, _ = vals()
	println("v =", v)
	_, err = vals()
	println("err =", err)
}
