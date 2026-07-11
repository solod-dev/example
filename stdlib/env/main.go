// The `so/os` package provides basic functions
// for accessing environment variables.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/os"
)

func main() {
	// To set a key/value pair, use os.Setenv.
	os.Setenv("TOKEN", "42")

	// To get a value for a key, use os.Getenv.
	fmt.Println("TOKEN:", os.Getenv("TOKEN"))
	// os.Getenv returns an empty string if the key
	// is not present in the environment.
	fmt.Println("SECRET:", os.Getenv("SECRET"))

	// To unset a key, use os.Unsetenv.
	os.Unsetenv("TOKEN")
	fmt.Println("TOKEN:", os.Getenv("TOKEN"))
}
