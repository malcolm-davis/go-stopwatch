package stopwatch

import (
	"log/slog"
	"time"

	"github.com/malcolm-davis/go-random"
)

// Define enum values using iota and bit shifting
const (
	LogStart = 1 << iota // 0001
	LogSplit             // 0010
	LogStop              // 0100
)

// Create a new watch, and allows for overriding the logging
func New(actionName string, logOn ...int) *StopWatch {
	//loggingOn := LogStart | LogSplit | LogStop
	loggingOn := 0
	if len(logOn) > 0 {
		loggingOn = logOn[0]
	}
	id := random.RandString(12)
	watch := &StopWatch{txnId: id, action: actionName, start: time.Time{}, stop: time.Time{}, logOn: loggingOn}
	return watch
}

// Create a new watch, and start the timing
func Start(actionName string, logOn ...int) *StopWatch {
	// loggingOn := LogStart | LogSplit | LogStop
	loggingOn := 0
	if len(logOn) > 0 {
		loggingOn = logOn[0]
	}

	id := random.RandString(12)
	watch := &StopWatch{txnId: id, action: actionName, start: time.Time{}, stop: time.Time{}, logOn: loggingOn}
	return watch.Start()
}

// StartAt starts a new Watch at the time supplied.
func StartAt(actionName string, t time.Time) *StopWatch {
	return &StopWatch{action: actionName, start: t, stop: time.Time{}}
}

type StopWatch struct {
	txnId, action      string
	start, stop, split time.Time
	logOn              int
	// User defined logger function.
	Logger func(string, ...interface{})
}

var now = func() time.Time {
	return time.Now()
}

func (s *StopWatch) Start() *StopWatch {
	s.start = now()

	if s.logOn&LogStart != 0 {
		s.Info("Start",
			"action", s.action,
			"txnId", s.txnId,
		)
	}
	return s
}

func (s *StopWatch) Stop() *StopWatch {
	s.stop = now()
	if s.logOn&LogStop != 0 {
		s.Info("Finish",
			"action", s.action,
			"txnId", s.txnId,
			"duration", s.duration().String(),
		)
	}
	return s
}

// stop with error
func (s *StopWatch) StopE(err error) *StopWatch {
	s.stop = now()

	msg := "Finish"
	if err != nil {
		msg += " with error"
	}

	if s.logOn&LogStop != 0 {
		s.Info(msg,
			"action", s.action,
			"txnId", s.txnId,
			"duration", s.duration().String(),
		)
	}
	return s
}

func (s *StopWatch) Split() *StopWatch {
	s.split = now()
	if s.logOn&LogSplit != 0 {
		s.Info("Split",
			"action", s.action,
			"txnId", s.txnId,
			"duration", s.splitDuration().String(),
		)
	}
	return s
}

func (s *StopWatch) String() string {
	// if the watch isn't stopped yet...
	if s.stop.IsZero() {
		return "0m0.00s"
	}
	return s.duration().String()
}

func (s *StopWatch) SplitString() string {
	// if the watch isn't split yet...
	if s.split.IsZero() {
		return "0m0.00s"
	}
	return s.splitDuration().String()
}

func (s *StopWatch) splitDuration() time.Duration {
	return s.split.Sub(s.start)
}

func (s *StopWatch) duration() time.Duration {
	return s.stop.Sub(s.start)
}

// Milliseconds returns the elapsed duration in milliseconds.
func (s *StopWatch) Milliseconds() time.Duration {
	return s.duration() / time.Millisecond
}

// Seconds returns the elapsed duration in seconds.
func (s *StopWatch) Seconds() time.Duration {
	return s.duration() / time.Second
}

// Minutes returns the elapsed duration in minutes.
func (s *StopWatch) Minutes() time.Duration {
	return s.duration() / time.Minute
}

// Hours returns the elapsed duration in hours.
func (s *StopWatch) Hours() time.Duration {
	return s.duration() / time.Hour
}

// Days returns the elapsed duration in days.
func (s *StopWatch) Days() time.Duration {
	return s.duration() / (24 * time.Hour)
}

// Info logs message either via defined user logger or via system one if no user logger is defined.
func (s *StopWatch) Info(msg string, args ...interface{}) {
	if s != nil && s.Logger != nil {
		s.Logger(msg, args...)
	} else {
		slog.Info(msg, args...)
	}
}

// Error logs message either via defined user logger or via system one if no user logger is defined.
func (s *StopWatch) Error(msg string, args ...interface{}) {
	if s != nil && s.Logger != nil {
		s.Logger(msg, args...)
	} else {
		slog.Error(msg, args...)
	}
}
