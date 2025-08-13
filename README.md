# stopwatch

A  package for timing code. The intention is to provide a simple,
light-weight library for benchmarking specific bits of code.

Designed to measure the amount of time that elapses between its activation and deactivation.

Goals over other stopwatch solutions
* Provide splits
* Provide standard output and logging using slog

## Install
```
go get github.com/malcolm-davis/go-stopwatch
```

## Example

```go
package main

import (
    "log/slog"
    "os"
    "time"

    "github.com/malcolm-davis/go-stopwatch"
)

func main() {
    // Create a new stopwatch and start it
    watch := stopwatch.Start("MyProcess")
    time.Sleep(2 * time.Second)
    watch.Split()
    time.Sleep(1 * time.Second)

    // Create a new stopwatch with a custom logger
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    watch2 := stopwatch.New("Proc2")

    // defer stop to the end of the method
    defer watch2.Stop()
    watch2.Logger = func(format string, v ...interface{}) {
        logger.Info(format, v...)
    }

    watch2.Start()
    time.Sleep(2 * time.Second)
    watch2.Split()
    time.Sleep(1 * time.Second)

    watch.Stop()
}```

