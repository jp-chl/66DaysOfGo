# Day 22 of #66DaysOfGo

_Last update:  Aug 10, 2023_.

---

Today, I've continued with the Design Patterns series, with the Singleton.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Singleton pattern

The Singleton is a design pattern ensuring a class has only one instance and provides a global access point to it. This pattern ensures a single instance, easy access to it, and controlled instantiation. The name derives from a [mathematical concept](https://en.wikipedia.org/wiki/Singleton_(mathematics)).

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/singleton/structure-en-2x.png" alt="Singleton Pattern UML example" width="350"/>

### Code example

```go
package main

import (
    "fmt"
    "sync"
)

type MyCustomData struct {
    data string
}

type SingletonProvider struct {
    once              sync.Once
    singletonInstance *MyCustomData
}

func (sp *SingletonProvider) GetInstance() *MyCustomData {
    sp.once.Do(func() {
        sp.singletonInstance = &MyCustomData{
            data: "test data",
        }
    })
    return sp.singletonInstance
}

func main() {
    provider := &SingletonProvider{}

    // several goroutines
    for i := 0; i < 10; i++ {
        go func(index int) {
            uniqueInstance := provider.GetInstance()
            demo(index, uniqueInstance)
        }(i)
    }

    fmt.Scanln() // wait for input before exiting
}

func demo(index int, variable interface{}) {
    fmt.Printf("index = %d, address of (%v) is: %p\n", index, variable, variable)
}
```

```bash
$ go run singleton.go
index = 2, address of (&{test data}) is: 0xc000100000
index = 1, address of (&{test data}) is: 0xc000100000
index = 8, address of (&{test data}) is: 0xc000100000
index = 0, address of (&{test data}) is: 0xc000100000
index = 4, address of (&{test data}) is: 0xc000100000
index = 3, address of (&{test data}) is: 0xc000100000
index = 5, address of (&{test data}) is: 0xc000100000
index = 9, address of (&{test data}) is: 0xc000100000
index = 6, address of (&{test data}) is: 0xc000100000
index = 7, address of (&{test data}) is: 0xc000100000
```

---

## References

- [https://refactoring.guru/design-patterns/singleton](https://refactoring.guru/design-patterns/singleton)
