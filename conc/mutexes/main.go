// In the atomic-counters example we saw how to manage simple counter
// state using atomic operations. For more complex state we can
// use a mutex to safely access data across multiple threads.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/sync"
)

// Container holds a map of counters; since we want to update it
// concurrently from multiple threads, we add a `Mutex` to
// synchronize access. Mutexes must not be copied, so the
// container is always passed around by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing `counters`; unlock it at
	// the end of the function using a `defer` statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	// So has no `m[k]++`, so we spell out the read-modify-write.
	c.counters[name] = c.counters[name] + 1
}

// job asks a worker to increment a counter
// with a given name n times.
type job struct {
	c    *Container
	name string
	n    int
}

// increment increments a named counter in a loop.
func increment(arg any) {
	j := arg.(*job)
	for range j.n {
		j.c.inc(j.name)
	}
}

func main() {
	c := Container{counters: map[string]int{"a": 0, "b": 0}}
	// Unlike Go, So's mutex zero value is not ready to use: we
	// initialize it here and free it when we're done.
	c.mu.Init()
	defer c.mu.Free()

	// A pool of worker threads runs the jobs; `Wait` lets us
	// wait for all of them, like Go's WaitGroup.
	pool := conc.NewPool(mem.System, conc.PoolOptions{NumThreads: 3})
	defer pool.Free()

	// Run several jobs concurrently; they all access the same
	// `Container`, and two of them access the same counter.
	// Each job needs its own storage that outlives the thread,
	// so we keep the structs in a slice.
	jobs := []job{
		{c: &c, name: "a", n: 10000},
		{c: &c, name: "a", n: 10000},
		{c: &c, name: "b", n: 10000},
	}
	for i := range jobs {
		pool.Go(increment, &jobs[i])
	}

	// Wait for the jobs to finish, then read the final counts.
	pool.Wait()
	println("a:", c.counters["a"], "b:", c.counters["b"])
}
