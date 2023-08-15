# Day 26 of #66DaysOfGo

_Last update:  Aug 14, 2023_.

---

Today, I've continued with the Design Patterns series, with the Composite.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Composite pattern

The composite pattern treats a group of objects as a single instance, useful for creating part-whole hierarchies. It defines a Component interface for both leaf and composite objects, allowing uniform treatment. This pattern simplifies client code, making it easier to implement, change, and reuse. It solves problems of treating part-whole hierarchies uniformly and representing them as tree structures. The composite pattern is applicable when treating individual and grouped objects similarly, reducing complexity in handling multiple objects with nearly identical code.

### UML diagram

<img src="https://i2.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Composite-Design-Pattern-1.jpg?w=537&ssl=1" alt="Composite Pattern UML example" width="450"/>

### Code example

```go

package main

import "fmt"

const CALCULATING_PRICE_MSG_FORMAT = "Calculating price for '%s'\n"
const ADDING_PRICE_MSG_FORMAT = "\tAdding '%.2f'...\n"

type CarComponent interface {
    getPrice() float64
}

type Car struct {
    components []CarComponent
    name       string
}

func (c *Car) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, c.name)
    var totalPrice float64 = 0.0
    for _, carComponent := range c.components {
        totalPrice += carComponent.getPrice()
    }
    return totalPrice
}
func (c *Car) addComponent(cc CarComponent) {
    c.components = append(c.components, cc)
}

type CarInterior struct {
    components []CarComponent
    name       string
}
type CarSeat struct {
    name  string
    price float64
}
type CarShift struct {
    name  string
    price float64
}
type CarEngine struct {
    name  string
    price float64
}

func (ci *CarInterior) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, ci.name)
    var totalPrice float64 = 0.0
    for _, carComponent := range ci.components {
        totalPrice += carComponent.getPrice()
    }
    return totalPrice
}
func (ci *CarInterior) addComponent(cc CarComponent) {
    ci.components = append(ci.components, cc)
}
func (cs *CarSeat) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, cs.name)
    fmt.Printf(ADDING_PRICE_MSG_FORMAT, cs.price)
    return cs.price
}
func (cs *CarShift) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, cs.name)
    fmt.Printf(ADDING_PRICE_MSG_FORMAT, cs.price)
    return cs.price
}
func (ce *CarEngine) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, ce.name)
    fmt.Printf(ADDING_PRICE_MSG_FORMAT, ce.price)
    return ce.price
}

type CarExterior struct {
    components []CarComponent
    name       string
}
type CarDoor struct {
    name  string
    price float64
}
type CarWindow struct {
    name  string
    price float64
}
type CarWheel struct {
    name  string
    price float64
}

func (ce *CarExterior) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, ce.name)
    var totalPrice float64 = 0.0
    for _, carComponent := range ce.components {
        totalPrice += carComponent.getPrice()
    }
    return totalPrice
}
func (ce *CarExterior) addComponent(cc CarComponent) {
    ce.components = append(ce.components, cc)
}
func (cd *CarDoor) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, cd.name)
    fmt.Printf(ADDING_PRICE_MSG_FORMAT, cd.price)
    return cd.price
}
func (cw *CarWindow) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, cw.name)
    fmt.Printf(ADDING_PRICE_MSG_FORMAT, cw.price)
    return cw.price
}
func (cw *CarWheel) getPrice() float64 {
    fmt.Printf(CALCULATING_PRICE_MSG_FORMAT, cw.name)
    fmt.Printf(ADDING_PRICE_MSG_FORMAT, cw.price)
    return cw.price
}

func main() {
    carDoor1 := &CarDoor{name: "door1", price: 10.0}
    carDoor2 := &CarDoor{name: "door2", price: 10.0}
    carDoor3 := &CarDoor{name: "door3", price: 10.0}
    carDoor4 := &CarDoor{name: "door4", price: 10.0}
    carWindow1 := &CarWindow{name: "window1", price: 5.0}
    carWindow2 := &CarWindow{name: "window2", price: 5.0}
    carWindow3 := &CarWindow{name: "window3", price: 5.0}
    carWindow4 := &CarWindow{name: "window4", price: 5.0}
    carWheel1 := &CarWheel{name: "wheel1", price: 2.0}
    carWheel2 := &CarWheel{name: "wheel2", price: 2.0}
    carWheel3 := &CarWheel{name: "wheel3", price: 2.0}
    carWheel4 := &CarWheel{name: "wheel4", price: 2.0}
    carExterior := &CarExterior{name: "Car Exterior"}
    carExterior.addComponent(carDoor1)
    carExterior.addComponent(carDoor2)
    carExterior.addComponent(carDoor3)
    carExterior.addComponent(carDoor4)
    carExterior.addComponent(carWindow1)
    carExterior.addComponent(carWindow2)
    carExterior.addComponent(carWindow3)
    carExterior.addComponent(carWindow4)
    carExterior.addComponent(carWheel1)
    carExterior.addComponent(carWheel2)
    carExterior.addComponent(carWheel3)
    carExterior.addComponent(carWheel4)

    carSeat1 := &CarSeat{name: "seat1", price: 30.0}
    carSeat2 := &CarSeat{name: "seat2", price: 30.0}
    carSeat3 := &CarSeat{name: "seat3", price: 30.0}
    carSeat4 := &CarSeat{name: "seat4", price: 30.0}
    carShift := &CarShift{name: "shift", price: 40.0}
    carEngine := &CarEngine{name: "engine", price: 100.0}
    carInterior := &CarInterior{name: "Car Interior"}
    carInterior.addComponent(carSeat1)
    carInterior.addComponent(carSeat2)
    carInterior.addComponent(carSeat3)
    carInterior.addComponent(carSeat4)
    carInterior.addComponent(carShift)
    carInterior.addComponent(carEngine)

    car := &Car{name: "Car"}
    car.addComponent(carInterior)
    car.addComponent(carExterior)

    totalPrice := car.getPrice()
    fmt.Printf("Total Car price: %.2f\n", totalPrice)
}
```

```bash
$ go run composite.go
Calculating price for 'Car'
Calculating price for 'Car Interior'
Calculating price for 'seat1'
    Adding '30.00'...
Calculating price for 'seat2'
    Adding '30.00'...
Calculating price for 'seat3'
    Adding '30.00'...
Calculating price for 'seat4'
    Adding '30.00'...
Calculating price for 'shift'
    Adding '40.00'...
Calculating price for 'engine'
    Adding '100.00'...
Calculating price for 'Car Exterior'
Calculating price for 'door1'
    Adding '10.00'...
Calculating price for 'door2'
    Adding '10.00'...
Calculating price for 'door3'
    Adding '10.00'...
Calculating price for 'door4'
    Adding '10.00'...
Calculating price for 'window1'
    Adding '5.00'...
Calculating price for 'window2'
    Adding '5.00'...
Calculating price for 'window3'
    Adding '5.00'...
Calculating price for 'window4'
    Adding '5.00'...
Calculating price for 'wheel1'
    Adding '2.00'...
Calculating price for 'wheel2'
    Adding '2.00'...
Calculating price for 'wheel3'
    Adding '2.00'...
Calculating price for 'wheel4'
    Adding '2.00'...
Total Car price: 328.00
```

---

## References

- [https://refactoring.guru/design-patterns/composite](https://refactoring.guru/design-patterns/composite)
