# Day 21 of #66DaysOfGo

_Last update:  Aug 09, 2023_.

---

Today, I've continued with the Design Patterns series, with the Builder.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Builder pattern

The builder pattern is a design pattern that addresses the construction of complex objects, separating the creation from its representation.

It allows flexibility by encapsulating creation and assembly in a separate Builder object, to which a class delegates object creation. Different Builder objects can create various representations of a complex object.

Advantages include varying internal representation and encapsulating construction code, while disadvantages involve the need for distinct ConcreteBuilder for each product type, mutability of Builder classes, and potential complications in dependency injection.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/8/87/W3sDesign_Builder_Design_Pattern_UML.jpg" alt="Builder Pattern UML example" width="900"/>

### Code example

```go
package main

import (
    "fmt"
)

type Car struct {
    brand string
    color string
    year  int
}

type CarBuilder struct {
    car Car
}

func NewCarBuilder() *CarBuilder {
    return &CarBuilder{
        car: Car{},
    }
}

func (cb *CarBuilder) WithBrand(brand string) *CarBuilder {
    cb.car.brand = brand
    return cb
}
func (cb *CarBuilder) WithColor(color string) *CarBuilder {
    cb.car.color = color
    return cb
}
func (cb *CarBuilder) WithYear(year int) *CarBuilder {
    cb.car.year = year
    return cb
}

func (cb *CarBuilder) Build() Car {
    return cb.car
}

func main() {
    bmw := Car{
        brand: "bmw",
        color: "red",
        year:  2000,
    }
    demoCar(bmw)

    porsche := NewCarBuilder().
        WithBrand("porsche").
        WithColor("white").
        // WithYear(2019).
        Build()
    demoCar(porsche)
}

func demoCar(c Car) {
    fmt.Printf("Brand: %s, Color: %s, Year: %d\n", c.brand, c.color, c.year)
}
```

```bash
$ go run builder.go
Brand: bmw, Color: red, Year: 2000
Brand: porsche, Color: white, Year: 0
```

---

## References

- [https://refactoring.guru/design-patterns/builder](https://refactoring.guru/design-patterns/builder)
