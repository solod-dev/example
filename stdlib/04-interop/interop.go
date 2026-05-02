package main

import "solod.dev/so/c"

// The `include` directive tell So to #include
// the given headers into the generated C code.

//so:include <stdlib.h>
//so:include <time.h>

// The `embed` directive embeds the contents of a .h or .c file
// into the generated C code. This is useful for defining helper
// functions or types, like we do here with `tm` (see below).

//so:embed interop.h
var interop_h string

// The `extern` directive declares a C symbol so that you can refer
// to it from So code. Writing a function without a body is the same
// as declaring it extern.

//so:extern
func getenv(name string) *c.Char {
	// String arguments auto-decay to C's `char*`.
	// The bodies of extern functions are ignored.
	return nil
}

//so:extern
func atoi(s string) int32 { return 0 }

//so:extern
func abs(n int32) int32 {
	// Scalar types like int32 map directly to C's int.
	return 0
}

// time_t is a C type for representing time values.
// We declare it as extern so that So can access it,
// but doesn't codegen it, since the type is already
// defined in the C header.
//
//so:extern
type time_t int64

// time returns the current calendar time
// as seconds since the Unix epoch.
func time(timer *time_t) time_t

// tm is C's broken-down time structure (defined in interop.h).
//
//so:extern
type tm struct{}

// localtime converts a time value to local time.
func localtime(timer *time_t) *tm

// strftime formats a time value into a string.
func strftime(buf *c.Char, maxsize uintptr, format string, timeptr *tm) int

// difftime computes the difference between two times.
func difftime(end, start time_t) float64
