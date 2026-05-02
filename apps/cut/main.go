// cut – cut out selected portions of each line of a file.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/mem"
	"solod.dev/so/os"
	"solod.dev/so/strconv"
	"solod.dev/so/strings"
)

const illegalListValue = "cut: [-f] list: illegal list value"

// Range represents a field selection range (1-based).
// hi == 0 means "to end of line".
type Range struct {
	lo int
	hi int
}

var fields string
var delim string
var useWhitespace bool

var ranges [128]Range
var nranges int

func main() {
	parseFlags()

	if fields == "" {
		fatal("must specify field list (-f)")
	}
	if useWhitespace && delim != "\t" {
		fatal("-w and -d may not be used together")
	}

	parseFields(fields)

	outDelim := delim
	if useWhitespace {
		outDelim = "\t"
	}

	args := flag.Args()
	if len(args) == 0 {
		cut(os.Stdin, outDelim)
		os.Exit(0)
	}
	for _, fname := range args {
		if fname == "-" {
			cut(os.Stdin, outDelim)
		} else {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Printf("cut: %s: No such file or directory\n", fname)
				os.Exit(1)
			}
			cut(&f, outDelim)
			f.Close()
		}
	}
}

// parseFlags parses command-line flags.
func parseFlags() {
	flag.StringVar(&fields, "f", "", "field list")
	flag.StringVar(&delim, "d", "\t", "field delimiter")
	flag.BoolVar(&useWhitespace, "w", false, "use whitespace as delimiter")
	flag.Parse()
}

// parseFields parses the field list specified by the -f flag.
func parseFields(list string) {
	nranges = 0
	i := 0
	n := len(list)
	for i < n {
		// Skip separators (comma, space, tab).
		for i < n && (list[i] == ',' || list[i] == ' ' || list[i] == '\t') {
			i++
		}
		if i >= n {
			break
		}
		start := i
		for i < n && list[i] != ',' && list[i] != ' ' && list[i] != '\t' {
			i++
		}
		parseToken(list[start:i])
	}
	if nranges == 0 {
		fatal(illegalListValue)
	}
}

// parseToken parses a single token from the field list,
// which can be one of the following forms:
//   - N: a single number, meaning field N
//   - -N: a leading dash, meaning fields 1 through N
//   - N-: a trailing dash, meaning fields N through end of line
//   - N-M: a range, meaning fields N through M
//
// After parsing, the token is added to the ranges array.
func parseToken(token string) {
	if len(token) == 0 {
		return
	}
	dashIdx := strings.IndexByte(token, '-')
	if dashIdx < 0 {
		// Single number: N.
		n, err := strconv.Atoi(token)
		if err != nil || n <= 0 {
			fatal(illegalListValue)
		}
		ranges[nranges] = Range{n, n}
		nranges++
		return
	}
	if dashIdx == 0 {
		// Leading dash: -N means fields 1 through N.
		if len(token) < 2 {
			fatal(illegalListValue)
		}
		n, err := strconv.Atoi(token[1:])
		if err != nil || n <= 0 {
			fatal(illegalListValue)
		}
		ranges[nranges] = Range{1, n}
		nranges++
		return
	}
	if dashIdx == len(token)-1 {
		// Trailing dash: N- means fields N through end.
		n, err := strconv.Atoi(token[:dashIdx])
		if err != nil || n <= 0 {
			fatal(illegalListValue)
		}
		ranges[nranges] = Range{n, 0}
		nranges++
		return
	}
	// Range: N-M.
	lo, err := strconv.Atoi(token[:dashIdx])
	if err != nil || lo <= 0 {
		fatal(illegalListValue)
	}
	hi, err := strconv.Atoi(token[dashIdx+1:])
	if err != nil || hi <= 0 {
		fatal(illegalListValue)
	}
	ranges[nranges] = Range{lo, hi}
	nranges++
}

// cut reads lines from r, extracts the selected fields,
// and writes them to standard output.
func cut(r io.Reader, outDelim string) {
	scanner := bufio.NewScanner(mem.System, r)
	for scanner.Scan() {
		cutLine(scanner.Text(), outDelim)
	}
	scanner.Free()
}

// cutLine extracts the selected fields from a single line
// and writes them to standard output.
func cutLine(line string, outDelim string) {
	var fields []string
	if useWhitespace {
		fields = strings.Fields(mem.System, line)
	} else {
		fields = strings.Split(mem.System, line, delim)
	}
	nfields := len(fields)

	first := true
	for i := 1; i <= nfields; i++ {
		if isSelected(i, nfields) {
			if !first {
				fmt.Print(outDelim)
			}
			fmt.Print(fields[i-1])
			first = false
		}
	}
	fmt.Println("")
	mem.FreeSlice(mem.System, fields)
}

// isSelected returns true if the given field number
// is selected by the field list.
func isSelected(field int, nfields int) bool {
	for i := 0; i < nranges; i++ {
		lo := ranges[i].lo
		hi := ranges[i].hi
		if hi == 0 {
			hi = nfields
		}
		if field >= lo && field <= hi {
			return true
		}
	}
	return false
}

func fatal(msg string) {
	fmt.Printf("cut: %s\n", msg)
	os.Exit(1)
}
