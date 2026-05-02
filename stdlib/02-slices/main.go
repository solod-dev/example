// The `so/slices` package provides various functions
// useful with slices of any type, including sorting,
// searching, and heap-allocated slice management.
package main

import (
	"solod.dev/so/mem"
	"solod.dev/so/slices"
)

func main() {
	// slices.Make allocates a slice on the heap.
	// The first argument is an allocator (nil uses the system allocator).
	s := slices.Make[int](mem.System, 5)
	defer slices.Free(mem.System, s)
	for i := range s {
		s[i] = (i + 1) * 10
	}
	println("s =", s[0], s[1], s[2], s[3], s[4])

	// slices.Append grows a heap-allocated slice, like Go's append.
	// It automatically reallocates as necessary.
	s = slices.Append(mem.System, s, 60, 70)
	println("s + 60 + 70 =", s[0], s[1], s[2], s[3], s[4], s[5], s[6])

	// slices.Extend appends all elements from another slice.
	// Also reallocates as necessary.
	a := slices.Make[int](mem.System, 3)
	defer slices.Free(mem.System, a)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println("a =", a[0], a[1], a[2])
	b := []int{4, 5, 6}
	println("b =", b[0], b[1], b[2])
	a = slices.Extend(mem.System, a, b)
	println("a + b =", a[0], a[1], a[2], a[3], a[4], a[5])

	// slices.MakeCap allocates with a specific length and capacity,
	// useful when you know you'll append later.
	prealloc := slices.MakeCap[int](mem.System, 0, 10)
	defer slices.Free(mem.System, prealloc)
	prealloc = slices.Append(mem.System, prealloc, 100, 200, 300)
	println("prealloc =", prealloc[0], prealloc[1], prealloc[2])
	println("prealloc len =", len(prealloc), "cap =", cap(prealloc))

	// slices.Contains checks whether a value is in the slice.
	println("s contains 30 =", slices.Contains(s, 30))
	println("s contains 99 =", slices.Contains(s, 99))

	// slices.Index returns the index of the first occurrence, or -1.
	println("s index of 40 =", slices.Index(s, 40))
	println("s index of 99 =", slices.Index(s, 99))

	// slices.Clone creates a shallow copy of a slice.
	c := slices.Clone(mem.System, s)
	defer slices.Free(mem.System, c)
	println("clone(s) =", c[0], c[1], c[2], c[3], c[4], c[5], c[6])

	// slices.Equal checks whether two slices have the same elements.
	println("s == clone?", slices.Equal(s, c))

	// slices.Min and slices.Max return the smallest/largest element.
	vals := []int{7, 2, 9, 4, 1}
	println("vals =", vals[0], vals[1], vals[2], vals[3], vals[4])
	println("min(vals) =", slices.Min(vals))
	println("max(vals) =", slices.Max(vals))

	// slices.Sort sorts a slice of ordered types in ascending order.
	slices.Sort(vals)
	println("vals asc =", vals[0], vals[1], vals[2], vals[3], vals[4])

	// slices.SortFunc sorts with a custom comparison function.
	// Here we sort in descending order using descInt defined above.
	slices.SortFunc(vals, descInt)
	println("vals desc =", vals[0], vals[1], vals[2], vals[3], vals[4])
}

// descInt is a comparison function for sorting ints in descending order.
func descInt(a, b any) int {
	return *b.(*int) - *a.(*int)
}
