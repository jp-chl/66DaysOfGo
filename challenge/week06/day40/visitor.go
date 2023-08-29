package main

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	VisitConcreteElement1(element *ConcreteElement1)
	VisitConcreteElement2(element *ConcreteElement2)
	VisitConcreteElement3(element *ConcreteElement3)
}

type ConcreteElement1 struct {
	name string
}
type ConcreteElement2 struct {
	name string
}
type ConcreteElement3 struct {
	name string
}

func (c *ConcreteElement1) Accept(visitor Visitor) {
	visitor.VisitConcreteElement1(c)
}
func (c *ConcreteElement2) Accept(visitor Visitor) {
	visitor.VisitConcreteElement2(c)
}
func (c *ConcreteElement3) Accept(visitor Visitor) {
	visitor.VisitConcreteElement3(c)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElement1(element *ConcreteElement1) {
	fmt.Println("Visited ConcreteElement 1:", element.name)
}
func (v *ConcreteVisitor) VisitConcreteElement2(element *ConcreteElement2) {
	fmt.Println("Visited ConcreteElement 2:", element.name)
}
func (v *ConcreteVisitor) VisitConcreteElement3(element *ConcreteElement3) {
	fmt.Println("Visited ConcreteElement 3:", element.name)
}

func main() {
	elements := []Element{
		&ConcreteElement1{name: "Element 1"},
		&ConcreteElement2{name: "Element 2"},
		&ConcreteElement3{name: "Element 3"},
	}

	visitor := &ConcreteVisitor{}

	for _, elem := range elements {
		elem.Accept(visitor)
	}
}
