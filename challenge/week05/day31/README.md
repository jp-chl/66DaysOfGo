# Day 31 of #66DaysOfGo

_Last update:  Aug 19, 2023_.

---

Today, I've continued with the Design Patterns series, with the Chain of Responsibility.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Chain of Responsibility pattern

This pattern is a behavioral design pattern. It involves a chain of processing objects, each capable of handling certain types of command objects. When a command is received, the processing objects are traversed until one of them handles the request. This pattern supports loose coupling and is similar in structure to the decorator pattern, differing in that only one class in the chain handles the request. It addresses problems like avoiding tight coupling between sender and receiver and enabling multiple receivers to handle requests. It defines a chain of receiver objects, allowing flexible forwarding of requests without sender-receiver coupling.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/6/6a/W3sDesign_Chain_of_Responsibility_Design_Pattern_UML.jpg" alt="Chain of Responsibility Pattern UML example" width="750"/>

### Code example

```go
package main

import (
    "fmt"
)

type Order struct {
    Purchased bool
    Cooked    bool
    Delivered bool
}

type OrderHandler interface {
    SetNext(OrderHandler) OrderHandler
    Handle(*Order) bool
}

type OrderProcessor struct {
    handler OrderHandler
}

func (o *OrderProcessor) SetHandler(handler OrderHandler) {
    o.handler = handler
}

func (o *OrderProcessor) Process(order *Order) bool {
    return o.handler.Handle(order)
}

type PaymentProcessor struct {
    next OrderHandler
}

type CookProcessor struct {
    next OrderHandler
}

type DeliverProcessor struct {
    next OrderHandler
}

func (p *PaymentProcessor) SetNext(next OrderHandler) OrderHandler {
    p.next = next
    return next
}

func (p *PaymentProcessor) Handle(order *Order) bool {
    if order.Purchased {
        if p.next != nil {
            return p.next.Handle(order)
        }
        return true
    }
    fmt.Println("Pending payment")
    return false
}

func (w *CookProcessor) SetNext(next OrderHandler) OrderHandler {
    w.next = next
    return next
}

func (w *CookProcessor) Handle(order *Order) bool {
    if order.Cooked {
        if w.next != nil {
            return w.next.Handle(order)
        }
        return true
    }
    fmt.Println("Cooking in progress")
    return false
}

func (s *DeliverProcessor) SetNext(next OrderHandler) OrderHandler {
    s.next = next
    return next
}

func (s *DeliverProcessor) Handle(order *Order) bool {
    if order.Delivered {
        if s.next != nil {
            return s.next.Handle(order)
        }
        return true
    }
    fmt.Println("Delivering in progress")
    return false
}

func main() {
    order := &Order{
        Purchased: true,
        Cooked:    false,
        Delivered: false,
    }

    orderProcessor := &OrderProcessor{}
    paymentProcessor := &PaymentProcessor{}
    cookingHandler := &CookProcessor{}
    deliverHandler := &DeliverProcessor{}

    paymentProcessor.SetNext(cookingHandler)
    cookingHandler.SetNext(deliverHandler)

    orderProcessor.SetHandler(paymentProcessor)

    if orderProcessor.Process(order) {
        fmt.Println("Order delivered")
    }
}
```

```bash
$ go run chain_of_responsibility.go
Cooking in progress
```

---

## References

- [https://refactoring.guru/design-patterns/chain-of-responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility)
