// The `so/maps` package provides a generic hashmap implementation
// similar to Go's built-in map[K]V, but with explicit allocation.
package main

import (
	"solod.dev/so/maps"
	"solod.dev/so/mem"
)

func main() {
	// maps.New creates a new map with a given initial capacity.
	// The first argument is an allocator (nil uses the system allocator).
	m := maps.New[string, int](mem.System, 8)
	defer m.Free()

	// Set adds or updates key-value pairs.
	m.Set("alice", 25)
	m.Set("bob", 30)
	m.Set("carol", 28)

	// Get retrieves a value by key (returns zero value if missing).
	println("alice =", m.Get("alice"))
	println("bob =", m.Get("bob"))
	println("missing =", m.Get("dave"))

	// Has checks whether a key exists.
	println("has alice =", m.Has("alice"))
	println("has dave =", m.Has("dave"))

	// Len returns the number of entries.
	println("len =", m.Len())

	// Iter iterates over all key-value pairs.
	// Order is not guaranteed.
	m.Set("bob", 31)
	println("all entries:")
	it := m.Iter()
	for it.Next() {
		println("  ", it.Key(), "=", it.Value())
	}

	// Delete removes a key-value pair.
	m.Delete("bob")
	println("after delete bob:")
	println("  has bob =", m.Has("bob"))
	println("  len =", m.Len())

	// Clear removes all entries but keeps allocated memory.
	m.Clear()
	println("after clear: len =", m.Len())
}
