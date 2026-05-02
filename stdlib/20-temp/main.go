// The `so/os` package provides basic functions
// for working with temporary files and directories.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/os"
)

func main() {
	var buf [os.MaxPathLen]byte

	// The easiest way to create a temporary file is by using os.CreateTemp.
	// It creates the file and opens it for reading and writing. We pass ""
	// as the first argument, so os.CreateTemp will create the file in the
	// default location for the OS.
	f, err := os.CreateTemp(buf[:], "", "tfile_")
	check(err)

	// The temp file name starts with the prefix given as the second argument
	// to os.CreateTemp and the rest is chosen automatically.
	fmt.Println("path =", f.Name())

	// Delete the file after we're done.
	os.Remove(f.Name())

	// To create a temporary directory, use os.MkdirTemp. It takes the same
	// arguments as os.CreateTemp, but returns the path to the new directory
	// instead of an open file.
	d, err := os.MkdirTemp(buf[:], "", "tdir_")
	check(err)
	fmt.Println("dir =", d)

	// Delete the directory after we're done.
	os.Remove(d)
}

// check panics on error, to make examples more concise.
// Don't do this in production code.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
