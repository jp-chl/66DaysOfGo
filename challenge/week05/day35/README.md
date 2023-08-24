# Day 34 of #66DaysOfGo

_Last update:  Aug 23, 2023_.

---

Today, I've continued with the Design Patterns series, with the Memento.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Memento pattern

This pattern exposes an object's internal state and is used for tasks like undoing changes, versioning, and custom serialization. It involves three objects: originator, caretaker, and memento. The originator has an internal state, the caretaker performs actions while preserving the ability to undo, and the memento stores the state. The pattern ensures encapsulation and allows saving/restoring state without violating it  

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/3/38/W3sDesign_Memento_Design_Pattern_UML.jpg" alt="Memento Pattern UML example" width="750"/>

### Code example

```go
package main

import (
    "fmt"
)

type Memento struct {
    state string
}

type Originator struct {
    state string
}

func (o *Originator) setState(state string) {
    o.state = state
}

func (o *Originator) getState() string {
    return o.state
}

func (o *Originator) createMemento() *Memento {
    return &Memento{state: o.state}
}

func (o *Originator) restoreMemento(m *Memento) {
    o.state = m.state
}

type Caretaker struct {
    mementos []*Memento
}

func (c *Caretaker) save(m *Memento) {
    c.mementos = append(c.mementos, m)
}

func (c *Caretaker) restore() *Memento {
    if len(c.mementos) == 0 {
        return nil
    }
    index := len(c.mementos) - 1
    memento := c.mementos[index]
    c.mementos = c.mementos[:index]
    return memento
}

func main() {
    caretaker := &Caretaker{}
    originator := &Originator{}

    originator.setState("State 1")
    fmt.Println("Current State:", originator.getState())

    caretaker.save(originator.createMemento())

    originator.setState("State 2")
    fmt.Println("Current State:", originator.getState())

    caretaker.save(originator.createMemento())

    originator.setState("State 3")
    fmt.Println("Current State:", originator.getState())

    memento := caretaker.restore()
    if memento != nil {
        originator.restoreMemento(memento)
        fmt.Println("Restored State:", originator.getState())

        memento = caretaker.restore()
        if memento != nil {
            originator.restoreMemento(memento)
            fmt.Println("Restored State:", originator.getState())
        }
    }

}

```

```bash
$ go run memento.go
Current State: State 1
Current State: State 2
Current State: State 3
Restored State: State 2
Restored State: State 1
```

---

## References

- [https://refactoring.guru/design-patterns/memento](https://refactoring.guru/design-patterns/memento)
