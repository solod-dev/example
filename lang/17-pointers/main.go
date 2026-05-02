// A pointer stores a memory address that refers to a specific value.
// The type *T is a pointer to a value of type T. The zero value of a
// pointer is nil.
package main

func main() {
	// `iptr` is a pointer to an int value; right now it's empty.
	var iptr *int
	println("iptr =", iptr)

	// The `&`` operator returns a pointer to a specific value:
	i := 42
	iptr = &i
	// Now `iptr` points to `i` - it holds the memory address of `i`.
	println("iptr =", iptr)

	// The `*` operator dereferences a pointer, accessing the value it points to.
	// You can read or change the value through the pointer.

	// Read the value of `i` through the pointer `iptr`.
	println("*iptr =", *iptr)
	// Set the value of `i` through the pointer `iptr`.
	*iptr = 21
	println("i =", i)

	// `any` can hold any pointer value;
	// it translates to (void*) in C.
	var n byte = 15
	var a any = n // void* a = &n
	println("any pointer =", a)

	// `any` can be converted back to a pointer of the original type,
	// or to a pointer of a different type (at your own risk).
	b := a.(*byte) // uint32_t* b = (uint32_t*)a
	println("any as byte =", *b)
	i32 := a.(*int32) // int32_t* i32 = (int32_t*)a
	println("any as int32 =", *i32)
}
