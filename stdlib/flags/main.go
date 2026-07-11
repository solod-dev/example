// The `so/flag` package provides a way
// to parse command-line arguments.
package main

import (
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/os"
)

func main() {
	// Basic flag declarations are available for string,
	// integer, and boolean options.

	// Here we declare a string flag `word` with a default
	// value "foo" and a short description. This flag.StringVar
	// function takes a pointer to a string variable which will
	// hold the value of the flag after parsing.
	var word string
	flag.StringVar(&word, "word", "foo", "a string")

	// Declare `num` and `ok` flags using a similar approach.
	var num int
	flag.IntVar(&num, "num", 42, "an int")
	var ok bool
	flag.BoolVar(&ok, "ok", false, "a bool")

	// Once all flags are declared, call flag.Parse
	// to parse the command-line arguments.
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments.
	fmt.Printf("word = %s\n", word)
	fmt.Printf("num = %d\n", num)
	fmt.Printf("ok = %d\n", ok)
	for i, arg := range flag.Args() {
		fmt.Printf("tail[%d] = %s\n", i, arg)
	}

	// The raw command-line arguments are available in os.Args.
	fmt.Println("os.Args:")
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
