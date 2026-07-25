// Channels are the pipes that connect concurrent threads.
// You can send values into a channel from one thread and
// receive those values in another.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
)

// ping sends a single message on the channel it is given.
func ping(arg any) any {
	messages := arg.(*conc.Chan[string])
	messages.Send("ping")
	return nil
}

func main() {
	// Create a new channel with `conc.NewChan`. Channels are
	// typed by the values they convey. The last argument is the
	// buffer size; 0 makes an unbuffered channel, where each send
	// blocks until a receiver is ready to take the value.
	messages := conc.NewChan[string](mem.System, 0)
	// Free the channel's resources once we are done with it.
	defer messages.Free()

	// Send a value into the channel with `Send`. Here we launch a
	// thread that sends "ping" into the `messages` channel.
	th := conc.Go(ping, &messages)
	defer th.Wait()

	// Receive a value from the channel with `Recv`. Here we receive
	// the "ping" message we sent above and print it out.
	var msg string
	messages.Recv(&msg)
	println(msg)

	// By default sends and receives block until both the sender and
	// receiver are ready. That is how we were able to wait for the
	// "ping" here without any other synchronization.
}
