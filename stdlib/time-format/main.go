// So supports time formatting and parsing via
// strftime/strptime-style layouts.
package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/time"
)

func main() {
	// `Time.Format` writes into a caller-provided buffer, so we
	// allocate one large enough for the longest layout below.
	buf := make([]byte, time.RFC3339NanoLen)

	// Here's a basic example of formatting a time according to RFC3339,
	// using the corresponding layout constant. The offset argument
	// says which timezone to render in; we use `time.UTC`.
	t := time.Now()
	println(t.Format(buf, time.RFC3339, time.UTC))

	// Time parsing uses the same layout values as `Format`.
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00", time.UTC)
	println(t1.String(buf))

	// `Format` and `Parse` use strftime/strptime verbs. Usually
	// you'll use a constant from `time` for these layouts, like
	// RFC3339, DateOnly or TimeOnly.
	println(t.Format(buf, time.RFC3339Nano, time.UTC))
	println(t.Format(buf, time.DateOnly, time.UTC))
	println(t.Format(buf, time.TimeOnly, time.UTC))
	// You can also use a custom layout. Here we format the time
	// as "hour:minute" in 24-hour format in the UTC+3 timezone.
	println(t.Format(buf, "%H:%M", time.Offset(3*3600)))

	// For purely numeric representations you can also use
	// standard string formatting with the extracted components
	// of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` will return an error on malformed input.
	_, err := time.Parse("%a %b %e %H:%M:%S %Y", "8:41PM", time.UTC)
	println(err)
}
