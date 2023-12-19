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
