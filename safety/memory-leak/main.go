// `mem.Tracker` helps catch memory leaks: it wraps any allocator and keeps
// track of every allocation and free that goes through it. This way, you can
// monitor your program's memory usage in real time instead of guessing.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/time"
)

// monitor periodically logs live allocation stats.
func monitor(arg any) any {
	t := arg.(*mem.Tracker)
	for {
		time.Sleep(100 * time.Millisecond)
		s := t.Stats()
		println("live:", s.Mallocs-s.Frees, "allocations,", s.Alloc, "bytes")
	}
	return nil
}

func main() {
	// Wrap the system allocator to count every allocation and free.
	heap := &mem.Tracker{Allocator: mem.System}

	// Watch memory from a background thread.
	conc.Go(monitor, heap).Detach()

	// Allocate through heap so the monitor sees it.
	for i := range 10 {
		v := mem.Alloc[int](heap) // intentionally not freeing it
		*v = i
		time.Sleep(50 * time.Millisecond)
	}
	// ...
}
