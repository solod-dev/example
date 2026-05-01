// Methods add behavior to struct types.
package main

// rect represents a rectangle figure.
type rect struct {
	width, height int
}

// This `area` method has a receiver type of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Call methods defined on `rect`.
	println("area =", r.area()) // same as (&r).area()
	println("perim =", r.perim())

	// So automatically converts between values and pointers when
	// calling methods. You might want to use a pointer receiver
	// to avoid copying the struct when calling methods or to let
	// the method change the struct.
	rp := &r
	println("area =", rp.area())
	println("perim =", rp.perim()) // same as (*rp).perim()
}
