package stopwatch

import (
	"log/slog"
	"time"

	"github.com/malcolm-davis/go-random"
)

// Create a new watch, and allows for overriding the logging
func New(processName string) *StopWatch {
	id := random.RandString(12)
	watch := &StopWatch{txnId: id, processName: processName, start: time.Time{}, stop: time.Time{}}
	return watch
}

// Create a new watch, and start the timing
func Start(processName string) *StopWatch {
	id := random.RandString(12)
	watch := &StopWatch{txnId: id, processName: processName, start: time.Time{}, stop: time.Time{}}
	return watch.Start()
}

// StartAt starts a new Watch at the time supplied.
func StartAt(processName string, t time.Time) *StopWatch {
	return &StopWatch{processName: processName, start: t, stop: time.Time{}}
}

type StopWatch struct {
	txnId, processName string
	start, stop, split time.Time
	// User defined logger function.
	Logger func(string, ...interface{})
}

var now = func() time.Time {
	return time.Now()
}

func (s *StopWatch) Stop() *StopWatch {
	s.stop = now()
	s.Info("TimerStop",
		"txnId", s.txnId,
		"process", s.processName,
		"duration", s.duration().String(),
	)
	return s
}

func (s *StopWatch) Start() *StopWatch {
	s.start = now()
	s.Info("TimerStart",
		"txnId", s.txnId,
		"process", s.processName,
	)
	return s
}

func (s *StopWatch) Split() *StopWatch {
	s.split = now()
	s.Info("TimerSplit",
		"txnId", s.txnId,
		"process", s.processName,
		"duration", s.splitDuration().String(),
	)
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
