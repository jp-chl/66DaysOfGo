# Day 26 of #66DaysOfGo

_Last update:  Aug 14, 2023_.

---

Today, I've continued with the Design Patterns series, with the Decorator.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Decorator pattern

This pattern offers solutions for flexible and reusable object-oriented software. It addresses adding responsibilities dynamically to objects, providing an alternative to subclassing for extending functionality. By using Decorator objects that transparently implement the interface of the decorated object, it enables dynamic runtime extension of an object's functionality.

### UML diagram

<img src="https://i1.wp.com/golangbyexample.com/wp-content/uploads/2021/04/Decorator-Design-Patter-min.jpg?w=781&ssl=1" alt="Decorator Pattern UML example" width="450"/>

### Code example

```go
// decorator-wrapper.go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func slowFunc() {
    fmt.Println("Function that takes time...")
    randomDuration := time.Duration(1+rand.Intn(3)) * time.Second
    time.Sleep(randomDuration)
}

func logFuncDuration(theFunc func()) {
    defer func(t time.Time) {
        fmt.Printf("[LOG] Function took %v\n", time.Since(t))
    }(time.Now())
    theFunc()
}

func main() {
    logFuncDuration(slowFunc)
}
```

```bash
$ go run decorator-wrapper.go
Function that takes time...
[LOG] Function took 3.001354329s
```

---

## References

- [https://refactoring.guru/design-patterns/composite](https://refactoring.guru/design-patterns/composite)
- [https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/](https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/)
- [https://medium.com/swlh/go-decorator-pattern-2379974077b1](https://medium.com/swlh/go-decorator-pattern-2379974077b1)
- [https://blog.devgenius.io/decorator-pattern-in-go-e3760957fd3a](https://blog.devgenius.io/decorator-pattern-in-go-e3760957fd3a)
- [https://github.com/alex-leonhardt/go-decorator-pattern](https://github.com/alex-leonhardt/go-decorator-pattern)
- [https://golangbyexample.com/decorator-pattern-golang/](https://golangbyexample.com/decorator-pattern-golang/)
