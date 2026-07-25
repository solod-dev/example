// Timeouts are important for programs that connect to external
// resources or that otherwise need to bound execution time.
// So implements them with the `RecvTimeout` and `SendTimeout`
// channel methods, which return a status instead of blocking forever.
package main

import (
	"solod.dev/so/conc"
	"solod.dev/so/mem"
	"solod.dev/so/time"
)

// call carries the channel to reply on and the result to send.
type call struct {
	ch  conc.Chan[string]
	msg string
}

// externalCall simulates a slow external call.
func externalCall(arg any) any {
	c := arg.(*call)
	time.Sleep(2 * time.Second)
	c.ch.Send(c.msg)
	return nil
}

func main() {
	// Suppose we're executing an external call that returns its
	// result on channel `c1` after 2s. The channel is buffered, so
	// the send in the thread is nonblocking, preventing a thread leak
	// if the result is never read.
	c1 := conc.NewChan[string](mem.System, 1)
	defer c1.Free()
	call1 := call{ch: c1, msg: "result 1"}
	th1 := conc.Go(externalCall, &call1)
	th1.Detach() // do not wait for the thread to finish

	// `RecvTimeout` awaits the result for up to 1s, returning
	// `conc.Ok` if it arrived or `conc.Timeout` if the deadline
	// passed first. Since the call takes 2s, we time out here.
	var res string
	if c1.RecvTimeout(&res, 1*time.Second) == conc.Ok {
		println(res)
	} else {
		println("timeout 1")
	}

	// If we allow a longer timeout of 3s, the receive from `c2`
	// succeeds and we print the result.
	c2 := conc.NewChan[string](mem.System, 1)
	defer c2.Free()
	call2 := call{ch: c2, msg: "result 2"}
	th2 := conc.Go(externalCall, &call2)
	th2.Detach()

	if c2.RecvTimeout(&res, 3*time.Second) == conc.Ok {
		println(res)
	} else {
		println("timeout 2")
	}
}
