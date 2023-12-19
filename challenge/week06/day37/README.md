# Day 37 of #66DaysOfGo

_Last update:  Aug 25, 2023_.

---

Today, I've continued with the Design Patterns series, with the State.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## State pattern

This pattern is used to change an object's behavior based on its internal state, akin to finite-state machines. It's similar to the strategy pattern, allowing behavior changes through method invocations. This pattern enhances runtime behavior changes without using conditionals, aiding maintainability. The pattern suggests using separate state objects or delegating behavior to handle state changes, making classes adaptable and independent of specific behavior implementations.

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/state/structure-en-2x.png" alt="State Pattern UML example" width="750"/>

### Code example

```go
package main

import "fmt"

type State interface {
    InsertCoin(context *VendingMachineContext)
    EjectCoin(context *VendingMachineContext)
    Dispense(context *VendingMachineContext)
}

type IdleState struct{}
type HasCoinState struct{}
type DispensingState struct{}

func (s *IdleState) InsertCoin(context *VendingMachineContext) {
    fmt.Println("Coin inserted. Machine now has a coin.")
    context.setState(&HasCoinState{})
}

func (s *IdleState) EjectCoin(context *VendingMachineContext) {
    fmt.Println("No coin to eject.")
}

func (s *IdleState) Dispense(context *VendingMachineContext) {
    fmt.Println("Payment required before dispensing.")
}

func (s *HasCoinState) InsertCoin(context *VendingMachineContext) {
    fmt.Println("Coin already inserted.")
}

func (s *HasCoinState) EjectCoin(context *VendingMachineContext) {
    fmt.Println("Coin ejected. Machine now idle.")
    context.setState(&IdleState{})
}

func (s *HasCoinState) Dispense(context *VendingMachineContext) {
    fmt.Println("Dispensing product.")
    context.setState(&DispensingState{})
}

func (s *DispensingState) InsertCoin(context *VendingMachineContext) {
    fmt.Println("Product dispensing. Please wait.")
}

func (s *DispensingState) EjectCoin(context *VendingMachineContext) {
    fmt.Println("Product dispensing. Cannot eject coin now.")
}

func (s *DispensingState) Dispense(context *VendingMachineContext) {
    fmt.Println("Product already dispensing.")
}

type VendingMachineContext struct {
    state State
}

func NewVendingMachineContext() *VendingMachineContext {
    return &VendingMachineContext{
        state: &IdleState{},
    }
}

func (c *VendingMachineContext) setState(state State) {
    c.state = state
}

func (c *VendingMachineContext) InsertCoin() {
    c.state.InsertCoin(c)
}

func (c *VendingMachineContext) EjectCoin() {
    c.state.EjectCoin(c)
}

func (c *VendingMachineContext) Dispense() {
    c.state.Dispense(c)
}

func main() {
    vendingMachine := NewVendingMachineContext()

    vendingMachine.InsertCoin()
    vendingMachine.Dispense()

    vendingMachine.InsertCoin()
    vendingMachine.EjectCoin()
    vendingMachine.Dispense()

    vendingMachine.InsertCoin()
    vendingMachine.Dispense()
}
```

```bash
$ go run state.go
Coin inserted. Machine now has a coin.
Dispensing product.
Product dispensing. Please wait.
Product dispensing. Cannot eject coin now.
Product already dispensing.
Product dispensing. Please wait.
Product already dispensing.
```

---

## References

- [https://refactoring.guru/design-patterns/state](https://refactoring.guru/design-patterns/state)
