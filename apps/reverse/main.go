// Read a file and print its lines in reverse order (byte-wise).
package main

import (
	"solod.dev/so/bufio"
	"solod.dev/so/flag"
	"solod.dev/so/fmt"
	"solod.dev/so/mem"
	"solod.dev/so/os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("usage: reverse <file>")
		os.Exit(1)
	}

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Println("error: could not open file")
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(mem.System, &f)
	defer scanner.Free()

	for scanner.Scan() {
		b := scanner.Bytes()
		reverse(b)
		fmt.Println(string(b))
	}
}

func reverse(b []byte) {
	i := 0
	j := len(b) - 1
	for i < j {
		tmp := b[i]
		b[i] = b[j]
		b[j] = tmp
		i++
		j--
	}
}
