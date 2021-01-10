// Package clock is a thin wrapper around time.Now() to help testing time
// dependent code.
//
// Testing code which uses `time.Now()` is hard and in many instances impossible.
// Package `clock` is an attempt to simplify this kind of tests by providing
// package level clock and few test clock implementations.
package clock

import (
	"sync"
	"time"
)

// Clock is a function signature for returning time.
type Clock func() time.Time

// clock represents package level clock.
var clock Clock = time.Now

// Now returns current clock time as defined by package level clock.
//
// By default it behaves exactly the same way as time.Now but its behaviour
// can be changed (especially in tests) using SetClock method.
func Now() time.Time {
	return clock()
}

// SetClock sets package level clock.
//
// Is reset the package level clock to default implementation use SetDefault
// function. SetClock is not thread safe.
func SetClock(clk Clock) {
	clock = clk
}

// SetDefault sets package level clock to time.Now. SetDefault is not
// thread safe.
func SetDefault() {
	clock = time.Now
}

// StartingAt returns a clock starting at given time.
//
// If StartingAt is used to set the package level clock then the behaviour
// of Now() will be as if you set the system time to tim.
func StartingAt(tim time.Time) Clock {
	now := time.Now()
	guard := sync.Mutex{}
	return func() time.Time {
		guard.Lock()
		defer guard.Unlock()
		return tim.Add(time.Now().Sub(now))
	}
}

// Fixed returns Clock which always returns the tim time.
func Fixed(tim time.Time) Clock {
	return func() time.Time {
		return tim
	}
}

// Deterministic returns a clock which advances given start time by tick every
// time you call it (no matter now fast or slow you do it).
func Deterministic(start time.Time, tick time.Duration) Clock {
	now := start.Add(-tick)
	guard := sync.Mutex{}
	return func() time.Time {
		guard.Lock()
		defer guard.Unlock()
		now = now.Add(tick)
		return now
	}
}
