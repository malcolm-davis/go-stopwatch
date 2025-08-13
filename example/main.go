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
	defer watch2.Stop()
	watch2.Logger = func(format string, v ...interface{}) {
		logger.Info(format, v...)
	}

	watch2.Start()
	time.Sleep(2 * time.Second)
	watch2.Split()
	time.Sleep(1 * time.Second)

	watch.Stop()
}
