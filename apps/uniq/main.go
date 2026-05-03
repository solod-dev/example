// uniq - filter out repeated lines in a file.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/mem"
	"solod.dev/so/os"
)

var showCount bool

func main() {
	parseFlags()
	args := flag.Args()

	if len(args) == 0 || args[0] == "-" {
		uniq(os.Stdin)
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("uniq: %s: No such file or directory\n", args[0])
			os.Exit(1)
		}
		uniq(&f)
		f.Close()
	}
}

// parseFlags parses command-line flags.
func parseFlags() {
	flag.BoolVar(&showCount, "c", false, "count occurrences of each line")
	flag.Parse()
}

// uniq writes the unique lines from r to standard output, optionally with counts.
func uniq(r io.Reader) {
	scanner := bufio.NewScanner(mem.System, r)
	defer scanner.Free()

	prev := ""
	hasPrev := false
	count := 1

	for scanner.Scan() {
		line := scanner.Text()
		if !hasPrev {
			prev = line
			hasPrev = true
			count = 1
			continue
		}
		if line == prev {
			count++
		} else {
			printLine(prev, count)
			prev = line
			count = 1
		}
	}
	if hasPrev {
		printLine(prev, count)
	}
}

// printLine prints a line with an optional count prefix.
func printLine(line string, count int) {
	if showCount {
		fmt.Printf("%4d %s\n", count, line)
	} else {
		fmt.Println(line)
	}
}
