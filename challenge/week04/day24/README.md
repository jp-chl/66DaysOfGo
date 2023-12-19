# Day 24 of #66DaysOfGo

_Last update:  Aug 12, 2023_.

---

Today, I've continued with the Design Patterns series, with the Adapter.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Adapter pattern

An adapter is a special object that converts one object's interface into a format understandable by another object. It wraps an object to simplify conversion, without the wrapped object being aware. For instance, it can convert data from metric to imperial units. Adapters facilitate data conversion and enable collaboration between objects with different interfaces by receiving calls from one object, passing them to another in the expected format, and even creating two-way conversions.

### UML diagram

<img src="https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Adapter-Design-Pattern-1.jpg?w=561&ssl=1" alt="Adapter Pattern UML example" width="550"/>

### Code example

```go
package main

import (
    "fmt"
    "math/rand"
)

type WeatherProvider interface {
    GetCelsius() int
}

type Provider1 struct{}
type Provider2 struct{}

func (p Provider1) GetCelsius() int {
    return rand.Intn(45)
}

func (p Provider2) GetFahrenheit() int {
    return rand.Intn(115)
}

type AdapterProvider2 struct {
    provider Provider2
}

func (p AdapterProvider2) GetCelsius() int {
    return int((float64)(p.provider.GetFahrenheit()-32) / 1.8)
}

func main() {
    weatherInCity1 := Provider1{}
    fmt.Print("City 1. ")
    demoTemperature(weatherInCity1)

    weatherInCity2 := AdapterProvider2{
        provider: Provider2{},
    }
    fmt.Print("City 2. ")
    demoTemperature(weatherInCity2)
}

func demoTemperature(provider WeatherProvider) {
    fmt.Printf("Weather in celsius: %v\n", provider.GetCelsius())
}
```

```bash
$ go run adapter.go
City 1. Weather in celsius: 10
City 2. Weather in celsius: 29
```

---

## References

- [https://refactoring.guru/design-patterns/adapter](https://refactoring.guru/design-patterns/adapter)
