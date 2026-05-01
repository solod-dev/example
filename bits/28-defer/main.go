// Defer makes sure a function is called later in the program,
// usually to clean things up.
package main

import "solod.dev/so/mem"

type Point struct {
	x, y int
}

// Suppose we wanted to allocate a Point on the heap,
// use it, and then free the memory when we're done.
// Here's how we could do that with `defer`.
func main() {
	// Immediately after allocating a Point object on the heap,
	// we defer the deallocation of that object. `mem.Free` will be
	// executed at the end of the enclosing function (`main`).
	p := mem.Alloc[Point](mem.System) // p is allocated on the heap
	defer mem.Free(mem.System, p)

	p.x = 11
	p.y = 22
	println(p.x, p.y)
	// `mem.Free` will be called here.
}
