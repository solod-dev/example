// Go's `encoding/json` relies on reflection to marshal arbitrary structs.
// Solod has no reflection, and uses a different approach: a token-level API.
// You read and write one JSON token at a time, and the `Encoder` and `Decoder`
// types take care of the syntax — adding commas and colons, checking UTF-8,
// and rejecting bad input.
package main

import (
	"solod.dev/so/encoding/json"
	"solod.dev/so/mem"
	"solod.dev/so/strings"
)

func encode() {
	// Encoding is done through a series of calls
	// that match the structure of your document.
	out := make([]byte, 256)
	sb := strings.FixedBuilder(out)
	enc := json.NewEncoder(&sb)

	enc.BeginObject()
	enc.Str("name")
	enc.Str("Alice")
	enc.Str("age")
	enc.Int(25)
	enc.EndObject()
	enc.Flush()

	println(sb.String())
	// {"name":"Alice","age":25}
}

func decode() {
	// Decoding pulls one validated token at a time with `Next`.
	// You can check each token with `Kind` and read its value using
	// typed getters like `Str`, `Int`, or `Bool`
	src := `{"name":"Alice","age":25}`
	dec := json.NewDecoder(mem.System, []byte(src))
	defer dec.Free()

	var name string
	var age int64

	dec.Next() // the opening {
	for dec.Next() && dec.Kind() == json.KindString {
		switch dec.Str() {
		case "name":
			dec.Next()
			name = dec.Str()
		case "age":
			dec.Next()
			age = dec.Int()
		default:
			dec.Next()
			dec.Skip()
		}
	}

	println(name, age)
	// Alice 25
}

func main() {
	encode()
	decode()
}
