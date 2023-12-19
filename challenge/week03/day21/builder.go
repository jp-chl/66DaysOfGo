package main

import (
	"fmt"
)

type Car struct {
	brand string
	color string
	year  int
}

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: Car{},
	}
}

func (cb *CarBuilder) WithBrand(brand string) *CarBuilder {
	cb.car.brand = brand
	return cb
}
func (cb *CarBuilder) WithColor(color string) *CarBuilder {
	cb.car.color = color
	return cb
}
func (cb *CarBuilder) WithYear(year int) *CarBuilder {
	cb.car.year = year
	return cb
}

func (cb *CarBuilder) Build() Car {
	return cb.car
}

func main() {
	bmw := Car{
		brand: "bmw",
		color: "red",
		year:  2000,
	}
	demoCar(bmw)

	porsche := NewCarBuilder().
		WithBrand("porsche").
		WithColor("white").
		// WithYear(2019).
		Build()
	demoCar(porsche)
}

func demoCar(c Car) {
	fmt.Printf("Brand: %s, Color: %s, Year: %d\n", c.brand, c.color, c.year)
}
