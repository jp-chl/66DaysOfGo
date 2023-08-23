# Day 34 of #66DaysOfGo

_Last update:  Aug 22, 2023_.

---

Today, I've continued with the Design Patterns series, with the Mediator.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Mediator pattern

This pattern encapsulates interactions between objects, reducing complexity as instances communicate through a mediator instead of directly. This pattern addresses tight coupling issues, enabling independent interaction changes and enhancing flexibility, reusability, and testability. By introducing a mediator object to control and coordinate interactions, objects become loosely coupled, leading to improved maintainability and reduced dependency.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/9/92/W3sDesign_Mediator_Design_Pattern_UML.jpg" alt="Mediator Pattern UML example" width="750"/>

### Code example

```go
package main

import "fmt"

type Mediator interface {
    Broadcast(message string, colleague Broker)
}

type Broker interface {
    GetId() string
    Broadcast(message string)
    ReceiveMessage(message, from string)
}

type ConcreteBroker struct {
    id       string
    mediator Mediator
}

func NewConcreteBroker(id string, mediator Mediator) *ConcreteBroker {
    return &ConcreteBroker{
        id:       id,
        mediator: mediator,
    }
}

func (c *ConcreteBroker) GetId() string {
    return c.id
}

func (c *ConcreteBroker) Broadcast(message string) {
    fmt.Printf("%s sends: %s\n", c.id, message)
    c.mediator.Broadcast(message, c)
}

func (c *ConcreteBroker) ReceiveMessage(message, from string) {
    fmt.Printf("%s receives from %s: %s\n", c.id, from, message)
}

type ConcreteMediator struct {
    subscribers map[string]Broker
}

func NewConcreteMediator() *ConcreteMediator {
    return &ConcreteMediator{
        subscribers: make(map[string]Broker),
    }
}

func (m *ConcreteMediator) Subscribe(colleague Broker) {
    m.subscribers[colleague.GetId()] = colleague
}

func (m *ConcreteMediator) Broadcast(message string, colleague Broker) {
    for name, c := range m.subscribers {
        if c.GetId() != colleague.GetId() {
            c.ReceiveMessage(message, name)
        }
    }
}

func main() {
    mediator := NewConcreteMediator()

    broker1 := NewConcreteBroker("Broker1", mediator)
    broker2 := NewConcreteBroker("Broker2", mediator)
    broker3 := NewConcreteBroker("Broker3", mediator)

    mediator.Subscribe(broker1)
    mediator.Subscribe(broker2)
    mediator.Subscribe(broker3)

    broker1.Broadcast("Hello!")
    fmt.Println()
    broker2.Broadcast("Ciao!")
    fmt.Println()
    broker3.Broadcast("Hallo!")
}
```

```bash
$ go run mediator.go
Broker1 sends: Hello!
Broker2 receives from Broker2: Hello!
Broker3 receives from Broker3: Hello!

Broker2 sends: Ciao!
Broker1 receives from Broker1: Ciao!
Broker3 receives from Broker3: Ciao!

Broker3 sends: Hallo!
Broker1 receives from Broker1: Hallo!
Broker2 receives from Broker2: Hallo!
```

---

## References

- [https://refactoring.guru/design-patterns/mediator](https://refactoring.guru/design-patterns/mediator)
