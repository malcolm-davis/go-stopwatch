# stopwatch

A  package for timing code. The intention is to provide a simple, light-weight library for benchmarking.

Designed to measure the amount of time that elapses between its activation and deactivation.

Goals over other stopwatch solutions
* Provide splits
* Provide standard output and logging using slog
* Allow for logging override

## Install
```
go get github.com/malcolm-davis/go-stopwatch
```

## Examples

### Basic

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
    watch.Stop()
}
```

**Output**
```
2025/08/13 21:28:48 INFO TimerStart txnId=J84QZ923K4VQ process=MyProcess
2025/08/13 21:28:50 INFO TimerSplit txnId=J84QZ923K4VQ process=MyProcess duration=2.0013811s
2025/08/13 21:28:51 INFO TimerStop txnId=J84QZ923K4VQ process=MyProcess duration=3.0025538s
```


### Stopwatch with a custom logger

```go
package main

import (
    "log/slog"
    "os"
    "time"

    "github.com/malcolm-davis/go-stopwatch"
)

func main() {

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
}
```

**Output**
```
{"time":"2025-08-13T21:28:51.4226696-05:00","level":"INFO","msg":"TimerStart","txnId":"77TD8CVEV1I0","process":"Proc2"}
{"time":"2025-08-13T21:28:53.4241304-05:00","level":"INFO","msg":"TimerSplit","txnId":"77TD8CVEV1I0","process":"Proc2","duration":"2.0014608s"}
{"time":"2025-08-13T21:28:54.4244162-05:00","level":"INFO","msg":"TimerStop","txnId":"77TD8CVEV1I0","process":"Proc2","duration":"3.0017466s"}
```
