# Day 2 of #66DaysOfGo

_Last update:  Jul 14, 2023_.

---

Today, I've started a Concurrency series, beginning with the basics.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- pre-commit: 3.3.3

---

## Setup

```bash
brew install pre-commit
```

```bash
pre-commit --version
3.3.3
```

Note:
You can run pre-commit before a git commit by executing:

```bash
pre-commit run --all-files
trim trailing whitespace.................................................Passed
fix end of files.........................................................Passed
check yaml...............................................................Passed
check for added large files..............................................Passed
go fmt...............................................(no files to check)Skipped
go vet...............................................(no files to check)Skipped
Check files aren't using go's testing package........(no files to check)Skipped
go-unit-tests........................................(no files to check)Skipped
```

---

## Concurrency basics - Goroutines

Golang provides one simple way to handle concurrency called a "goroutine" which is basically a lightweight thread managed by the Go runtime.

For instance, the following program will print the numbers sequentially.

```go
func myFunc(num int) {
	fmt.Println(num)
}

func main() {
	myFunc(1)
	myFunc(2)
	myFunc(3)
	myFunc(4)
	fmt.Println("end...")
}
```

The output is, obviously sequential:

```bash
go run goroutines.go
1
2
3
4
end...
```

To run them concurrently ([not in parallel](https://www.youtube.com/watch?v=f6kdp27TYZs)) just add the `go` keyword before the function call, and it will fork a new child process (aka goroutine)

```go
func main() {
	go myFunc(1) // each go call will spawn an independent child process
	go myFunc(2)
	go myFunc(3)
	go myFunc(4)

	fmt.Println("end...")
}
```

If you run it several times, normally it'll print `end...` without waiting for `myFunc` function, and sometimes it would print a number.

```bash
go run goroutines.go
end...
```

```bash
go run goroutines.go
end...
```

```bash
go run goroutines.go
3
end...
```

```bash
go run goroutines.go
end...
```

Let's add a wait (`time.Sleep`):

```go
import "time"

// ...

func main() {
	go myFunc(1)
	go myFunc(2)
	go myFunc(3)
	go myFunc(4)

	time.Sleep(time.Second * 1)

	fmt.Println("end...")
}
```

If you try executing this program repeatedly, it'll print all the numbers, but not necessarily in order:

```bash
go run goroutines.go
2
4
3
1
end...
```

```bash
go run goroutines.go
4
1
2
3
end...
```

```bash
go run goroutines.go
4
2
3
1
end...
```

However, with this code, the main process does not sync with each goroutine, and even more each goroutine doesn't coordinate with each other, unless we implement it (for example with [channels](https://go.dev/tour/concurrency/2)).

---

## Introduction to channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

Creating a channel can be very simple. Just call make with the keyword `chan` defining the channel data type.

For instance:

```go
func main() {
	myChannel := make(chan string)

	// anonymous function
	go func() {
		myChannel <- "some data"
	}()

	message := <-myChannel

	fmt.Println(message)
}
```

The `<-` allows to send/receive data from/to the channel.

The output will be:

```bash
go run channels.go
some data
```

---

## References

- [Golang Concurrency Tour](https://go.dev/tour/concurrency/1)
- [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Channels (Go by Example)](https://gobyexample.com/channels)
