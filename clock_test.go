package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rzajac/clock"
)

var past = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
var future = time.Now().Add(time.Hour).Truncate(time.Second)

func Test_SetClock_SetDefault(t *testing.T) {
	// --- Given ---
	fixed := clock.Fixed(past)

	// --- When ---
	clock.SetClock(fixed)
	defer clock.SetDefault()

	// --- Then ---
	assert.True(t, clock.Now().Equal(past))
	assert.True(t, clock.Now().Equal(past))

	// Test set default is working properly.
	clock.SetDefault()
	assert.True(t, clock.Now().Sub(time.Now()) < time.Millisecond)
}

func Test_StartingAt_past(t *testing.T) {
	// --- Given ---
	clk := clock.StartingAt(past)

	// --- When ---
	tim1 := clk()
	tim2 := clk()

	// --- Then ---
	assert.True(t, tim1.Before(tim2))
	assert.True(t, tim2.Before(time.Now()))
}

func Test_StartingAt_future(t *testing.T) {
	// --- Given ---
	clk := clock.StartingAt(future)

	// --- When ---
	tim1 := clk()
	tim2 := clk()

	// --- Then ---
	assert.True(t, tim1.Before(tim2))
	assert.True(t, tim2.After(time.Now()))
}

func Test_SetClock_Deterministic(t *testing.T) {
	// --- Given ---
	clock.SetClock(clock.Deterministic(past, time.Second))
	defer clock.SetDefault()

	// --- When ---
	tim0 := clock.Now()
	tim1 := clock.Now()
	tim2 := clock.Now()

	// --- Then ---
	assert.Exactly(t, past.Add(0*time.Second), tim0)
	assert.Exactly(t, past.Add(1*time.Second), tim1)
	assert.Exactly(t, past.Add(2*time.Second), tim2)
}

func Test_SetClock_Fixed(t *testing.T) {
	// --- Given ---
	clock.SetClock(clock.Fixed(past))
	defer clock.SetDefault()

	// --- When ---
	tim0 := clock.Now()
	tim1 := clock.Now()
	tim2 := clock.Now()

	// --- Then ---
	assert.Exactly(t, past, tim0)
	assert.Exactly(t, past, tim1)
	assert.Exactly(t, past, tim2)
}

func BenchmarkClockNow(b *testing.B) {
	var now time.Time
	for i := 0; i < b.N; i++ {
		now = clock.Now()
	}
	_ = now
}

func BenchmarkTimeNow(b *testing.B) {
	var now time.Time
	for i := 0; i < b.N; i++ {
		now = time.Now()
	}
	_ = now
}
