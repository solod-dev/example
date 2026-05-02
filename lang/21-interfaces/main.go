// Interfaces are named collections of method signatures.
package main

import "solod.dev/so/math"

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// Let's implement this interface on `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface, you need to implement all the methods
// in the interface. Here we implement `geometry` on `rect`s.
func (r *rect) area() float64 {
	return r.width * r.height
}
func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// And here we implement `geometry` on `circle`s.
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, we can call any methods
// defined in that interface. Here's a generic `measure` function
// that uses this to work with any `geometry` value.
func measure(g geometry) {
	println("area =", g.area(), "perim =", g.perim())
}

// Sometimes it's useful to know the concrete type of an interface
// value. You can do this with a type assertion - `iface.(type)`.
func detectCircle(g geometry) {
	if _, ok := g.(*circle); ok {
		c := g.(*circle)
		println("this is a circle, radius =", c.radius)
	}
}

// Go also has a type switch construct `switch x := iface.(type)`,
// but So does not support it.

func main() {
	r := &rect{width: 3, height: 4}
	c := &circle{radius: 5}

	// The `circle` and `rect` struct types both implement the
	// `geometry` interface so we can use instances of these
	// structs as arguments to `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
