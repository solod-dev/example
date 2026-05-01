// An enum is a type that has a fixed number of possible values,
// each with a distinct name. So doesn't have an enum type as a
// distinct language feature, so enums are  typically implemented
// using constants of a specific type.
package main

// Our enum type `ServerState` has an underlying `int` type.
type ServerState int

// The possible values for `ServerState` are defined as
// constants. The special keyword `iota` generates successive
// constant values automatically; in this case 0, 1, 2 and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

func main() {
	println("initial state:", StateIdle)
	state1 := transition(StateIdle)
	println("transitioned to:", state1)

	// If we have a value of type `int`, we cannot pass it to `transition` -
	// the compiler will complain about type mismatch. This provides some
	// degree of compile-time type safety for enums.

	state2 := transition(state1)
	println("transitioned to:", state2)
}

// transition emulates a state transition for a
// server; it takes the existing state and returns
// a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicates here to
		// determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic("unknown state")
	}
}
