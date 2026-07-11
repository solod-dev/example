// A line filter is a common type of program that reads input on stdin,
// processes it, and then prints some derived result to stdout.
// grep and sed are common line filters.
//
// Here's an example line filter in that writes a
// capitalized version of all input text.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/os"
	"solod.dev/so/strings"
)

func main() {
	// Wrapping the unbuffered os.Stdin with a buffered
	// scanner gives us a convenient Scan method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.
	scanner := bufio.NewScanner(mem.System, os.Stdin)
	defer scanner.Free()

	for scanner.Scan() {
		// Text returns the current token, here the next line,
		// from the input.
		str := strings.ToUpper(mem.System, scanner.Text())
		// Write out the uppercased line.
		fmt.Println(str)
		// Free the occupied memory.
		mem.FreeString(mem.System, str)
	}

	// Check for errors during Scan. End of file is
	// expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
