// The `so/path` package provides functions
// to parse and construct Unix-like paths.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/path"
	"solod.dev/so/strings"
)

func main() {
	// Join takes any number of arguments and constructs
	// a hierarchical path from them.
	p := path.Join(mem.System, "dir1", "dir2", "filename")
	fmt.Println(p)
	mem.FreeString(mem.System, p)

	// Join also normalizes paths by removing extra separators and directory changes.
	p = path.Join(mem.System, "dir1//", "filename")
	fmt.Println(p)
	mem.FreeString(mem.System, p)

	p = path.Join(mem.System, "dir1/../dir1", "filename")
	fmt.Println(p)
	mem.FreeString(mem.System, p)

	const fpath = "path/to/file.txt"
	println(fpath)

	// Dir and Base split a path into the directory and the file.
	// Alternatively, Split returns both at once.
	dir := path.Dir(mem.System, fpath)
	fmt.Println("- dir =", dir)
	mem.FreeString(mem.System, dir)
	base := path.Base(fpath)
	fmt.Println("- base =", base)

	// Use Ext to get the file extension.
	ext := path.Ext(fpath)
	fmt.Println("- ext =", ext)

	// To find the file's name with the extension removed,
	// use strings.TrimSuffix.
	stem := strings.TrimSuffix(base, ext)
	fmt.Println("- stem =", stem)
}
