// So integrates with C using so:include, so:extern,
// and so:embed directives. You can also write raw C code
// using c.Val and c.Raw compiler intrinsics.
package main

import "solod.dev/so/c"

// See interop.go for C bindings.

func main() {
	// Read an environment variable. getenv returns char*
	// so we use c.String to convert it to a So string.
	home := c.String(getenv("HOME"))
	println("HOME =", home)

	// Convert a string to a number.
	n := atoi("12345")
	println("atoi(12345) =", n)

	// Compute an absolute value.
	println("abs(-42) =", abs(-42))

	// Get the current time.
	now := time(nil)
	println("seconds since epoch =", now)

	// Format the current time as a string.
	// localtime returns a *tm (C's broken-down time).
	// strftime writes the formatted time into a buffer.
	local := localtime(&now)
	var buf [64]byte
	cbuf := (*c.Char)(&buf[0])
	strftime(cbuf, 64, "%Y-%m-%d %H:%M:%S", local)
	println("current time =", c.String(cbuf))

	// Compute the difference between two times.
	start := now
	end := start + 3600
	diff := difftime(end, start)
	println("seconds in an hour =", diff)

	// c.Val emits a typed C expression.
	// Use it to access C constants, macros, or call C
	// functions inline, without declaring them as extern.
	println("EXIT_SUCCESS =", c.Val[int32]("EXIT_SUCCESS"))

	// c.Raw emits a raw block of C code as a statement.
	var squared int32
	c.Raw(`
	int a = 7;
	squared = a * a;
	`)
	println("7*7 =", squared)
}
