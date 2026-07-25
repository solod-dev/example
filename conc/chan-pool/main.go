// In this example we'll look at how to implement a worker pool
// using `conc.Go` and `conc.Chan`.
//
// For an alternative implementation with `conc.Pool`,
// see the worker-pool example.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/time"
)

// worker carries a worker's id and the input/output channels.
type worker struct {
	id      int
	jobs    conc.Chan[int]
	results conc.Chan[int]
}

// work is the worker's function, we run several of these concurrently.
// Each worker receives work from the `jobs` channel and sends results
// to the `results` channel. We sleep a second per job to simulate an
// expensive task.
func work(arg any) any {
	w := arg.(*worker)
	var j int
	for w.jobs.Recv(&j) {
		println("worker", w.id, "started job", j)
		time.Sleep(time.Second)
		println("worker", w.id, "finished job", j)
		w.results.Send(j * 2)
	}
	return nil
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	// Create channels for jobs (inputs) and results (outputs)
	// that our workers will use.
	jobs := conc.NewChan[int](mem.System, numJobs)
	defer jobs.Free()
	results := conc.NewChan[int](mem.System, numJobs)
	defer results.Free()

	// Start 3 workers, initially blocked because there are no jobs yet.
	// Each worker needs its own storage that outlives the thread, so we
	// keep the structs in a slice.
	workers := make([]worker, numWorkers)
	threads := make([]conc.Thread, numWorkers)
	for i := range workers {
		workers[i] = worker{id: i + 1, jobs: jobs, results: results}
		threads[i] = conc.Go(work, &workers[i])
	}

	// Send 5 jobs, then close `jobs` to signal there is no more work.
	for j := 1; j <= numJobs; j++ {
		jobs.Send(j)
	}
	jobs.Close()

	// Collect all the results.
	var r int
	for range numJobs {
		results.Recv(&r)
		println("got result:", r)
	}

	// Join the workers before the deferred `Free` calls run, so that no
	// worker is still inside `Recv` when the channels are freed.
	for i := range threads {
		threads[i].Wait()
	}

	// Like the pool version, this takes about 2 seconds despite
	// doing 5 seconds of total work, thanks to 3 concurrent workers.
}
