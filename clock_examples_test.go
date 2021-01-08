package clock_test

import (
	"fmt"
	"time"

	"github.com/rzajac/clock"
)

func ExampleSetClock_globally() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	clock.SetClock(start)
	defer clock.ResetClock()

	fmt.Println(clock.Now().Truncate(time.Millisecond).Format(time.RFC3339))

	// Output: 2020-01-01T00:00:00Z
}


func ExampleSetClockTick() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	clock.SetClockTick(start, time.Second)
	defer clock.ResetClock()

	fmt.Println(clock.Now().Format(time.RFC3339))
	fmt.Println(clock.Now().Format(time.RFC3339))
	fmt.Println(clock.Now().Format(time.RFC3339))

	// Output:
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:01Z
	// 2020-01-01T00:00:02Z
}

func ExampleWatch_start() {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	w := clock.New(clock.WatchStart(start))

	fn := func(clk clock.Clock) {
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
	}

	fn(w)

	// Output: 2020-01-01T00:00:00Z
}

func ExampleWatch_tick() {
	fn := func(clk clock.Clock) {
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
	}

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	w := clock.New(clock.WatchTick(start, time.Second))

	fn(w) // Inject clock.

	// Output:
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:01Z
	// 2020-01-01T00:00:02Z
}

func ExampleWatch_static() {
	fn := func(clk clock.Clock) {
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
		fmt.Println(clk.Now().Truncate(time.Second).Format(time.RFC3339))
	}

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	w := clock.New(clock.WatchStatic(start))

	fn(w) // Inject clock.

	// Output:
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:00Z
	// 2020-01-01T00:00:00Z
}
