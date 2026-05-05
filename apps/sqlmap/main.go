// A simple key-value store backed by an SQLite database.
// Run with `LDFLAGS="-lsqlite3" so run apps/sqlmap`.
package main

import "solod.dev/so/mem"

func main() {
	m, err := NewSQLMap(mem.System, ":memory:")
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.SetString("name", "Alice")
	m.SetInt("age", 42)

	name, err := m.GetString("name")
	println("name =", name, "err =", err)
	mem.FreeString(m.Alloc, name)
	// name = Alice, err = <nil>

	age, err := m.GetInt("age")
	println("age =", age, "err =", err)
	// age = 42, err = <nil>

	m.SetString("name", "Bob")
	name, err = m.GetString("name")
	println("name =", name, "err =", err)
	mem.FreeString(m.Alloc, name)
	// name = Bob, err = <nil>

	m.Delete("name")
	name, err = m.GetString("name")
	println("name =", name, "err =", err)
	mem.FreeString(m.Alloc, name)
	// name = , err = sqlmap: not found
}
