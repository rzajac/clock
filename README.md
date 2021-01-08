[![Go Report Card](https://goreportcard.com/badge/github.com/rzajac/clock)](https://goreportcard.com/report/github.com/rzajac/clock)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rzajac/clock)

# clock

Testing code which uses `time.Now()` is hard and in many instances impossible.
Package `clock` is an attempt to simplify this kind of tests by providing
methods and interfaces allowing easy way to inject different "clock"
implementations.

Besides `Clock` interface package also provides package level `clock.Now()`
method which behaviour might be changed in tests. See examples for use cases.

# Installation

```
go get github.com/rzajac/clock
```

# Usage

## Inject clock

```
func NewStruct(clk clock.Clock) {
	// ...
}

start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
watch := clock.New(clock.WatchTick(start, time.Second))

s := NewStruct(watch)
```

## Set clock globally

```
now := clock.Now().Format(time.RFC3339)
fmt.Println(now) // 2021-01-08T15:02:53+01:00

clock.SetClock(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
defer clock.ResetClock() // Make sure you reset the clock!

now = clock.Now().Truncate(time.Millisecond).Format(time.RFC3339)
fmt.Println(now) // 2020-01-01T00:00:00Z
```

## Make calls to `clock.Now()` predictable

```
start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
clock.SetClockTick(start, time.Second)
defer clock.ResetClock()

fmt.Println(clock.Now().Format(time.RFC3339)) // 2020-01-01T00:00:00Z
fmt.Println(clock.Now().Format(time.RFC3339)) // 2020-01-01T00:00:01Z
fmt.Println(clock.Now().Format(time.RFC3339)) // 2020-01-01T00:00:02Z
```

## Benchmarks

```
BenchmarkClockNow-12   27557296    42.5 ns/op    0 B/op    0 allocs/op
BenchmarkTimeNow-12    28931214    41.4 ns/op    0 B/op    0 allocs/op
```

## License

Apache License, Version 2.0
