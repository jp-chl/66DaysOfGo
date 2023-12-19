# Day 40 of #66DaysOfGo

_Last update:  Aug 28, 2023_.

---

Today, I've continued with the Design Patterns series, with the Visitor.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Visitor pattern

This pattern separates algorithms from object structures, enabling the addition of new operations to existing structures without modification. It supports the open/closed principle in object-oriented programming by allowing new functions through a visitor class. This class implements specializations of functions via double dispatch, while languages with sum types can simplify this. It's especially useful for extending APIs externally without source modification.

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/visitor/structure-en-2x.png" alt="Visitor Pattern UML example" width="450"/>

### Code example

```go
package main

import "fmt"

type Element interface {
    Accept(visitor Visitor)
}

type Visitor interface {
    VisitConcreteElement1(element *ConcreteElement1)
    VisitConcreteElement2(element *ConcreteElement2)
    VisitConcreteElement3(element *ConcreteElement3)
}

type ConcreteElement1 struct {
    name string
}
type ConcreteElement2 struct {
    name string
}
type ConcreteElement3 struct {
    name string
}

func (c *ConcreteElement1) Accept(visitor Visitor) {
    visitor.VisitConcreteElement1(c)
}
func (c *ConcreteElement2) Accept(visitor Visitor) {
    visitor.VisitConcreteElement2(c)
}
func (c *ConcreteElement3) Accept(visitor Visitor) {
    visitor.VisitConcreteElement3(c)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElement1(element *ConcreteElement1) {
    fmt.Println("Visited ConcreteElement 1:", element.name)
}
func (v *ConcreteVisitor) VisitConcreteElement2(element *ConcreteElement2) {
    fmt.Println("Visited ConcreteElement 2:", element.name)
}
func (v *ConcreteVisitor) VisitConcreteElement3(element *ConcreteElement3) {
    fmt.Println("Visited ConcreteElement 3:", element.name)
}

func main() {
    elements := []Element{
        &ConcreteElement1{name: "Element 1"},
        &ConcreteElement2{name: "Element 2"},
        &ConcreteElement3{name: "Element 3"},
    }

    visitor := &ConcreteVisitor{}

    for _, elem := range elements {
        elem.Accept(visitor)
    }
}
```

```bash
$ go run visitor.go
Visited ConcreteElement 1: Element 1
Visited ConcreteElement 2: Element 2
Visited ConcreteElement 3: Element 3
```

---

## References

- [https://refactoring.guru/design-patterns/visitor](https://refactoring.guru/design-patterns/visitor)
