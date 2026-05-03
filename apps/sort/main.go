// sort - sort text files by lines.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/c"
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/mem"
	"solod.dev/so/os"
	"solod.dev/so/slices"
	"solod.dev/so/strings"
)

var reverse bool
var keyField int
var separator string

func main() {
	parseFlags()
	args := flag.Args()

	lines := slices.MakeCap[string](mem.System, 0, 64)
	defer freeLines(lines)
	exitCode := 0

	if len(args) == 0 {
		lines = readLines(os.Stdin, lines)
	} else {
		for _, fname := range args {
			if fname == "-" {
				lines = readLines(os.Stdin, lines)
				continue
			}
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "sort: %s: No such file or directory\n", fname)
				exitCode = 2
				continue
			}
			lines = readLines(&f, lines)
			f.Close()
		}
	}

	slices.SortFunc(lines, compareLines)

	for _, line := range lines {
		fmt.Println(line)
	}

	os.Exit(exitCode)
}

// parseFlags parses command-line flags.
func parseFlags() {
	flag.BoolVar(&reverse, "r", false, "sort in reverse order")
	flag.IntVar(&keyField, "k", 0, "1-based field index to sort by")
	flag.StringVar(&separator, "t", "", "field separator character")
	flag.Parse()
}

// readLines reads lines from r and appends them to the lines slice.
func readLines(r io.Reader, lines []string) []string {
	scanner := bufio.NewScanner(mem.System, r)
	defer scanner.Free()
	for scanner.Scan() {
		line := strings.Clone(mem.System, scanner.Text())
		lines = slices.Append(mem.System, lines, line)
	}
	return lines
}

// compareLines compares two lines for sorting, using
// the specified key field and separator if applicable.
func compareLines(a, b any) int {
	sa := *c.PtrAs[string](a)
	sb := *c.PtrAs[string](b)

	if keyField > 0 {
		sa = extractField(sa)
		sb = extractField(sb)
	}

	result := strings.Compare(sa, sb)
	if reverse {
		return -result
	}
	return result
}

// freeLines frees the memory used by the lines slice and its contents.
func freeLines(lines []string) {
	for _, line := range lines {
		mem.FreeString(mem.System, line)
	}
	slices.Free(mem.System, lines)
}

// extractField extracts the specified key field from the line.
func extractField(line string) string {
	if len(separator) > 0 {
		return fieldBySep(line, keyField, separator[0])
	}
	return fieldBySpace(line, keyField)
}

// fieldBySpace extracts the specified field from the line,
// using whitespace as the separator.
func fieldBySpace(line string, field int) string {
	n := len(line)
	current := 0
	i := 0

	for i < n {
		// Skip whitespace.
		for i < n && isSpace(line[i]) {
			i++
		}
		if i >= n {
			break
		}
		current++
		start := i
		// Skip non-whitespace.
		for i < n && !isSpace(line[i]) {
			i++
		}
		if current == field {
			return line[start:i]
		}
	}
	return ""
}

// fieldBySep extracts the specified field from the line,
// using the specified separator character.
func fieldBySep(line string, field int, sep byte) string {
	n := len(line)
	current := 1
	start := 0

	for i := 0; i <= n; i++ {
		if i == n || line[i] == sep {
			if current == field {
				return line[start:i]
			}
			current++
			start = i + 1
		}
	}
	return ""
}

// isSpace reports whether ch is a whitespace character.
func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
