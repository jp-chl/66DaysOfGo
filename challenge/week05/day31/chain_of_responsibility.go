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
