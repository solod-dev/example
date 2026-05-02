// The `so/math/rand` package provides functions for generating
// random numbers, similar to Go's `math/rand/v2` package.
package main

import "solod.dev/so/math/rand"

func main() {
	// rand.IntN returns a random integer n, 0 <= n < 100.
	print("random ints = ")
	print(rand.IntN(100), "")
	print(rand.IntN(100))
	println()

	// rand.Float64 returns a floating-point f, 0.0 <= f < 1.0.
	println("random float =", rand.Float64())

	// You can generate random floats in other ranges, like 5.0 <= f' < 10.0.
	print("random floats = ")
	print((rand.Float64()*5)+5, "")
	print((rand.Float64() * 5) + 5)
	println()

	// If you want a known seed, create a new rand.Source and pass
	// it into the rand.New constructor. NewPCG creates a new
	// permuted congruential generator (PCG) source that requires
	// a seed of two uint64 numbers.
	pcg := rand.NewPCG(42, 1024)
	rnd := rand.New(&pcg)
	print("random ints with custom seed = ")
	print(rnd.IntN(100), "")
	print(rnd.IntN(100), "")
	print(rnd.IntN(100))
	println()
}
