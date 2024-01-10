# Go Logger interface with implementations

This library provides a common interface for logging in Go. It also provides an implementation of the interface for [slog](https://pkg.go.dev/log/slog),
which is part of the standard library.

Coming soon in independent repositories: Implementations for two other popular log libraries: [Logrus](https://github.com/sirupsen/logrus)
and [zap](https://github.com/uber-go/zap).

## Why is this interface useful?

When we create libraries in general we shouldn't be logging but at times we do have to log, debug what the library is doing or trace the log.

We cannot implement a library with one log library and expect other applications to use the same log library. Here is where this interface comes in.
It allows others to change the log library at any time without changing the code.

## Installation

To install `go-logger`, use the following command:

	go get -u github.com/secondtruth/go-logger

## Quick Start

### Example for [slog](https://pkg.go.dev/log/slog)

```go
package main

import (
	"os"
	"log/slog"

	"github.com/secondtruth/go-logger/logger"
)

func main() {
	slogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLog := slog.New(slogHandler)
	log, _ := logger.NewSlogLogger(slogLog)
	
	log.WithFields(logger.Fields{
		"foo": "bar",
	}).Info("message")
}
```
