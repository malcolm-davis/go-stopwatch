package stopwatch

import (
	"testing"
)

// const (
// 	expectedMilliseconds = (86400 * 1000)
// )

// func withNow(fn func() time.Time, callback func()) {
// 	oldNow := now
// 	defer func() {
// 		now = oldNow
// 	}()

// 	now = fn
// 	callback()
// }

// func withNowOffset(t time.Duration, callback func()) {
// 	fn := func() time.Time {
// 		return time.Now().Add(t)
// 	}

// 	withNow(fn, callback)
// }

// // need a better test
// func TestStopWatchString(t *testing.T) {
// 	exp := `^30\.(\d+)ms$`
// 	rexp := regexp.MustCompile(exp)

// 	var watch *StopWatch

// 	withNowOffset(-30*time.Millisecond, func() {
// 		watch = Start("TestProcess")
// 	})
// 	watch.Stop()

// 	// Not millisecond accurate above, so...
// 	if !rexp.MatchString(watch.String()) {
// 		t.Fatalf("expected `%s` to match `%s`", watch, exp)
// 	}
// }

func TestZero(t *testing.T) {
	watch := New("TestAction")
	defer watch.Stop()

	// Not millisecond accurate above, so...
	zero := "0m0.00s"
	test := watch.String()
	if zero != test {
		t.Fatalf("expected `%s`, got `%s`", zero, test)
	}

	testSplit := watch.SplitString()
	if zero != testSplit {
		t.Fatalf("expected `%s`, got `%s`", zero, testSplit)
	}
}
