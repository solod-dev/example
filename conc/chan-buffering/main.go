// By default channels are unbuffered, meaning that they
// will only accept sends if there is a corresponding receive
// ready to take the value. Buffered channels accept a
// limited number of values without a corresponding receiver.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
)

func main() {
	// Here we make a channel of strings buffering up to 2 values.
	messages := conc.NewChan[string](mem.System, 2)
	defer messages.Free()

	// Because this channel is buffered, we can send these values
	// into the channel without a corresponding concurrent receive.
	messages.Send("buffered")
	messages.Send("channel")

	// Later we can receive these two values as usual.
	var msg string
	messages.Recv(&msg)
	println(msg)
	messages.Recv(&msg)
	println(msg)
}
