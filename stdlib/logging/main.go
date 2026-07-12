// The `so/log/slog` package provides structured logging
// similar to Go's `log/slog` package.
package main

import (
	"solod.dev/so/log/slog"
	"solod.dev/so/os"
)

func main() {
	// Configure the logger to write to standard output
	// anything with level INFO or higher.
	handler := slog.NewTextHandler(os.Stdout, slog.LevelInfo)
	log := slog.New(&handler)

	// Log a message with two attributes: a string "user" and an integer "count".
	log.Info("hello world", slog.String("user", "john"), slog.Int("count", 42))

	// Debug messages are ignored.
	log.Debug("this won't be logged", slog.String("user", "john"))

	// You can also use the default logger (os.Stderr + LevelInfo).
	slog.Info("hello world", slog.Bool("ok", true))
}
