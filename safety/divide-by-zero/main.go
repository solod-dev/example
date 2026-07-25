// Integer division or modulo by zero is undefined in C. On x86 it traps, but
// on arm64 it silently yields 0, so a bug computes garbage instead of failing.
// So guards every integer division whose divisor isn't a known-nonzero
// constant, turning a zero divisor into a diagnosable panic that honors
// `-panic`:
//
//	so run -panic=trace .  # print a stack trace, then exit(1)
//	so run -panic=exit  .  # just exit(1) after the message
//
// The guard is skipped where it can't fire: floating-point division (IEEE
// defines it) and a constant divisor (Go rejects a constant zero divisor at
// compile time), so `x / 2` and `y / 3.0` stay as plain C.
package main

// avg returns the mean of the values. It divides by len(vals), which is zero
// for an empty slice, so the call below panics a couple of frames deep.
func avg(vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum / len(vals)
}

func main() {
	scores := []int{}
	println(avg(scores))
}
