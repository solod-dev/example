// Use os.Exit to immediately exit with a given status.
//
// Note that unlike C, So does not use an integer return value
// from main to indicate exit status. If you'd like to exit with
// a non-zero status, use os.Exit.
package main

import "solod.dev/so/os"

func main() {
	// defers do not run when using os.Exit,
	// so this println will never be called.
	defer println("!")

	// Exit with status 3.
	os.Exit(3)
}
