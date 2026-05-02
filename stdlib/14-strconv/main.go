// The `so/strconv` package provides functions for converting between
// strings and other types, similar to Go's `strconv` package.
package main

import "solod.dev/so/strconv"

func main() {
	// With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	println(f)

	// For ParseInt, 0 means infer the base from the string,
	// and 64 requires that the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	println(i)

	// ParseInt will recognizes hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	println(d)

	// ParseUint is also available.
	u, _ := strconv.ParseUint("789", 0, 64)
	println(u)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135") // same as ParseInt("135", 10, 0)
	println(k)

	// Parse functions return an error on bad input.
	_, err := strconv.Atoi("wat")
	println(err)
}
