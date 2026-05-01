// So supports constants of character, string,
// boolean, and numeric values.
package main

import "solod.dev/so/math"

// `const` declares a constant value.
const s string = "constant"

func main() {
	println("s =", s)

	// You can also declare a constant inside a function.
	const n = 500000000

	const ch = 'a'
	println("ch =", ch) // 97 - the ASCII code for 'a'

	// Constant expressions perform arithmetic.
	// This only works for constants declared inside functions.
	const d = 3e20 / n
	println("d =", d)

	// Integer constants are `int` by default, but
	// you can cast them to another compatible type.
	println("int64(d) =", int64(d))

	// The compiler will automatically convert to
	// a compatible type if possible. For example,
	// `math.Sqrt` expects a `float64` here.
	println("math.Sqrt(n) =", math.Sqrt(n))
}
