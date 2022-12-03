package async

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/platinummonkey/go-utils/atomic"
)

var multiplier = NewAtomic[float64](1.0)

func init() {
	for i, arg := range os.Args[1:] {
		if arg == "asyncWaitMultiplier" {
			newMultiplier, err := strconv.ParseFloat(os.Args[i+1], 32)
			if err == nil {
				SetTimingMultiplier(newMultiplier)
			}
		}
	}
}

// SetTimingMultiplier will set the current timing multiplier.
// Normally this would be set through the above args, however, for
// programmatic access this exists.
func SetTimingMultiplier(m float64) {
	if m > 0.0 {
		multiplier.Store(m)
	}
}

// GetTimingMultiplier will return the current timing multiplier
func GetTimingMultiplier() float64 {
	return multiplier.Load()
}

// ConditionalWait is a test utility to wait for a specific async condition to occur within a time range.
// for example if a go-routine has to do some work or wait for results you can conditional wait on the side-effect
// before trying to proceed with the tests. This will help prevent flaky time.Sleep()s and try to move on as fast as
// possible. It's important to set a decent sleep duration to allow the async go-routine to have some cpu-cycles.
func ConditionalWait(
	within time.Duration,
	maxAttempts int64,
	predicate func(counter int64) bool,
	msg string,
	args ...interface{},
) error {
	conditionMet := false
	within = TimingMultiplier(within)
	sleepDuration := within / time.Duration(maxAttempts)
	for i := int64(0); i < maxAttempts; i++ {
		conditionMet = predicate(i)
		if conditionMet {
			return nil
		}
		time.Sleep(sleepDuration)
		runtime.Gosched()
	}
	if !conditionMet {
		errorMsg := "condition not met by predicate"
		if msg != "" {
			errorMsg += ": " + msg
		}

		if len(args) == 0 {
			return fmt.Errorf(errorMsg)
		}
		return fmt.Errorf(errorMsg, args...)
	}
	return nil
}

// TimingMultiplier adjusts d based on the global multiplier set for the test run.
// In slow environments, you can set multiplier > 1 to make timeouts longer than
// tests indicate.
func TimingMultiplier(d time.Duration) time.Duration {
	return time.Duration(float64(d.Nanoseconds())*GetTimingMultiplier()) * time.Nanosecond
}

// FinishWithin attempts to run f within d, returning an error if f took longer
// than d or if f returned an error.
func FinishWithin(d time.Duration, f func() error) error {
	d = TimingMultiplier(d)
	done := make(chan struct{})
	var err error
	go func() {
		err = f()
		close(done)
	}()

	runtime.Gosched()

	select {
	case <-time.After(d):
		return errors.New("expired")
	case <-done:
		return err
	}
}
