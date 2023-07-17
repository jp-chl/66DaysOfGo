# Day 3 of #66DaysOfGo

_Last update:  Jul 15, 2023_.

---

Today, I've continued with the Concurrency series, this time with some concurrency patterns.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- pre-commit: 3.3.3

---

## Channel capacity

A channel has a buffer capacity. A 0 buffer size is called an "_unbuffered_" channel.

```go
unBufferedChannel1 := make(chan int)
unBufferedChannel2 := make(chan int, 0)
```

By default, a channel has a buffer capacity of 0, i.e. every single send will block until another goroutine receives from the channel.

For instance:

```go
myChannel := make(chan int, 1)
c <- 1 // does not block
c <- 2 // blocks until another goroutine receives from the channel
```

## Select

Golang provides a `select` statement that lets a goroutine wait on multiple communication operations.
A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

For example, the following snippet will print a channel message as soon as it's available.

```go
func main() {
    channel1 := make(chan string)
    channel2 := make(chan string)

    go func() {
        channel1 <- "data for channel 1"
    }()

    go func() {
        channel2 <- "data for channel 2"
    }()

    select {
    case msgFromChannel1 := <-channel1:
        fmt.Println(msgFromChannel1)
    case msgFromChannel2 := <-channel2:
        fmt.Println(msgFromChannel2)
    }
}
```

```bash
while true; do go run select.go; done
data for channel 2
data for channel 2
data for channel 2
data for channel 2
data for channel 2
data for channel 1
data for channel 2
data for channel 2
data for channel 2
data for channel 2
^C
```

## For-Select pattern

This pattern allows to act whenever a channel is ready. If multiple are ready, a random one is picked.

```go
    myChannel1 := make(chan int, 3)
    myChannel2 := make(chan int, 5)

    mySlice := []int{1, 2, 3, 4, 5, 6}

    for _, number := range mySlice {
        select {
        case myChannel1 <- number:
            fmt.Println("Sending to Channel 1")
        case myChannel2 <- number:
            fmt.Println("Sending to Channel 2")
        }
    }

    close(myChannel1)
    close(myChannel2)

    for message := range myChannel2 {
        fmt.Println("Receiving number ", message)
    }

    for message := range myChannel1 {
        fmt.Println("Receiving number ", message)
    }
```

Sample outputs:

```bash
$ go run buffered.go
Sending to Channel 2
Sending to Channel 2
Sending to Channel 1
Sending to Channel 2
Sending to Channel 1
Sending to Channel 2
Receiving number  1
Receiving number  2
Receiving number  4
Receiving number  6
Receiving number  3
Receiving number  5
```

```bash
$ go run buffered.go
Sending to Channel 2
Sending to Channel 1
Sending to Channel 1
Sending to Channel 2
Sending to Channel 2
Sending to Channel 2
Receiving number  1
Receiving number  4
Receiving number  5
Receiving number  6
Receiving number  2
Receiving number  3
```

## "Done-channel" pattern

A simple way to notify a channel to stop processing is via a "done" channel.
A channel can be passed as a parameter and act as a stop trigger in the select statement.

```go
func working(done <-chan bool) {
    for {
        select {
        default:
            fmt.Println("working...")
            time.Sleep(time.Second * 1)
        case <-done:
            fmt.Println("stopping")
            return
        }
    }
}

func main() {
    done := make(chan bool)

    go working(done)

    time.Sleep(time.Second * 3)

    close(done)

    // To print "stopping", you can add a delay after close
    //time.Sleep(time.Millisecond * 5)
}
```

## References

- [Concurrency (Effective Go)](https://go.dev/doc/effective_go#concurrency)
- [Go Concurrency Patterns: Pipelines and cancellation (Go Blog)](https://go.dev/blog/pipelines)
