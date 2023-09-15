# Day 54 of #66DaysOfGo

_Last update:  Sep 14, 2023_.

---

Today, I've learned how to run a simple benchmark.

---

> Based on _https://hackernoon.com/how-to-write-benchmarks-in-golang-like-an-expert-0w1834gs_

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- CPUs: 4

## Scope

Compare two versions of Fibonacci (sequential and recursive).

## Setup

go mod init example.com

## Code

```go
// fibonacci.go
func RecursiveFibonacci(n uint) uint {
    if n <= 1 {
        return n
    }
    return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}

func SequentialFibonacci(n uint) uint {
    if n <= 1 {
        return n
    }
    var n1 uint = 1
    var n2 uint = 0
    for i := uint(2); i < n; i++ {
        temp := n1 + n2
        n2 = n1
        n1 = temp
    }
    return n2 + n1
}
```

```go
// fibonacci_test.go
package main

import "testing"

func TestRecursiveFibonacci(t *testing.T) {
    testFibonacci(t)
}

func TestSequentialFibonacci(t *testing.T) {
    testFibonacci(t)
}

func testFibonacci(t *testing.T) {
    data := []struct {
        n    uint
        want uint
    }{
        {0, 0},
        {1, 1},
        {2, 1},
        {3, 2},
        {4, 3},
        {5, 5},
        {6, 8},
        {10, 55},
        {42, 267914296},
    }
    for _, d := range data {
        if got := RecursiveFibonacci(d.n); got != d.want {
            t.Errorf("got: %d, want: %d", got, d.want)
        }
    }
}

func BenchmarkTestRecursiveFibonacci_10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        RecursiveFibonacci(10)
    }
}

func BenchmarkTestRecursiveFibonacci_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        RecursiveFibonacci(20)
    }
}

func BenchmarkTestSequentialFibonacci_10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SequentialFibonacci(10)
    }
}

func BenchmarkTestSequentialFibonacci_20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SequentialFibonacci(20)
    }
}
```

---

## Running the benchmark

Just run `go test -bench=.`. By default it uses all CPUs available.

```bash
# Using 4 CPUs
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: example.com
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkTestRecursiveFibonacci_10-4         3102602           373.8 ns/op
BenchmarkTestRecursiveFibonacci_20-4           24805         46829 ns/op
BenchmarkTestSequentialFibonacci_10-4       140303704            8.415 ns/op
BenchmarkTestSequentialFibonacci_20-4       72031556            16.61 ns/op
PASS
ok      example.com    10.501s
```

```bash
go test -cpu=2 -bench=.      ✔  took 16s   at 22:26:32 
goos: darwin
goarch: amd64
pkg: example.com
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkTestRecursiveFibonacci_10-2         3059385           372.7 ns/op
BenchmarkTestRecursiveFibonacci_20-2           25893         46689 ns/op
BenchmarkTestSequentialFibonacci_10-2       139686907            8.518 ns/op
BenchmarkTestSequentialFibonacci_20-2       69582042            17.19 ns/op
PASS
ok      example.com    10.632s
```

---

## References

- [https://hackernoon.com/how-to-write-benchmarks-in-golang-like-an-expert-0w1834gs](https://hackernoon.com/how-to-write-benchmarks-in-golang-like-an-expert-0w1834gs)
- [https://www.golinuxcloud.com/golang-benchmark/](https://www.golinuxcloud.com/golang-benchmark/)
- [https://medium.com/hyperskill/testing-and-benchmarking-in-go-e33a54b413e](https://medium.com/hyperskill/testing-and-benchmarking-in-go-e33a54b413e)
