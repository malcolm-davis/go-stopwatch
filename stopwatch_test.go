package stopwatch

import (
	"regexp"
	"testing"
	"time"
)

const (
	expectedMilliseconds = (86400 * 1000)
)

func withNow(fn func() time.Time, callback func()) {
	oldNow := now
	defer func() {
		now = oldNow
	}()

	now = fn
	callback()
}

func withNowOffset(t time.Duration, callback func()) {
	fn := func() time.Time {
		return time.Now().Add(t)
	}

	withNow(fn, callback)
}

func TestStopWatchString(t *testing.T) {
	exp := `^30\.(\d+)ms$`
	rexp := regexp.MustCompile(exp)

	var watch *StopWatch

	withNowOffset(-30*time.Millisecond, func() {
		watch = Start("TestProcess")
	})

	watch.Stop()

	// Not millisecond accurate above, so...
	if !rexp.MatchString(watch.String()) {
		t.Fatalf("expected `%s` to match `%s`", watch, exp)
	}
}
