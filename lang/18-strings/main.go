// A So string is a read-only slice of bytes. The language
// and standard library handle strings in a special way -
// as containers for text encoded in UTF-8.
//
// In other languages, strings are made of "characters".
// In So, the concept of a character is called a `rune` -
// it's an integer that represents a Unicode code point.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/unicode/utf8"
)

func main() {
	// `s` is a `string` assigned a literal value representing the
	// word "hello" in the Thai language. So string literals are
	// UTF-8 encoded text.
	const s = "สวัสดี"

	// Since a string is the same as `[]byte`, `len(s)`
	// gives you the number of raw bytes stored inside.
	println("len =", len(s))

	// Indexing into a string gives you the raw byte values
	// at each position. This loop prints the hex values of
	// all the bytes that make up the code points in `s`.
	for i := range len(s) {
		fmt.Printf("%x ", s[i])
	}
	println()

	// To count how many runes are in a string, we can use the
	// `so/utf8` package. This is an O(n) time operation, because
	// it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by multi-byte UTF-8 code
	// points, so the result of this count may be surprising.
	n := utf8.RuneCountInString(s)
	println("rune count =", n)

	// A `range` loop handles strings specially and yields
	// each rune along with its offset in the string.
	for idx, r := range s {
		fmt.Printf("0x%x starts at %d\n", r, int32(idx))
		examineRune(r)
	}
}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals.
	// You can compare a rune value to a rune literal directly.
	if r == 't' {
		println("found tee")
	} else if r == 'ส' {
		println("found so sua")
	}
}
