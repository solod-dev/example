// In this example we'll look at how to implement a worker pool
// using `conc.Pool`. A pool runs a fixed number of worker threads
// that pull jobs from a queue, which is the recommended approach
// for a large number of short-lived jobs.
//
// For an alternative worker pool implementation with channels,
// see the chan-pool example.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/time"
)

// job holds input and the result, written in place.
type job struct {
	id     int
	result int
}

// process handles a single job. We sleep a second to simulate an
// expensive job, then write the result (the input doubled).
func process(arg any) {
	j := arg.(*job)
	println("started job", j.id)
	time.Sleep(time.Second)
	println("finished job", j.id)
	j.result = j.id * 2
}

func main() {
	const numJobs = 5

	// Start a pool of 3 worker threads. Each submitted job is
	// picked up by whichever worker is free; we don't manage the
	// workers or track which one runs a given job.
	pool := conc.NewPool(mem.System, conc.PoolOptions{NumThreads: 3})
	defer pool.Free()

	// Submit 5 jobs. Each job writes into its own struct, so we
	// keep the structs alive in a slice until the jobs finish.
	jobs := make([]job, numJobs)
	for i := range jobs {
		jobs[i].id = i + 1
		pool.Go(process, &jobs[i])
	}

	// Wait blocks until all submitted jobs have finished.
	pool.Wait()

	// Now that every job is done, collect the results.
	for i := range jobs {
		println("result for job", jobs[i].id, "is", jobs[i].result)
	}

	// This program takes about 2 seconds despite doing 5 seconds of
	// total work, because 3 workers run concurrently.
}
