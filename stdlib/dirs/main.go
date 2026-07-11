// The `so/os` package provides basic functions
// for working with the file system.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/os"
)

func main() {
	// Create a new sub-directory in the current working directory.
	const dirname = "subdir"
	err := os.Mkdir(dirname, 0755)
	check(err)

	// When creating temporary directories, it's good
	// practice to defer their removal.
	defer removeAll(dirname)

	createFile(dirname + "/file1")
	createFile(dirname + "/file2")
	createFile(dirname + "/file3")

	// ReadDir lists directory contents, returning a
	// slice of os.DirEntry objects.
	entries, err := os.ReadDir(mem.System, dirname)
	check(err)
	defer os.FreeDirEntry(mem.System, entries)

	fmt.Println("Listing", dirname)
	for _, entry := range entries {
		fmt.Println("-", entry.Name)
	}

	// Change the current working directory with Chdir.
	err = os.Chdir("subdir")
	check(err)
	err = os.Chdir("..")
	check(err) // back to the original directory
}

// Helper function to create a new empty file.
func createFile(fpath string) {
	data := []byte("")
	err := os.WriteFile(fpath, data, 0644)
	check(err)
}

func removeAll(dirname string) {
	// os.Remove deletes a file or an empty directory.
	os.Remove(dirname + "/file1")
	os.Remove(dirname + "/file2")
	os.Remove(dirname + "/file3")
	os.Remove(dirname)
}

// check panics on error, to make examples more concise.
// Don't do this in production code.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
