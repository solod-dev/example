// Count the frequency of each word in a text file
// and print the top N most common words.
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/c"
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/io"
	"solod.dev/so/maps"
	"solod.dev/so/mem"
	"solod.dev/so/os"
	"solod.dev/so/slices"
	"solod.dev/so/strings"
)

func main() {
	var top int
	flag.IntVar(&top, "top", 10, "number of top words to display")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: wordfreq [-top N] <file>")
		os.Exit(1)
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("Error: cannot open %s\n", args[0])
		os.Exit(1)
	}
	defer file.Close()

	entries := countWords(&file)
	defer freeEntries(entries)

	n := min(top, len(entries))
	for _, entry := range entries[:n] {
		fmt.Printf("%d %s\n", entry.count, entry.word)
	}
}

// wordCount holds a word and its frequency count.
type wordCount struct {
	word  string
	count int
}

// countWords reads words from the reader, counts their frequencies,
// and returns a slice of word-count pairs sorted by count descending.
func countWords(r io.Reader) []wordCount {
	scanner := bufio.NewScanner(mem.System, r)
	defer scanner.Free()
	scanner.Split(bufio.ScanWords)

	// Count the frequency of each word using a map.
	wordMap := maps.New[string, int](mem.System, 8)
	defer wordMap.Free()

	for scanner.Scan() {
		word := scanner.Text()
		if wordMap.Has(word) {
			n := wordMap.Get(word)
			wordMap.Set(word, n+1)
		} else {
			// Clone the word to ensure it outlives the scanner's buffer.
			word = strings.Clone(mem.System, word)
			wordMap.Set(word, 1)
		}
	}

	// Collect word-count pairs into a slice and sort by count descending.
	entries := slices.MakeCap[wordCount](mem.System, 0, wordMap.Len())
	it := wordMap.Iter()
	for it.Next() {
		entry := wordCount{word: it.Key(), count: it.Value()}
		entries = slices.Append(mem.System, entries, entry)
	}
	slices.SortFunc(entries, compareByCount)

	return entries
}

// compareByCount compares two word-count entries by count descending.
func compareByCount(a, b any) int {
	wa := c.PtrAs[wordCount](a)
	wb := c.PtrAs[wordCount](b)
	if wa.count != wb.count {
		return wb.count - wa.count
	}
	return strings.Compare(wa.word, wb.word)
}

// freeEntries frees the memory used by the word-count entries.
func freeEntries(entries []wordCount) {
	if len(entries) == 0 {
		return
	}
	for _, entry := range entries {
		mem.FreeString(mem.System, entry.word)
	}
	slices.Free(mem.System, entries)
}
