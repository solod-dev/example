// The `so/strings` package provides string manipulation
// functions, similar to Go's `strings` package.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/strings"
)

func main() {
	contains := strings.Contains("test", "es")
	fmt.Printf("Contains(test, es) = %d\n", contains)

	count := strings.Count("test", "t")
	fmt.Printf("Count(test, t) = %d\n", int32(count))

	hasPrefix := strings.HasPrefix("test", "te")
	fmt.Printf("HasPrefix(test, te) = %d\n", hasPrefix)

	hasSuffix := strings.HasSuffix("test", "st")
	fmt.Printf("HasSuffix(test, st) = %d\n", hasSuffix)

	index := strings.Index("test", "e")
	fmt.Printf("Index(test, e) = %d\n", int32(index))

	// Some functions allocate new strings to return the result.
	// Use mem.FreeString to free them and avoid memory leaks.
	joined := strings.Join(mem.System, []string{"a", "b"}, "-")
	defer mem.FreeString(mem.System, joined)
	fmt.Printf("Join(a, b, -) = %s\n", joined)

	repeated := strings.Repeat(mem.System, "a", 5)
	defer mem.FreeString(mem.System, repeated)
	fmt.Printf("Repeat(a, 5) = %s\n", repeated)

	replacedAll := strings.ReplaceAll(mem.System, "foo", "o", "0")
	defer mem.FreeString(mem.System, replacedAll)
	fmt.Printf("ReplaceAll(foo, o, 0) = %s\n", replacedAll)

	replacedOnce := strings.Replace(mem.System, "foo", "o", "0", 1)
	defer mem.FreeString(mem.System, replacedOnce)
	fmt.Printf("Replace(foo, o, 0, 1) = %s\n", replacedOnce)

	splitted := strings.Split(mem.System, "a-b-c-d-e", "-")
	defer mem.FreeSlice(mem.System, splitted)
	fmt.Println("Split(a-b-c-d-e, -):")
	for i, s := range splitted {
		fmt.Printf("  [%d] = %s\n", i, s)
	}

	lowered := strings.ToLower(mem.System, "TEST")
	defer mem.FreeString(mem.System, lowered)
	fmt.Printf("ToLower(TEST) = %s\n", lowered)

	uppered := strings.ToUpper(mem.System, "test")
	defer mem.FreeString(mem.System, uppered)
	fmt.Printf("ToUpper(test) = %s\n", uppered)
}
