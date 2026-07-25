// We can use channels to synchronize execution across threads.
// Here's an example of using a blocking receive to wait for a
// thread to finish.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/time"
)

// This is the function we'll run in a thread. The `done` channel
// notifies the main thread that this function's work is done.
func worker(arg any) any {
	done := arg.(*conc.Chan[bool])
	print("working...")
	time.Sleep(time.Second)
	println("done")

	// Send a value to notify that we're done.
	done.Send(true)
	return nil
}

func main() {
	// Start a worker thread, giving it the channel to notify on.
	done := conc.NewChan[bool](mem.System, 0)
	defer done.Free()
	th := conc.Go(worker, &done)
	defer th.Wait()

	// Block until we receive a notification from the worker.
	var ok bool
	done.Recv(&ok)
}
