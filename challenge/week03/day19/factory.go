package main

import (
	"fmt"
)

type iCar interface {
	setName(name string)
	getName() string
	setPower(hp int)
	getPower() int
}

type car struct {
	name  string
	power int
}

func (c *car) setName(name string) {
	c.name = name
}
func (c *car) getName() string {
	return c.name
}
func (c *car) setPower(power int) {
	c.power = power
}
func (c *car) getPower() int {
	return c.power
}

type bmw struct {
	car
}

func newBmw() iCar {
	return &bmw{
		car: car{
			name:  "bmw",
			power: 300,
		},
	}
}

type mercedes struct {
	car
}

func newMercedes() iCar {
	return &mercedes{
		car: car{
			name:  "mercedes",
			power: 400,
		},
	}
}

func getCar(carType string) (iCar, error) {
	switch carType {
	case "bmw":
		return newBmw(), nil
	case "mercedes":
		return newMercedes(), nil
	default:
		return nil, fmt.Errorf("Unknown car type: [%s]", carType)
	}
}

func printCarDetails(c iCar) {
	fmt.Printf("Car: %s, power: %d", c.getName(), c.getPower())
	fmt.Println()
}

func demoCar(carType string) {
	car, err := getCar(carType)
	if err == nil {
		printCarDetails(car)
	} else {
		fmt.Printf(err.Error())
	}
}

func main() {
	demoCar("mercedes")
	fmt.Println()
	demoCar("bmw")
}
