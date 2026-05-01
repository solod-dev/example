// Maps are So's built-in associative (key-value) data type.
// Built-in maps in So are fixed-size and stack-allocated,
// so only use them when you have a small, fixed number of
// key-value pairs (<1024). For anything else, use heap-allocated
// maps from the `so/maps` package.
package main

import "solod.dev/so/maps"

func main() {
	// To create a map, use the `make` builtin:
	m := make(map[string]int, 3) // 3 is the fixed map capacity

	// Set key/value pairs using typical `name[key] = val` syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	println("m[k1] =", v1)

	// If the key doesn't exist, the zero value of the
	// value type is returned.
	v3 := m["k3"]
	println("m[k3] =", v3)

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	println("len(m) =", len(m))

	// `delete` and `clear` are not supported.
	// delete(m, "k2")
	// clear(m)

	// The optional second return value when getting a value from a map
	// indicates if the key was present in the map. This can be used to
	// disambiguate between missing keys and keys with zero values like
	// `0` or `""`. Here we didn't need the value itself, so we ignored
	// it with the blank identifier `_`.
	_, ok := m["k2"]
	println("k2 in m?", ok)

	// You can also declare and initialize a new map as literal.
	n := map[string]int{"foo": 1, "bar": 2}
	println("foo:", n["foo"], "bar:", n["bar"])

	// Use the so/maps package for heap-allocated maps
	// that can grow and shrink dynamically.
	mapa := maps.New[string, int](nil, 0)
	defer mapa.Free() // remember to free heap-allocated maps
	mapa.Set("abc", 11)
	mapa.Set("def", 22)
	mapa.Set("xyz", 33)
	println("len(mapa) =", mapa.Len())
	println("mapa[abc] =", mapa.Get("abc"))
}
