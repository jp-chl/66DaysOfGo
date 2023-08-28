# Day 39 of #66DaysOfGo

_Last update:  Aug 27, 2023_.

---

Today, I've continued with the Design Patterns series, with the Template.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Template pattern

This pattern involves a superclass method defining the structure of an operation using high-level steps, which are implemented by helper methods in the same class. These helpers can be abstract methods, requiring subclass implementation, or hook methods with empty bodies. Subclasses customize the operation by overriding hook methods, allowing for overall structure definition while permitting step refinements. The pattern achieves consistent algorithm execution across subclasses while enabling runtime algorithm selection. This demonstrates inversion of control, and hook methods provide a means for fine-tuning without overriding the template method.

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/template-method/structure-2x.png" alt="Template Pattern UML example" width="450"/>

### Code example

```go
package main

import "fmt"

type EaterType interface {
    pick()
    prepare()
    ingest()
}

type Eater struct {
    eaterType EaterType
}

func (e *Eater) eat() {
    e.eaterType.pick()
    e.eaterType.prepare()
    e.eaterType.ingest()
}
func (e *Eater) setEaterType(eaterType EaterType) {
    e.eaterType = eaterType
}

type MeatEater struct{}
type Vegetarian struct{}
type Vegan struct{}

func (m *MeatEater) pick() {
    fmt.Println("Selecting meat with high protein")
}
func (m *MeatEater) prepare() {
    fmt.Println("Adding vegetables, and roasting the meat")
}
func (m *MeatEater) ingest() {
    fmt.Println("Enjoying meat + vegetables")
}

func (v *Vegetarian) pick() {
    fmt.Println("Selecting beans, eggs and vegetables")
}
func (v *Vegetarian) prepare() {
    fmt.Println("Adding pepper, salt and spicy stuff")
}
func (v *Vegetarian) ingest() {
    fmt.Println("Enjoying a healthy food made of eggs, beans and vegetables")
}

func (v *Vegan) pick() {
    fmt.Println("Selecting only non-animal food")
}
func (v *Vegan) prepare() {
    fmt.Println("Preparing a healthy food")
}
func (v *Vegan) ingest() {
    fmt.Println("Enjoying a super healthy food made of plants")
}

func main() {
    meatEater := &MeatEater{}
    vegetarian := &Vegetarian{}
    vegan := &Vegan{}

    eater := &Eater{
        eaterType: meatEater,
    }
    eater.eat()
    fmt.Println("==================================")

    eater.setEaterType(vegetarian)
    eater.eat()
    fmt.Println("==================================")

    eater.setEaterType(vegan)
    eater.eat()
}
```

```bash
$ go run template.go
Selecting meat with high protein
Adding vegetables, and roasting the meat
Enjoying meat + vegetables
==================================
Selecting beans, eggs and vegetables
Adding pepper, salt and spicy stuff
Enjoying a healthy food made of eggs, beans and vegetables
==================================
Selecting only non-animal food
Preparing a healthy food
Enjoying a super healthy food made of plants
```

---

## References

- [https://refactoring.guru/design-patterns/template-method](https://refactoring.guru/design-patterns/template-method)
