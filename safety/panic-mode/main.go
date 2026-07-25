// A panic prints a message and terminates the program. The `-panic`
// build flag controls what happens after the message is printed:
//
//	so run -panic=exit  .  # call exit(1) (the default)
//	so run -panic=abort .  # call abort(), raising SIGABRT for a debugger
//	so run -panic=trace .  # print a stack trace, then exit(1)
//
// Run this program with `-panic=trace` to see the backtrace. The frames
// are C symbols (`main_Parse`, `main_Field`, ...), which map directly onto
// the So functions below.
package main

// Field returns the value at position i, or panics if i is out of range.
func Field(row []int, i int) int {
	if i >= len(row) {
		panic("field: index out of range")
	}
	return row[i]
}

// Parse reads the expected columns out of a row.
func Parse(row []int) int {
	id := Field(row, 0)
	value := Field(row, 5) // no such column, will panic
	return id + value
}

func main() {
	row := []int{42}
	println(Parse(row))
}
