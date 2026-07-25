// The `so/conc` package provides the tools to work with OS threads.
package main

import "solod.dev/so/conc"

// A thread entry point takes its argument as an `any` and returns
// a result as an `any`. Here `f` prints its label three times.
func f(arg any) any {
	// `any` is always a pointer. By asserting it to the value type
	// we dereference the pointer and get the value. We can also do
	// `arg.(*string)` to get the original pointer.
	from := arg.(string)
	for i := range 3 {
		println(from, ":", i)
	}
	return nil
}

// g prints a single message. We use it to show that a thread
// can run any matching function, not just `f`.
func g(arg any) any {
	msg := arg.(string)
	println(msg)
	return nil
}

func main() {
	// Suppose we have a function call `f(s)`. Here's how we'd
	// call that in the usual way, running it synchronously.
	direct := "direct"
	f(direct)

	// To invoke this function on a separate OS thread, use
	// `conc.Go`. This new thread will execute concurrently
	// with the calling one. The argument must point to storage
	// that outlives the thread.
	thread := "thread"
	th := conc.Go(f, &thread)

	// You can also start a thread for a different function.
	// Even if you try to pass the argument by value rather than
	// by pointer (like we do here), So will implicitly pass a pointer
	// to it (`arg = &going`), so the value must still be addressable
	// and outlive the thread.
	going := "going"
	other := conc.Go(g, going)

	// Our two calls are running on separate threads now. Unlike
	// Go's goroutines, an OS thread must always be joined with
	// `Wait` (or handed to the runtime with `Detach`), otherwise
	// its resources leak. `Wait` blocks until the thread returns.
	th.Wait()
	other.Wait()

	println("done")

	// OS threads are not cheap to start, so spawning one per task
	// does not scale. For a large number of short-lived tasks, use
	// `conc.Pool` instead (see the worker-pool example).
}
