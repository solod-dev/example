// Reading and writing files are basic tasks needed for many So
// programs. First we'll look at some examples of reading files.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/mem"
	"solod.dev/so/os"
)

func main() {
	const fpath = "/etc/hosts"
	readFull(fpath)
	println("---")
	readParts(fpath)
}

func readFull(fpath string) {
	// One of the simplest ways to read a file is to load
	// its entire contents into memory with os.ReadFile.
	data, err := os.ReadFile(mem.System, fpath)
	check(err)
	println(fpath)
	os.Stdout.Write(data)
	mem.FreeSlice(mem.System, data)
}

func readParts(fpath string) {
	// You'll often want more control over how and what parts of
	// a file are read. For these tasks, open the file with os.Open
	// and use the returned os.File as an io.Reader and io.Seeker.
	f, err := os.Open(fpath)
	check(err)
	defer f.Close() // Ensure the file is closed when we're done.

	println(fpath)
	{
		// Read some bytes from the beginning of the file.
		// Allow up to 5 to be read but also note how many
		// actually were read.
		data := make([]byte, 5)
		n1, err := f.Read(data)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(data[:n1]))
	}

	{
		// You can also Seek to a known location in the file
		// and Read from there.
		offset, err := f.Seek(6, io.SeekStart)
		check(err)
		data := make([]byte, 2)
		n2, err := f.Read(data)
		check(err)
		fmt.Printf("%d bytes @ %d: ", n2, offset)
		fmt.Printf("%s\n", string(data[:n2]))
	}

	{
		// Other methods of seeking are relative to the
		// current cursor position,
		_, err = f.Seek(2, io.SeekCurrent)
		check(err)

		// and relative to the end of the file.
		_, err = f.Seek(-4, io.SeekEnd)
		check(err)

		// There is no built-in rewind, but
		// Seek(0, io.SeekStart) accomplishes this.
		_, err = f.Seek(0, io.SeekStart)
		check(err)
	}

	{
		// The `bufio` package provides a buffered reader,
		// which makes small reads more efficient.
		r := bufio.NewReader(mem.System, &f)
		defer r.Free()
		data, err := r.Peek(5)
		check(err)
		fmt.Printf("5 bytes: %s\n", string(data))
	}
}

// check panics on error, to make examples more concise.
// Don't do this in production code.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
