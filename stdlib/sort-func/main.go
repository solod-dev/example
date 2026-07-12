// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts in So.
package main

import (
	"solod.dev/so/cmp"
	"solod.dev/so/fmt"
	"solod.dev/so/slices"
)

// We implement a comparison function for string lengths.
// A `cmp.Func` receives its arguments as `any` pointing to
// the slice elements, so we type-assert them back to
// `string`.
func lenCmp(a, b any) int {
	sa := a.(string)
	sb := b.(string)
	return cmp.Compare(len(sa), len(sb))
}

// We can use the same technique to sort a slice of values
// that aren't built-in types.
type Person struct {
	name string
	age  int
}

// ageCmp compares two people by age.
func ageCmp(a, b any) int {
	pa := a.(*Person)
	pb := b.(*Person)
	return cmp.Compare(pa.age, pb.age)
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Now we can call `slices.SortFunc` with our custom
	// comparison function to sort `fruits` by name length.
	slices.SortFunc(fruits, lenCmp)
	fmt.Printf("[%s %s %s]\n", fruits[0], fruits[1], fruits[2])

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// Sort `people` by age using `slices.SortFunc`.
	slices.SortFunc(people, ageCmp)
	fmt.Printf("[{%s %d} {%s %d} {%s %d}]\n",
		people[0].name, people[0].age,
		people[1].name, people[1].age,
		people[2].name, people[2].age)
}
