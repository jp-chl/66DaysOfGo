# Day 4 of #66DaysOfGo

_Last update:  Jul 16, 2023_.

---

Today, I've continued with the Concurrency series, this time with the pipeline and the fan-in patterns.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- pre-commit: 3.3.3

---

## Pipeline pattern

A series of stages connected by channels, where each step is a group of goroutines.
Each stage, can receive values via _inbound_ channels (read), perform some function on that data and send values via _outbound_ channels (write).

An intermediate stage has inbound and outbound channels. The first stage (aka the _source_ or _producer_) only has an output channel, and the last stage (aka _sink_ or _consumer_) only an input one.

The following example creates a pipeline that put a list of integers into a channel, cube them all, and finally outputs the average of the cube list.

```go
// first stage
func listToChannel(nums ...int) <-chan int {
    out := make(chan int)

    go func() {
        for _, num := range nums {
            out <- num
        }
        close(out)
    }()

    return out
}
```

```go
// second stage
func cube(in <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        for num := range in {
            out <- num * num * num
        }
        close(out)
    }()

    return out
}
```

```go
// final stage
func average(in <-chan int) float64 {
    sum := 0
    elements := 0
    for num := range in {
        sum += num
        elements++
    }
    return (float64)(sum / elements)
}
```

```go
// pipeline
func main() {
    nums := []int{1, 2, 3, 4, 5}

    listAsChannel := listToChannel(nums...)
    cubesChannel := cube(listAsChannel)
    result := average(cubesChannel)

    fmt.Println("Average of the cubes is", result)
}
```

```bash
go run pipelines.go
Average of the cubes is 45
```

## Fan-out, fan-in patterns

Fan-out in Go involves multiple functions reading concurrently from the same channel for parallel processing, until the channel is closed. Fan-in, on the other hand, merges multiple channel inputs into a single channel that's closed when all the input are. Both patterns help in optimizing CPU and I/O usage in a concurrent system.

A fan-in pattern can be implemented by leveraging the (native) `sync` package.

```go
func successor(in <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        for num := range in {
            out <- num + 1
        }
        close(out)
    }()

    return out
}
```

```go
func main() {
    listAsChannel1 := listToChannel(1, 2, 3, 4, 5)
    listAsChannel2 := listToChannel(10, 11, 12, 13, 14)

    c1 := successor(listAsChannel1)
    c2 := successor(listAsChannel2)

    for num := range merge(c1, c2) {
        fmt.Println(num)
    }
}
```

```go
// fan-in
func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }

    wg.Add(len(channels))

    for _, c := range channels {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

Let's dig into the `merge` function:

- `sync.WaitGroup`: a simple struct type that allows to wait for a collection of goroutines to finish executing. The zero value for a WaitGroup is a group with no goroutines.
- `wg.Add(len(channels)`: adds the number of input channels `cs` to the wait group `wg`.
- `wg.Done()`: decreases the WaiGroup counter by one. When each goroutine finishes reading values from its input channel and writes to the `out` channel, then signals that it's done.
- `wg.Wait()`: blocks the goroutine where it's called until the WaitGroup counter is zero, i.e. it waits until all goroutines added with `wg.Add()` have called `wg.Done()`.
- To avoid a `panic` due to sending data to a closed channel, the `close` is done after the `Wait` call.

Notice that the output can vary on each execution, is dependent on the timing of the goroutines, and is not deterministic.

```bash
go run fan_in.go
11
12
2
3
13
4
14
15
5
6
```

```bash
go run fan_in.go
2
3
4
5
11
6
12
13
14
```

## References

- [Go Concurrency Patterns: Pipelines and cancellation (Go Blog)](https://go.dev/blog/pipelines)
