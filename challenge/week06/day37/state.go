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
