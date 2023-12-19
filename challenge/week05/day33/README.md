# Day 33 of #66DaysOfGo

_Last update:  Aug 21, 2023_.

---

Today, I've continued with the Design Patterns series, with the Iterator.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Iterator pattern

This pattern involves using an iterator to navigate a container's elements, decoupling algorithms from containers. This pattern addresses problems like accessing elements without exposing their structure and adding new traversal operations to containers without altering their interface. It introduces a separate iterator object for traversal, allowing clients to access aggregates without knowing their structure. This enables diverse traversal methods using different iterators and independently defining new access operations.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/c/c5/W3sDesign_Iterator_Design_Pattern_UML.jpg" alt="Iterator Pattern UML example" width="750"/>

### Code example

```go
package main

import (
    "fmt"
)

type Iterator interface {
    HasNext() bool
    Next() interface{}
}

type ConcreteIterator struct {
    collection []interface{}
    index      int
}

func NewConcreteIterator(collection []interface{}) *ConcreteIterator {
    return &ConcreteIterator{
        collection: collection,
        index:      0,
    }
}

func (ci *ConcreteIterator) HasNext() bool {
    return ci.index < len(ci.collection)
}

func (ci *ConcreteIterator) Next() interface{} {
    if ci.HasNext() {
        item := ci.collection[ci.index]
        ci.index++
        return item
    }
    return nil
}

type Aggregate interface {
    CreateIterator() Iterator
}

type ConcreteAggregate struct {
    items []interface{}
}

func NewConcreteAggregate() *ConcreteAggregate {
    return &ConcreteAggregate{
        items: []interface{}{},
    }
}

func (ca *ConcreteAggregate) AddItem(item interface{}) {
    ca.items = append(ca.items, item)
}

func (ca *ConcreteAggregate) CreateIterator() Iterator {
    return NewConcreteIterator(ca.items)
}

func main() {
    aggregate := NewConcreteAggregate()

    aggregate.AddItem("Item 1")
    aggregate.AddItem("Item 2")
    aggregate.AddItem("Item 3")

    iterator := aggregate.CreateIterator()

    for iterator.HasNext() {
        item := iterator.Next()
        fmt.Println(item)
    }
}
```

```bash
$ go run iterator.go
Item 1
Item 2
Item 3
```

---

## References

- [https://refactoring.guru/design-patterns/iterator](https://refactoring.guru/design-patterns/iterator)
