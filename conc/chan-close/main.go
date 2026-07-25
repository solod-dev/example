// Closing a channel indicates that no more values will be sent
// on it. This can be useful to communicate completion to the
// channel's receivers.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
)

// job carries the `jobs` channel that delivers work and the
// `done` channel used to signal completion back to main.
type job struct {
	jobs conc.Chan[int]
	done conc.Chan[bool]
}

// work repeatedly receives from `jobs`.
func work(arg any) any {
	j := arg.(*job)
	for {
		var n int
		// `Recv` returns false once the channel is closed and drained,
		// at which point we notify on `done` and return.
		if j.jobs.Recv(&n) {
			println("received job", n)
		} else {
			println("received all jobs")
			j.done.Send(true)
			return nil
		}
	}
}

func main() {
	// We use a `jobs` channel to communicate work from main to a
	// worker thread, and `done` to learn when the worker is done.
	jobs := conc.NewChan[int](mem.System, 5)
	defer jobs.Free()
	done := conc.NewChan[bool](mem.System, 0)
	defer done.Free()

	// Start the worker thread.
	j := job{jobs: jobs, done: done}
	th := conc.Go(work, &j)
	defer th.Wait()

	// Send 3 jobs to the worker over the `jobs` channel, then
	// close it to signal that there is no more work.
	for n := 1; n <= 3; n++ {
		jobs.Send(n)
		println("sent job", n)
	}
	jobs.Close()
	println("sent all jobs")

	// Await the worker using the synchronization approach
	// we saw in the chan-sync example.
	var ok bool
	done.Recv(&ok)

	// Receiving from a closed and drained channel immediately
	// returns false; there are no more jobs to get.
	var n int
	more := jobs.Recv(&n)
	println("received more jobs:", more)
}
