// head - display first lines of a file.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/mem"
	"solod.dev/so/os"
)

var nLines int
var nBytes int

func main() {
	parseFlags()
	args := flag.Args()

	if len(args) == 0 {
		head(os.Stdin)
		os.Exit(0)
	}

	exitCode := 0
	for i, fname := range args {
		if len(args) > 1 {
			if i > 0 {
				fmt.Println("")
			}
			fmt.Println("==>", fname, "<==")
		}
		f, err := os.Open(fname)
		if err != nil {
			fmt.Printf("head: %s: No such file or directory\n", fname)
			exitCode = 1
			continue
		}
		head(&f)
		f.Close()
	}
	os.Exit(exitCode)
}

// parseFlags parses command-line flags.
func parseFlags() {
	flag.IntVar(&nLines, "n", 10, "print count lines of each of the specified files")
	flag.IntVar(&nBytes, "c", 0, "print bytes of each of the specified files")
	flag.Parse()
}

// head writes the first nLines lines or nBytes bytes
// from r to standard output.
func head(r io.Reader) {
	if nBytes > 0 {
		headBytes(r, nBytes)
	} else {
		headLines(r, nLines)
	}
}

// headLines writes the first count lines from r to standard output.
func headLines(r io.Reader, count int) {
	scanner := bufio.NewScanner(mem.System, r)
	defer scanner.Free()
	printed := 0
	for printed < count && scanner.Scan() {
		fmt.Println(scanner.Text())
		printed++
	}
}

// headBytes writes the first count bytes from r to standard output.
func headBytes(r io.Reader, count int) {
	io.CopyN(os.Stdout, r, int64(count))
}
