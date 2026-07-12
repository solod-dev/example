// A common requirement in programs is getting the number of
// seconds, milliseconds, or nanoseconds since the Unix epoch.
// Here's how to do it in So.

package main

import "solod.dev/so/time"

func main() {
	// `Time.String` formats a time into a caller-provided
	// buffer, so we allocate one big enough for the ISO 8601
	// representation and reuse it below.
	buf := make([]byte, time.RFC3339Len)

	// Use `time.Now` with `Unix`, `UnixMilli` or `UnixNano`
	// to get elapsed time since the Unix epoch in seconds,
	// milliseconds or nanoseconds, respectively.
	now := time.Now()
	println(now.String(buf))

	println(now.Unix())
	println(now.UnixMilli())
	println(now.UnixNano())

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	unixNow := time.Unix(now.Unix(), 0)
	println(unixNow.String(buf))
	unixNanoNow := time.Unix(0, now.UnixNano())
	println(unixNanoNow.String(buf))
}
