// Basic sends and receives on channels are blocking. However,
// passing a zero timeout to `SendTimeout` or `RecvTimeout`
// makes the operation non-blocking.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
)

func main() {
	messages := conc.NewChan[string](mem.System, 0)
	defer messages.Free()
	signals := conc.NewChan[bool](mem.System, 0)
	defer signals.Free()

	// Here's a non-blocking receive. If a value is available on
	// `messages`, `RecvTimeout` returns `conc.Ok`. If not, the
	// zero timeout makes it return `conc.Timeout` immediately.
	var msg string
	if messages.RecvTimeout(&msg, 0) == conc.Ok {
		println("received message", msg)
	} else {
		println("no message received")
	}

	// A non-blocking send works similarly. Here "hi" cannot be sent,
	// because the channel has no buffer and no receiver, so
	// `SendTimeout` returns `conc.Timeout` immediately.
	out := "hi"
	if messages.SendTimeout(out, 0) == conc.Ok {
		println("sent message", out)
	} else {
		println("no message sent")
	}

	// Since So does not have `select`, a multi-way non-blocking check
	// is just a sequence of non-blocking receives. Here we try both
	// `messages` and `signals`, falling through to "no activity".
	var sig bool
	if messages.RecvTimeout(&msg, 0) == conc.Ok {
		println("received message", msg)
	} else if signals.RecvTimeout(&sig, 0) == conc.Ok {
		println("received signal", sig)
	} else {
		println("no activity")
	}
}
