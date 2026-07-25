// Returning a pointer to a stack-allocated value is a classic C footgun.
// So catches the common cases at compile time.
package main

type Point struct{ x, y int }

func newPoint(x, y int) *Point {
	return &Point{x: x, y: y}
	//     ^ compile-time error: stack-allocated
	//       value escapes function frame
}

func main() {
	p := newPoint(3, 4)
	println(p.x, p.y)
}
