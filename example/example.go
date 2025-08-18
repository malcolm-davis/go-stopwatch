package main

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/malcolm-davis/go-stopwatch"
)

func main() {
	// Create a new stopwatch and start it
	watchNL := stopwatch.Start("No log")
	time.Sleep(2 * time.Second)
	watchNL.Stop()
	slog.Info(">>> Watch completed", "duration", watchNL.String())
	slog.Info("-----------------------------------------------------------------")

	// Create a new stopwatch and start it
	watch := stopwatch.Start("my-process", stopwatch.LogStart|stopwatch.LogStop|stopwatch.LogSplit)
	time.Sleep(2 * time.Second)
	watch.Split()
	time.Sleep(1 * time.Second)
	watch.Stop()
	slog.Info(">>> Watch completed", "duration", watch.String())
	slog.Info("-----------------------------------------------------------------")

	// Create a new stopwatch with a custom logger defer stop
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	watch2 := stopwatch.New("Proc2", stopwatch.LogStart|stopwatch.LogStop|stopwatch.LogSplit)
	defer watch2.Stop()
	watch2.Logger = func(format string, v ...interface{}) {
		logger.Info(format, v...)
	}
	watch2.Start()
	time.Sleep(2 * time.Second)
	watch2.Split()
	time.Sleep(1 * time.Second)
	slog.Info("-----------------------------------------------------------------")

	watch3 := stopwatch.Start("My long running method", stopwatch.LogStop)
	time.Sleep(1 * time.Second)
	watch3.Split()
	time.Sleep(1 * time.Second)
	watch3.Stop()
	slog.Info(">>> Watch3 completed", "duration", watch3.String())
	slog.Info("-----------------------------------------------------------------")

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	watch4 := stopwatch.Start("Proc4 global handler", stopwatch.LogStop)
	time.Sleep(1 * time.Second)
	watch4.Split()
	time.Sleep(1 * time.Second)
	watch4.Stop()
	slog.Info(">>> Watch4 completed", "duration", watch4.String())

	slog.Info("-----------------------------------------------------------------")
	watch5 := stopwatch.Start("Mock Method", stopwatch.LogStop)
	time.Sleep(1 * time.Second)
	watch5.StopE(errors.New("Something bad happened"))
	slog.Info(">>> Watch5 completed", "duration", watch5.String())

	callMe()
}

func callMe() (err error) {
	watch := stopwatch.Start("callMe", stopwatch.LogStop)
	defer func() { watch.StopE(err) }()
	time.Sleep(1 * time.Second)
	return errors.New("An error occurred in callMe")
	//return nil
}
