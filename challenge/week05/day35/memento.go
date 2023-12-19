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
