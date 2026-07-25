// The primary mechanism for managing state in So is communication
// over channels, but there are also other options. Here we'll look at
// using the `so/sync/atomic` package for atomic counters accessed
// by multiple threads.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/sync/atomic"
)

// increment atomically bumps the shared counter 1000 times.
// Every task receives a pointer to the same `ops`, so the
// updates accumulate across all of them.
func increment(arg any) {
	ops := arg.(*atomic.Uint64)
	for range 1000 {
		// To atomically increment the counter we use `Add`.
		ops.Add(1)
	}
}

func main() {
	// We'll use an atomic integer type to represent our
	// (always-positive) counter.
	var ops atomic.Uint64

	// A pool of worker threads runs the tasks; `Wait` lets us
	// wait for all of them to finish, like Go's WaitGroup.
	pool := conc.NewPool(mem.System, conc.PoolOptions{NumThreads: 8})
	defer pool.Free()

	// We'll start 50 tasks that each increment the counter
	// exactly 1000 times.
	for range 50 {
		pool.Go(increment, &ops)
	}

	// Wait until all the tasks are done.
	pool.Wait()

	// Here no tasks are writing to `ops`, but using `Load` it's
	// safe to atomically read a value even while other threads
	// are (atomically) updating it.
	println("ops:", ops.Load())

	// We expect to get exactly 50,000 operations. Had we used a
	// non-atomic integer and incremented it with `ops++`, we'd
	// likely get a different number, changing between runs,
	// because the threads would interfere with each other.
}
