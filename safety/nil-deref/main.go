// If you try to dereference a nil pointer, it will cause a panic
// at runtime instead of a raw segmentation fault.
package main

type Rect struct{ width, height int }

func (r *Rect) area() int {
	return r.width * r.height
	//     ^ runtime error: nil pointer dereference
}

func main() {
	var r *Rect
	println(r.area())
}
