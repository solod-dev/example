// Writing files in So follows similar patterns
// to what we saw earlier for reading.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/os"
	"solod.dev/so/path"
)

func main() {
	writeFull()
	writeParts()
}

func writeFull() {
	// To start, here's how to dump a string (or just
	// bytes) into a file.
	data := []byte("hello\ngo\n")
	fpath := path.Join(mem.System, os.TempDir(), "dat1")
	defer mem.FreeString(mem.System, fpath)
	err := os.WriteFile(fpath, data, 0644)
	check(err)
}

func writeParts() {
	// For more granular writes, open a file for writing.
	fpath := path.Join(mem.System, os.TempDir(), "dat2")
	defer mem.FreeString(mem.System, fpath)
	f, err := os.Create(fpath)
	check(err)
	defer f.Close()

	// Write some bytes to the file.
	data := []byte{115, 111, 109, 101, 10}
	n, err := f.Write(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)

	// Write a string to the file.
	n, err = f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n)

	// bufio provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(mem.System, &f)
	defer w.Free()
	n, err = w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n)

	// Use Flush to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()
}

// check panics on error, to make examples more concise.
// Don't do this in production code.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
