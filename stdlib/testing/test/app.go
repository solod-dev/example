// Unit testing is an important part of writing So programs.
// The `so/testing` package provides the tools we need to write
// unit tests, and the `so test` command runs them.

// By convention, testing code lives in the `test` subdirectory of the
// package we are testing - that's where the `so test` command looks for it.
//
// Do not name the file with tests `main.go` - that file will be
// auto-generated when you run `so test` and will contain the `main`
// function that runs the tests.
//
// Do not use the `_test.go` suffix for test files.
package app_test

import (
	"app"

	"solod.dev/so/fmt"
	"solod.dev/so/testing"
)

// A test is created by writing a function with a name
// beginning with `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := app.IntMin(2, -2)
	if ans != -2 {
		// t.Error reports a test failure but continue executing the test.
		// t.Fatal reports a test failure and must be followed by a return
		// statement to stop executing the test.
		buf := make([]byte, 64)
		msg := fmt.Sprintf(buf, "IntMin(2, -2) = %d, want -2", ans)
		t.Error(msg)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	// Writing tests can be repetitive, so it's idiomatic to use
	// a table-driven style, where test inputs and expected outputs
	// are listed in a table and a single loop walks over them
	// and performs the test logic.
	type testCase struct {
		a, b int
		want int
	}
	var tests = []testCase{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	buf := make([]byte, 64)
	for _, tt := range tests {
		ans := app.IntMin(tt.a, tt.b)
		if ans != tt.want {
			msg := fmt.Sprintf(buf, "IntMin(%d, %d) = %d, want %d", tt.a, tt.b, ans, tt.want)
			t.Error(msg)
		}
	}
}
