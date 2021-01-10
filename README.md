[![Go Report Card](https://goreportcard.com/badge/github.com/rzajac/clock)](https://goreportcard.com/report/github.com/rzajac/clock)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rzajac/clock)

# clock

Testing code which uses `time.Now()` is hard and in many instances impossible.
Package `clock` is an attempt to simplify this kind of tests by providing
package level clock and few test clock implementations.

# Installation

```
go get github.com/rzajac/clock
```

# Usage

## Inject clock implementation.

Production code:

```
func NewStruct(clk clock.Clock) {
	// ...
	now := clk()
	// ...
}

s := NewStruct(time.Now)

// or

s := NewStruct(clock.Now) // Works exaclty the same way as time.Now().
```

Tests:

```
tim := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
s := NewStruct(clock.Fixed(start))
```

## Set clock globally

In case your code already uses `time.Now()` and you are not able to use
injection method, use package level clock and override its implementation in
tests.

Production code:

```
func doStuff() {
    // ...
    doOtherStuff(clock.Now())
    // ...
}
```

Tests:

```
tim := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
clock.SetClock(clock.Fixed(tim))
defer clock.SetDefault() // Make sure you set package level clock to default!

// Perform your tests.
```

## Useful clock implementations.

Package `clock` provides few useful "clock" you can use in your tests.

- `StartingAt(tim time.Time)` - sets current time.
- `Fixed(tim time.Time)` - clock always returning the same time.
    ```
    tim := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
    clk := clock.Fixed(tim)
    
    fmt.Println(clk().Format(time.RFC3339)) // 2020-01-01T00:00:00Z
    fmt.Println(clk().Format(time.RFC3339)) // 2020-01-01T00:00:00Z
    fmt.Println(clk().Format(time.RFC3339)) // 2020-01-01T00:00:00Z
    ```  
- `Deterministic(start time.Time, tick time.Duration)` - clock which advances
  given start time by tick every time you call it (no matter now fast or slow
  you do it).
    ```
    start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
    clk := clock.Deterministic(start, time.Second)

    fmt.Println(clk().Format(time.RFC3339)) // 2020-01-01T00:00:00Z
    fmt.Println(clk().Format(time.RFC3339)) // 2020-01-01T00:00:01Z
    fmt.Println(clk().Format(time.RFC3339)) // 2020-01-01T00:00:02Z
    ```

## Benchmarks

```
BenchmarkClockNow-12    27322105    42.9 ns/op    0 B/op    0 allocs/op
BenchmarkTimeNow-12     28653739    42.0 ns/op    0 B/op    0 allocs/op
```

## License

Apache License, Version 2.0
