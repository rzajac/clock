package clock_test

import (
	"fmt"
	"time"

	"github.com/rzajac/clock"
)

func ExampleSetClock_globally() {
	tim := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	clock.SetClock(clock.StartingAt(tim))
	defer clock.SetDefault()

	fmt.Println(clock.Now().Truncate(time.Millisecond).Format(time.RFC3339))

	// Output:
	// 2020-01-01T00:00:00Z
}

func ExampleSetClock_deterministic() {
	tim := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	clock.SetClock(clock.Deterministic(tim, time.Second))
	defer clock.SetDefault()

	fmt.Println(clock.Now().Format(time.RFC3339))
	fmt.Println(clock.Now().Format(time.RFC3339))
	fmt.Println(clock.Now().Format(time.RFC3339))

	// Output:
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:01Z
	// 2020-01-01T00:00:02Z
}

func ExampleSetClock_fixed() {
	tim := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	clock.SetClock(clock.Fixed(tim))
	defer clock.SetDefault()

	fmt.Println(clock.Now().Format(time.RFC3339))
	fmt.Println(clock.Now().Format(time.RFC3339))
	fmt.Println(clock.Now().Format(time.RFC3339))

	// Output:
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:00Z
}
