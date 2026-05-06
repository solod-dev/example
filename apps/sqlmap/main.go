// A simple key-value store backed by an SQLite database.
// Run with `LDFLAGS="-lsqlite3" so run <args>`.
package main

import (
	"solod.dev/so/flag"
	"solod.dev/so/mem"
	"solod.dev/so/os"
)

var (
	opFlag  string
	keyFlag string
	valFlag string
)

func parseFlags() {
	flag.StringVar(&opFlag, "op", "", "operation: get, set, or del")
	flag.StringVar(&keyFlag, "key", "", "key name")
	flag.StringVar(&valFlag, "val", "", "value (for set operation)")
	flag.Parse()
}

func main() {
	parseFlags()

	m, err := NewSQLMap("sqlmap.db")
	check(err)
	defer m.Close()

	switch opFlag {
	case "set":
		err = m.SetString(keyFlag, valFlag)
		check(err)
	case "get":
		val, err := m.GetString(mem.System, keyFlag)
		check(err)
		println(val)
		mem.FreeString(mem.System, val)
	case "del":
		err = m.Delete(keyFlag)
		check(err)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil && err != ErrNotFound {
		panic(err)
	}
}
