// A struct groups fields together into a single record.
package main

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

// `newPerson` constructs a new person struct with the given name.
func newPerson(name string) person {
	// Unlike Go, So is NOT a garbage collected language.
	// You should never return a pointer to a local stack-allocated
	// variable - the stack is freed when the function returns.
	p := person{name: name}
	p.age = 42
	return p // value, not pointer
}

func main() {
	// Create a new struct value using a struct literal.
	p := person{"Bob", 20}
	printPerson(p)

	// You can name the fields when initializing a struct.
	p = person{name: "Alice", age: 30}
	printPerson(p)

	// Omitted fields will be zero-valued.
	p = person{name: "Fred"}
	printPerson(p)

	// The `&` operator yields a pointer to the struct.
	pptr := &person{name: "Ann", age: 40}
	printPerson(*pptr)

	// It's idiomatic to create new structs values
	// using constructor functions.
	p = newPerson("Jon")
	printPerson(p)

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	println("s.name =", s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	println("sp.age =", sp.age) // same as (*sp).age

	// Structs are mutable.
	sp.age = 51
	println("sp.age =", sp.age)

	// If a struct type is only used for one value, we don't need
	// to give it a name. The value can use an anonymous struct type.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	println(dog.name, dog.isGood)
}

func printPerson(p person) {
	println("{", p.name, p.age, "}")
}
