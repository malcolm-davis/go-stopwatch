# stopwatch

A  package for timing code. The intention is to provide a simple, light-weight library for benchmarking.

Designed to measure the amount of time that elapses between its activation and deactivation.

Goals over other stopwatch solutions

* Provide splits
* Provide standard output and logging using slog
* Provide a default trxId for the logging output
* Allow for logging override

## Install
```
go get github.com/malcolm-davis/go-stopwatch
```


## Examples

### Basic - no event logging

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
    watchNL := stopwatch.Start("No log")
    time.Sleep(2 * time.Second)
    watchNL.Split()
    time.Sleep(1 * time.Second)
    watchNL.Stop()
    slog.Info("Watch completed", "duration", watchNL.String())
}
```


**Output**
```
Watch completed duration=3.0004139ss
```

### Basic - all event logging

```go
package main

import (
    "log/slog"
    "os"
    "time"

    "github.com/malcolm-davis/go-stopwatch"
)

func main() {
	watch := stopwatch.Start("my-process", stopwatch.LogStart|stopwatch.LogStop|stopwatch.LogSplit)
	time.Sleep(2 * time.Second)
	watch.Split()
	time.Sleep(1 * time.Second)
	watch.Stop()
	slog.Info("Watch completed", "duration", watch.String())
}
```

**Output**
```
2025/08/18 14:33:58 INFO Start action=my-process txnId=PH0L3L504M5V
2025/08/18 14:34:00 INFO Split action=my-process txnId=PH0L3L504M5V duration=2.000541s
2025/08/18 14:34:01 INFO Finish action=my-process txnId=PH0L3L504M5V duration=3.0013657s
2025/08/18 14:34:01 INFO Watch completed duration=3.0013657s
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
{"time":"2025-08-13T21:28:51.4226696-05:00","level":"INFO","msg":"Start","action":"Proc2","txnId":"77TD8CVEV1I0"}
{"time":"2025-08-13T21:28:53.4241304-05:00","level":"INFO","msg":"Split","action":"Proc2","txnId":"77TD8CVEV1I0","duration":"2.0014608s"}
{"time":"2025-08-13T21:28:54.4244162-05:00","level":"INFO","msg":"Finish","action":"Proc2","txnId":"77TD8CVEV1I0","duration":"3.0017466s"}
```

### Log only stop

```go
package main

import (
    "log/slog"
    "os"
    "time"

    "github.com/malcolm-davis/go-stopwatch"
)

func main() {
	watch3 := stopwatch.Start("My long running method", stopwatch.LogStop)
	time.Sleep(1 * time.Second)
	watch3.Split()
	time.Sleep(1 * time.Second)
	watch3.Stop()
}
```

**Output**
```
Finish action="My long running method" txnId=54DFH8SWNSMD duration=2.0008301s
```



### Log if finish due to error

Sample defer  error condition for logging
```go
func callMe() (err error) {
	watch := stopwatch.Start("callMe", stopwatch.LogStop)
	defer func() { watch.StopE(err) }()
	time.Sleep(1 * time.Second)
	return errors.New("An error occurred in callMe")
	//return nil
}
```

**Output**
```
{"time":"2025-08-18T14:47:29.8251759-05:00","level":"INFO","msg":"Finish with error","action":"callMe","txnId":"9WEX9N7V2RV5","duration":"1.0004543s"}
```
