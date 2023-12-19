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
