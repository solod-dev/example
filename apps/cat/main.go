// cat - concatenate and print files.
package main

import (
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	exitCode := 0

	if len(args) == 0 {
		if cat(os.Stdin) != nil {
			exitCode = 1
		}
		os.Exit(exitCode)
	}

	for _, fname := range args {
		if fname == "-" {
			if cat(os.Stdin) != nil {
				writeErr("-", "Error reading standard input")
				exitCode = 1
			}
			continue
		}

		f, err := os.Open(fname)
		if err != nil {
			writeErr(fname, "No such file or directory")
			exitCode = 1
			continue
		}

		if cat(&f) != nil {
			writeErr(fname, "Error reading file")
			exitCode = 1
		}
		f.Close()
	}

	os.Exit(exitCode)
}

// cat writes all data from r to standard output.
func cat(r io.Reader) error {
	_, err := io.Copy(os.Stdout, r)
	return err
}

// writeErr writes an error message to standard error.
func writeErr(name string, msg string) {
	fmt.Printf("cat: %s: %s\n", name, msg)
}
