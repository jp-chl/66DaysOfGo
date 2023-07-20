# Day 7 of #66DaysOfGo

_Last update:  Jul 19, 2023_.

---

Today, I've continued with the Concurrency series, this time understanding good and bad practices about it.

---

## Bad practices

Here there are some common bad practices in Golang:

- **Data Races**: Data races occur when two or more goroutines access the same memory location concurrently and at least one of them is a write. This can lead to unpredictable results because the outcome depends on which goroutine finishes first.
- **Ignoring Goroutine Leaks**: Goroutines are cheap but not free. Each one consumes a certain amount of memory. Failing to properly control goroutines can lead to "_goroutine leaks_", where goroutines are started but never stop. This can consume resources and ultimately cause your program to run out of memory.
- **Incorrect use of the `sync` package**: The sync package in Go provides primitives like `Mutex`, `RWMutex`, `Cond`, `WaitGroup`, `Pool`, `Once`, etc. which are used for synchronization. Misusing or incorrectly using these primitives can lead to issues like deadlocks or data races.
- **Not using context for cancellation**: In Go, context is often used to control the lifecycle of goroutines, particularly for cancellation, timeouts, and passing request-scoped data. Ignoring or improperly using the context package can lead to problems like goroutines that keep running after they are no longer needed.
- **Overusing channels**: Channels in Go provide a way for goroutines to communicate with each other and synchronize their execution. However, overusing channels, or using them where simple locking would suffice, can complicate your code and make it harder to understand.
- **Blocking on Unbuffered Channels**: When sending or receiving on unbuffered channels, goroutines will block if there isn't a corresponding receive or send on the other side. This can lead to deadlocks if not handled correctly.
- **Ignoring Errors**: Go encourages explicit error handling, and this includes errors related to concurrency such as sending to a closed channel. Ignoring these errors can lead to your program behaving unexpectedly or crashing.
- **Assuming Memory Visibility**: When multiple goroutines read and write to the same variables, one must use a synchronization primitive (like mutex or atomic operations) to guarantee that all goroutines see the updates. Failing to do so may lead to some goroutines seeing stale or incorrect data, causing data races.
- **Starvation**: When using mutexes, one should be aware of starvation, where a goroutine can be blocked indefinitely if others are constantly acquiring a lock. Starvation can lead to unresponsive or slowly performing software.
- **Misuse of `sync/atomic` package**: The sync/atomic package provides low-level locking and synchronization primitives. Misuse can lead to subtle and hard-to-find bugs. Atomic operations should only be used when absolutely necessary and by developers who understand their implications fully.
- **Overcomplicating Concurrency Model**: Go’s concurrency model is quite powerful, but it's easy to make it more complex than necessary. One should strive for the simplest model that suits their needs. The more complex your concurrency model is, the more likely it is to have hard-to-find concurrency-related bugs.
- **Improper use of `select`**: The `select` statement lets a goroutine wait on multiple communication operations. Misusing select or failing to understand its subtleties (like the fact that select chooses at random when multiple cases are ready) can lead to bugs or inefficient code.

## Good practices

And here there are some common good practices in Golang:

- **Prefer `sync` primitives**: Go's `sync` package provides several primitives, like `Mutex`, `WaitGroup`, and `Cond`, which are much easier and safer to use than managing synchronization manually with channels. Use these primitives whenever possible.
- **Use `context` for cancellation**: As mentioned before, the context package in Go can be used to control the lifecycle of goroutines, particularly for cancellation. Using context correctly can help prevent goroutine leaks and make your code easier to understand and maintain.
- **Use channels when necessary**: Channels in Go are a powerful way to synchronize and communicate between goroutines. However, they can be overused. Use them when they provide a clear benefit, like when you need to pass ownership of data, or when you need to coordinate tasks across multiple goroutines.
- **Avoid shared state**: Concurrency issues often arise from multiple goroutines accessing and mutating shared state. Whenever possible, avoid shared state, or use synchronization primitives to protect it.
- **Use `select` for multiplexing**: The select statement in Go allows a goroutine to wait on multiple communication operations. It's a powerful tool for managing multiple channels and should be used when you need to multiplex input or output from multiple channels.
- **Keep goroutines small and focused**: Just like functions, goroutines should do one thing and do it well. Keeping your goroutines small and focused will make them easier to understand, test, and debug.
- **Handle errors**: Ignoring errors from goroutines or from channel operations can lead to subtle bugs and unpredictable behavior. Always check and handle errors appropriately.
- **Limit the number of goroutines**: Goroutines are lightweight but not free. Each one consumes some amount of memory and CPU. Creating an unlimited number of goroutines can lead to high memory usage and scheduler contention. Use tools like semaphores (for example, `sync.Semaphore`) to limit the number of concurrent goroutines.
- **Use buffered channels when necessary**: Unbuffered channels cause senders to block until the receiver is ready. This can cause performance issues or even deadlocks. Buffered channels can help prevent this by allowing senders to send values without blocking, up to the capacity of the buffer.
- **Use `sync.Once` for one-time initialization**: If you have some state that should be initialized only once (like a singleton or a lazy-loaded variable), use `sync.Once`. This will ensure that the initialization code is executed only once, even if multiple goroutines attempt to initialize it simultaneously.
- **Test your code with the race detector**: Go provides a race detector tool that can help you catch data races in your code. While it won't catch every possible race condition, it's a valuable tool for finding and fixing these issues.
- **Think about the worst-case scenario**: When designing with concurrency, think about what can go wrong, such as goroutines that never terminate, blocked goroutines, deadlocks, etc. Design your code to handle these scenarios appropriately.

## References

- [https://en.wikipedia.org/wiki/Communicating_sequential_processes](Tony Hoare's CSP paper)
- [https://divan.dev/posts/go_concurrency_visualize/](Visualizing Concurrency in Go)
- [https://github.com/golang-basics/concurrency](Concurrency in Go series)
