// wc - word, line, character, and byte count.
package main

import (
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/os"
)

type counts struct {
	lines int
	words int
	bytes int
	chars int
}

var showLines bool
var showWords bool
var showBytes bool
var showChars bool

func main() {
	parseFlags()

	if !showLines && !showWords && !showBytes && !showChars {
		showLines = true
		showWords = true
		showBytes = true
	}

	args := flag.Args()

	if len(args) == 0 {
		c := wc(os.Stdin)
		printCounts(c, "")
		os.Exit(0)
	}

	exitCode := 0
	var total counts
	for _, fname := range args {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %s: No such file or directory\n", fname)
			exitCode = 1
			continue
		}
		c := wc(&f)
		printCounts(c, fname)
		total.lines += c.lines
		total.words += c.words
		total.bytes += c.bytes
		total.chars += c.chars
		f.Close()
	}

	if len(args) > 1 {
		printCounts(total, "total")
	}

	os.Exit(exitCode)
}

// parseFlags parses command-line flags.
func parseFlags() {
	flag.BoolVar(&showLines, "l", false, "count lines")
	flag.BoolVar(&showWords, "w", false, "count words")
	flag.BoolVar(&showBytes, "c", false, "count bytes")
	flag.BoolVar(&showChars, "m", false, "count characters")
	flag.Parse()
}

// wc counts the lines, words, bytes, and characters in r.
func wc(r io.Reader) counts {
	var c counts
	buf := make([]byte, 4096)
	inWord := false

	for {
		n, err := r.Read(buf)
		c.bytes += n
		for i := range n {
			b := buf[i]
			if b == '\n' {
				c.lines++
			}
			if b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '\f' || b == '\v' {
				inWord = false
			} else if !inWord {
				inWord = true
				c.words++
			}
			if showChars && (b&0xC0) != 0x80 { // start of UTF-8 character
				c.chars++
			}
		}
		if err != nil {
			break
		}
	}
	return c
}

// printCounts prints the counts in c, optionally with a filename.
func printCounts(c counts, fname string) {
	if showLines {
		fmt.Printf("%8d", c.lines)
	}
	if showWords {
		fmt.Printf("%8d", c.words)
	}
	if showBytes {
		fmt.Printf("%8d", c.bytes)
	}
	if showChars {
		fmt.Printf("%8d", c.chars)
	}
	if fname != "" {
		fmt.Printf(" %s", fname)
	}
	fmt.Printf("\n")
}
