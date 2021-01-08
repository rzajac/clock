// Package clock is a thin wrapper around time.Now() to help testing time
// dependent code.
//
// Testing code which uses time.Now() is hard and in many instances impossible.
// Package clock is an attempt to simplify this kind of tests by providing
// methods and interfaces allowing easy way to inject different "clock"
// implementations.
package clock

import (
	"time"
)

// Clock is the interface that wraps the Now method.
//
// Now returns the current local time.
type Clock interface {
	Now() time.Time
}

// clock returns current local time. By default it's set to time.Now.
var clock = time.Now

// Now returns current clock time. By default it behaves exactly the same way
// as time.Now but its behaviour in tests might be changed using: SetClock,
// SetClockTick, SetClockStatic methods.
func Now() time.Time {
	return clock()
}

// SetClock sets package level clock to start time. The behaviour of Now()
// after calling this function is as if you set system time to start.
func SetClock(start time.Time) {
	now := time.Now()
	clock = func() time.Time {
		return start.Add(time.Now().Sub(now))
	}
}

// SetClockTick sets package level clock to start time and makes all future
// calls to Now() return times incremented by tick (no matter now fast or slow
// you call it).
func SetClockTick(start time.Time, tick time.Duration) {
	now := start
	clock = func() time.Time {
		ret := now
		now = now.Add(tick)
		return ret
	}
}

// SetClockStatic sets package level clock to start time and always return the
// same time from Now() function.
func SetClockStatic(start time.Time) {
	clock = func() time.Time {
		return start
	}
}

// ResetClock resets clock to return the same time as time.Now function.
func ResetClock() {
	clock = time.Now
}

// WatchTick is a Watch configuration method advancing start time by one second
// every time Now() is called (no matter now fast or slow you call it).
func WatchTick(start time.Time, tick time.Duration) func(*Watch) {
	now := start
	return func(w *Watch) {
		w.clock = func() time.Time {
			ret := now
			now = now.Add(tick)
			return ret
		}
	}
}

// WatchStart is a Watch configuration method setting clock to start time. This
// option makes New() method of the Watch instance to behave as if you set
// system time to start.
func WatchStart(start time.Time) func(*Watch) {
	return func(w *Watch) {
		now := time.Now()
		w.clock = func() time.Time {
			return start.Add(time.Now().Sub(now))
		}
	}
}

// WatchStatic is a Watch configuration method setting clock to start t. This
// option makes New() method of the Watch instance to always return the same
// time.
func WatchStatic(t time.Time) func(*Watch) {
	return func(w *Watch) {
		w.clock = func() time.Time {
			return t
		}
	}
}

// New returns new Watch instance.
func New(opts ...func(*Watch)) Clock {
	w := &Watch{time.Now}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

// Watch is an implementation of the Clock interface.
type Watch struct {
	clock func() time.Time
}

// Now returns current clock time.
func (w *Watch) Now() time.Time {
	return w.clock()
}
