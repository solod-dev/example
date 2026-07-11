// The `so/time` package provides time-related types
// and functions, similar to Go's `time` package.
package main

import "solod.dev/so/time"

func main() {
	// Pre-allocate a buffer for string formatting.
	buf := make([]byte, 64)

	// We'll start by getting the current time.
	now := time.Now()
	println("now =", now.String(buf))

	// You can build a Time struct by providing the year, month, day,
	// etc. Times are always in UTC. You can provide a UTC offset in
	// seconds, and the time will be adjusted accordingly.
	then := time.Date(2026, 2, 1, 12, 34, 56, 555555555, time.UTC)
	println("then =", then.String(buf))

	// Extract the various components of the time value.
	println("now year =", now.Year())
	println("now month =", now.Month())
	println("now day =", now.Day())
	println("now hour =", now.Hour())
	println("now minute =", now.Minute())
	println("now second =", now.Second())
	println("now nano =", now.Nanosecond())

	// Weekday returns 1 for Monday, 2 for Tuesday, and so on.
	println("now weekday =", now.Weekday())

	// These methods compare two times to check if the first one
	// happens before, after, or at the same time as the second.
	println("then before now =", then.Before(now))
	println("then after now =", then.After(now))
	println("then equal now =", then.Equal(now))

	// The Sub methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	println("diff (now - then) =", diff.String(buf))

	// Calculate the duration length in different units.
	println("diff hours =", diff.Hours())
	println("diff minutes =", diff.Minutes())
	println("diff seconds =", diff.Seconds())
	println("diff nanos =", diff.Nanoseconds())

	// Use Add to advance a time by a given duration
	// (or go back by using a negative duration).
	println("then + diff =", then.Add(diff).String(buf))
	println("then - diff =", then.Add(-diff).String(buf))
}
