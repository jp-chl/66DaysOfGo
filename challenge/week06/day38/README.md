# Day 38 of #66DaysOfGo

_Last update:  Aug 26, 2023_.

---

Today, I've continued with the Design Patterns series, with the Strategy.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Strategy pattern

This pattern allows runtime selection of algorithms. It separates algorithm implementation from client code, promoting flexibility and reusability. It enables dynamic algorithm choice for tasks like data validation based on factors such as data type or source. Strategies are encapsulated and interchangeable, avoiding duplication, and are stored as references for retrieval, often using function pointers or classes.

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/strategy/structure-2x.png" alt="Strategy Pattern UML example" width="750"/>

### Code example

```go
package main

import "fmt"

type PaymentStrategy interface {
   Pay(amount float64)
}

type DebitCardPayment struct{}
type CreditCardPayment struct{}
type BitcoinPayment struct{}

func (c *DebitCardPayment) Pay(amount float64) {
   fmt.Printf("Paying %.1f using Debit Card\n", amount)
}

func (p *CreditCardPayment) Pay(amount float64) {
   fmt.Printf("Paying %.1f using Credit Cards\n", amount)
}

func (b *BitcoinPayment) Pay(amount float64) {
   fmt.Printf("Paying %.1f using Bitcoin\n", amount)
}

type PaymentProcessor struct {
   paymentMethod PaymentStrategy
}

func (pp *PaymentProcessor) SetPaymentMethod(method PaymentStrategy) {
   pp.paymentMethod = method
}

func (pp *PaymentProcessor) ProcessPayment(amount float64) {
   pp.paymentMethod.Pay(amount)
}

func main() {
   debitCard := &DebitCardPayment{}
   creditCard := &CreditCardPayment{}
   bitcoin := &BitcoinPayment{}

   paymentProcessor := &PaymentProcessor{
      paymentMethod: debitCard,
   }
   paymentProcessor.ProcessPayment(200.5)

   paymentProcessor.SetPaymentMethod(creditCard)
   paymentProcessor.ProcessPayment(30.2)

   paymentProcessor.SetPaymentMethod(bitcoin)
   paymentProcessor.ProcessPayment(123.7)
}
```

```bash
$ go run strategy.go
Paying 200.5 using Debit Card
Paying 30.2 using Credit Cards
Paying 123.7 using Bitcoin
```

---

## References

- [https://refactoring.guru/design-patterns/strategy](https://refactoring.guru/design-patterns/strategy)
