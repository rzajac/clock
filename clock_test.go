package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rzajac/clock"
)

var past = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
var future = time.Now().Add(time.Hour).Truncate(time.Second)

func Test_Watch_WatchStart(t *testing.T) {
	// --- Given ---
	w := clock.New(clock.WatchStart(past))

	// --- When ---
	tim0 := w.Now()

	// --- Then ---
	assert.Exactly(t, past, tim0.Truncate(time.Second))
}

func Test_Watch_WatchTick(t *testing.T) {
	// --- Given ---
	w := clock.New(clock.WatchTick(past, time.Second))

	// --- When ---
	tim0 := w.Now()
	tim1 := w.Now()
	tim2 := w.Now()

	// --- Then ---
	assert.Exactly(t, past.Add(0*time.Second), tim0)
	assert.Exactly(t, past.Add(1*time.Second), tim1)
	assert.Exactly(t, past.Add(2*time.Second), tim2)
}

func Test_Watch_WatchStatic(t *testing.T) {
	// --- Given ---
	w := clock.New(clock.WatchStatic(past))

	// --- When ---
	tim0 := w.Now()
	tim1 := w.Now()
	tim2 := w.Now()

	// --- Then ---
	assert.Exactly(t, past, tim0)
	assert.Exactly(t, past, tim1)
	assert.Exactly(t, past, tim2)
}

func Test_SetClock_past(t *testing.T) {
	// --- Given ---
	clock.SetClock(past)
	defer clock.ResetClock()

	// --- When ---
	time.Sleep(time.Second)
	now := clock.Now()

	// --- Then ---
	assert.True(t, past.Add(time.Second).Equal(now.Truncate(time.Second)))
}

func Test_SetClock_future(t *testing.T) {
	// --- Given ---
	clock.SetClock(future)
	defer clock.ResetClock()

	// --- When ---
	time.Sleep(time.Second)
	now := clock.Now()

	// --- Then ---
	assert.True(t, future.Add(time.Second).Equal(now.Truncate(time.Second)))
}

func Test_SetClockTick(t *testing.T) {
	// --- Given ---
	clock.SetClockTick(past, time.Second)
	defer clock.ResetClock()

	// --- When ---
	tim0 := clock.Now()
	tim1 := clock.Now()
	tim2 := clock.Now()

	// --- Then ---
	assert.Exactly(t, past.Add(0*time.Second), tim0)
	assert.Exactly(t, past.Add(1*time.Second), tim1)
	assert.Exactly(t, past.Add(2*time.Second), tim2)
}

func Test_SetClockStatic(t *testing.T) {
	// --- Given ---
	clock.SetClockStatic(past)
	defer clock.ResetClock()

	// --- When ---
	tim0 := clock.Now()
	tim1 := clock.Now()
	tim2 := clock.Now()

	// --- Then ---
	assert.Exactly(t, past, tim0)
	assert.Exactly(t, past, tim1)
	assert.Exactly(t, past, tim2)
}
