// `switch` represents a condition with multiple branches.
package main

import "solod.dev/so/time"

func main() {
	{
		// Unlike C, you don't need to use `break`. So only runs
		// the matching case and doesn't fall through to the next one.
		i := 2
		print("Write", i, "as ")
		switch i {
		case 1:
			println("one")
		case 2:
			println("two")
		case 3:
			println("three")
		}
	}

	// Unlike Go, So doesn't support the `fallthrough` keyword.

	{
		// You can include multiple expressions in one branch.
		// The `default` branch will run if none of the others match.
		switch day := time.Now().Weekday(); day {
		case time.Saturday, time.Sunday:
			println(day, "is a weekend")
		default:
			println(day, "is a weekday")
		}
	}

	{
		// Expressions in case branches don't have to be constants.
		// A `switch` can be used like an `if`.
		t := time.Now()
		switch {
		case t.Hour() < 12:
			println("It's before noon")
		default:
			println("It's after noon")
		}
	}
}
