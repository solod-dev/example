// Measuring performance is an important part of writing So programs.
// The `so/testing` package provides the tools we need to write
// benchmarks, and the `so bench` command runs them.

// By convention, benchmark code lives in the `bench` subdirectory of the
// package we are testing - that's where the `so bench` command looks for it.
//
// Do not name the file with benchmarks `main.go` - that file will be
// auto-generated when you run `so bench` and will contain the `main`
// function that runs the benchmarks.
//
// Do not use the `_test.go` suffix for benchmark files.
package main

import (
	"app"

	"solod.dev/so/testing"
)

// A sink variable prevents the compiler from
// optimizing away the benchmarked code.
//
//so:volatile
var sink int

func BenchmarkIntMin(b *testing.B) {
	// Any code that's required for the benchmark to run but should
	// not be measured goes before this loop.
	for b.Loop() {
		// The benchmark runner will automatically execute this loop
		// body many times to determine a reasonable estimate of the
		// run-time of a single iteration.
		sink = app.IntMin(1, 2)
		// You can also use testing.Keep instead of the global sink
		// variable to prevent the compiler from optimizing away the
		// benchmarked code.
	}
}
