// Branching with `if` and `else` in So is straightforward.
package main

func main() {
	// You don't need parentheses around the condition,
	// but that the braces are required.
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	// You can have an `if` statement without an `else`.
	if 8%4 == 0 {
		println("8 is divisible by 4")
	}

	// Logical operators like `&&` and `||` are often
	// useful in conditions.
	if 8%2 == 0 || 7%2 == 0 {
		println("either 8 or 7 are even")
	}

	// You can put a statement before a condition. Any variables
	// declared in that statement will be available in the current
	// branch and all branches that follow.
	if num := 9; num < 0 {
		println(num, "is negative")
	} else if num < 10 {
		println(num, "has 1 digit")
	} else {
		println(num, "has multiple digits")
	}
}
