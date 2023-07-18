# Day 5 of #66DaysOfGo

_Last update:  Jul 17, 2023_.

---

Today, I've continued with the Concurrency series, this time understanding how to avoid deadlocks in pipelines by using a done-channel.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- pre-commit: 3.3.3

---

## Stopping short, unbuffered vs buffered channels

Keep in mind that an unbuffered channel needs a consumer right away.

For instance, the following code will raise: `fatal error: all goroutines are asleep - dealock!`

```go
myChannel := make(chan int)
myChannel <- 2 // This line will raise the error
number := <-myChannel
fmt.Println("number = ", number)
```

The latter will be solved by using a goroutine, i.e.

```go
go func() {
    myChannel <- 2
}()
```

However, with a buffered channel, a goroutine might not even be needed. For example, the following code will run ok:

```go
myChannel := make(chan int, 1) // buffer capacity: 1
myChannel <- 1
number := <-myChannel
fmt.Println("number = ", number)
```

In a pipeline pattern in Golang, stages close their outbound channels once all sends are done, and they keep receiving values from inbound channels until those are closed. However, in real pipelines, stages might not receive all inbound values. This can lead to issues like goroutines blocking indefinitely and causing a resource leak, if not all values are consumed.

One solution is to change outbound channels to buffered ones, allowing them to hold a fixed number of values and causing send operations to complete immediately if there's room. However, this approach can be fragile and dependent on the number of values sent and received. A more effective solution would be to provide a way for downstream stages to signal upstream ones that they will stop accepting input.

## Explicit cancellation ("done-channel" in pipelines)

In Golang, explicit cancellation can be used to tell upstream stages to stop sending values if downstream stages have stopped receiving them.

This can be achieved by creating a '`done`' channel. When a goroutine decides to stop receiving, it sends values into the '`done`' channel, which the sender goroutines check for before performing their send operation.

However, this approach can be complicated if the number of potentially blocked senders is unknown. To resolve this, '`done`' channel can be closed to signal all senders to stop. The close operation is a broadcast signal to the senders. Pipeline stages can then exit as soon as '`done`' is closed, ensuring there are no resource leaks. It's crucial that pipeline construction follows certain guidelines to prevent blocked senders, either by using sufficient buffer space or by explicitly signalling when receivers may abandon the channel.

One approach could be modifying the main function of the [previous pipeline example](../day04/fan_in.go), by adding a `done` buffered channel and then explicitly send it a close signal:

```go
func main() {
    listAsChannel1 := listToChannel(1, 2, 3, 4, 5)
    listAsChannel2 := listToChannel(10, 11, 12, 13, 14)

    done := make(chan bool, 2) // it could also be make(chan struct{}, 2), i.e. data type does not matter much

    c1 := successor(listAsChannel1)
    c2 := successor(listAsChannel2)

    for num := range merge(c1, c2) {
        fmt.Println(num)
    }

    done <- bool
    done <- bool
}
```

Then the merge function would have needed a for-select pattern to identify the `done` signal.

```go
func merge(done <-chan bool, channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        for n := range c {
            select {
            case out <- n:
            case <-done:
        }
        wg.Done()
    }
```

However, the latter approach enforce each downstream receiver to know the number of potentially blocked upstream senders and arrange to signal those senders on early return.

One way to address this type of issue is to close the `done` channel for example by deferring the closure. In this case, we don't even need a buffered channel.

The main function will be like the following. Notice all the pipeline functions have the additional (`done`) parameter.

```go
func main() {
    done := make(chan bool)
    defer close(done)

    listAsChannel1 := listToChannel(done, 1, 2, 3, 4, 5)
    listAsChannel2 := listToChannel(done, 10, 11, 12, 13, 14)

    c1 := successor(done, listAsChannel1)
    c2 := successor(done, listAsChannel2)

    for num := range merge(done, c1, c2) {
        fmt.Println(num)
    }

    // done will be closed by the deferred call
}
```

In the merge function, the `wg.Done` can also be deferred:

```go
func merge(done <-chan bool, channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        defer wg.Done() // <---
        for n := range c {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }
```

The [complete code](./done-pipeline.go) is in the same folder as this readme.

## References

- [Go Concurrency Patterns: Pipelines and cancellation (Go Blog)](https://go.dev/blog/pipelines)
