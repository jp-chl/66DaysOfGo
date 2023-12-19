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
