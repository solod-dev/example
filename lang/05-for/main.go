// `for` is the only kind of loop in So.
// Here are some examples of how to use it.
package main

func main() {
	// The most basic kind, with a single condition.
	i := 1
	for i <= 3 {
		println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		println(j)
	}

	// Loop from 0 to n-1 (range over integers).
	const n = 10
	for i := range n {
		print(i)
	}
	println()

	// Range also works without a loop variable, if you don't need it.
	for range n {
		print(".")
	}
	println()

	// Infinite loop runs until a `break`` statement exits the loop
	// or a `return` statement exits the function.
	for {
		println("loop")
		break
	}

	// `continue` jumps  to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
