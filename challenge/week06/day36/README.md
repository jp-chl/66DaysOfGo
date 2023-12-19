# Day 36 of #66DaysOfGo

_Last update:  Aug 24, 2023_.

---

Today, I've continued with the Design Patterns series, with the Observer.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Observer pattern

This pattern maintains a list of observers and notifies them of its state changes. It's useful for event-driven systems and decouples objects. The pattern addresses one-to-many dependencies without tight coupling, allowing automatic updates for multiple dependent objects. It involves defining subject and observer objects, where the subject notifies registered observers of state changes, achieving loose coupling and dynamic addition/removal of observers.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/0/01/W3sDesign_Observer_Design_Pattern_UML.jpg" alt="Observer Pattern UML example" width="750"/>

### Code example

```go
package main

import "fmt"

type Publisher interface {
    addSubscriber(subscriber Subscriber)
    removeSubscriber(subscriberId string)
    broadcast(message string)
}

type Subscriber interface {
    id() string
    receive(message string)
}

type ConcretePublisher struct {
    subscribers map[string]Subscriber
}

func NewConcretePublisher() ConcretePublisher {
    return ConcretePublisher{
        subscribers: make(map[string]Subscriber),
    }
}

func (p ConcretePublisher) addSubscriber(subscriber Subscriber) {
    p.subscribers[subscriber.id()] = subscriber
}

func (p ConcretePublisher) removeSubscriber(subscriberId string) {
    delete(p.subscribers, subscriberId)
}

func (p ConcretePublisher) broadcast(message string) {
    for _, subscriber := range p.subscribers {
        subscriber.receive(message)
    }
}

type ConcreteSubscriber struct {
    subscriberId string
}

func NewConcreteSubscriber(subscriberId string) ConcreteSubscriber {
    return ConcreteSubscriber{
        subscriberId: subscriberId,
    }
}

func (s ConcreteSubscriber) id() string {
    return s.subscriberId
}

func (s ConcreteSubscriber) receive(message string) {
    fmt.Printf("Subscriber %s receiving message: [%s]\n", s.subscriberId, message)
}

func main() {
    publisher := NewConcretePublisher()

    subscriber1 := NewConcreteSubscriber("id1")
    subscriber2 := NewConcreteSubscriber("id2")
    publisher.addSubscriber(subscriber1)
    publisher.addSubscriber(subscriber2)

    publisher.broadcast("hello subscribers")
}

```

```bash
$ go run observer.go
Subscriber id1 receiving message: [hello subscribers]
Subscriber id2 receiving message: [hello subscribers]
```

---

## References

- [https://refactoring.guru/design-patterns/observer](https://refactoring.guru/design-patterns/observer)
