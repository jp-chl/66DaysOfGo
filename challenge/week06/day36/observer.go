package main

import "fmt"

type Publisher interface {
	addSubscriber(subscriber Subscriber)
	removeSubscriber(subscriberId string)
	broadcast(message string)
}

type Subscriber interface {
	id() string
	receive(message string)
}

type ConcretePublisher struct {
	subscribers map[string]Subscriber
}

func NewConcretePublisher() ConcretePublisher {
	return ConcretePublisher{
		subscribers: make(map[string]Subscriber),
	}
}

func (p ConcretePublisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[subscriber.id()] = subscriber
}

func (p ConcretePublisher) removeSubscriber(subscriberId string) {
	delete(p.subscribers, subscriberId)
}

func (p ConcretePublisher) broadcast(message string) {
	for _, subscriber := range p.subscribers {
		subscriber.receive(message)
	}
}

type ConcreteSubscriber struct {
	subscriberId string
}

func NewConcreteSubscriber(subscriberId string) ConcreteSubscriber {
	return ConcreteSubscriber{
		subscriberId: subscriberId,
	}
}

func (s ConcreteSubscriber) id() string {
	return s.subscriberId
}

func (s ConcreteSubscriber) receive(message string) {
	fmt.Printf("Subscriber %s receiving message: [%s]\n", s.subscriberId, message)
}

func main() {
	publisher := NewConcretePublisher()

	subscriber1 := NewConcreteSubscriber("id1")
	subscriber2 := NewConcreteSubscriber("id2")
	publisher.addSubscriber(subscriber1)
	publisher.addSubscriber(subscriber2)

	publisher.broadcast("hello subscribers")
}
