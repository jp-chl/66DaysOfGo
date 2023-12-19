package main

import "fmt"

type Car interface {
	print()
	clone() Car
}

type Bmw struct {
	year  int
	color string
}

func (b *Bmw) clone() Car {
	return &Bmw{
		year:  b.year,
		color: b.color,
	}
}

func (b *Bmw) print() {
	fmt.Printf("Brand: BMW, Year: %d, Color: %s\n", b.year, b.color)
}

type Porsche struct {
	brand string
	year  int
	power int
}

func (p *Porsche) clone() Car {
	return &Porsche{
		year:  p.year,
		power: p.power,
	}
}

func (p *Porsche) print() {
	fmt.Printf("Brand: Porsche, Year: %d, Power: %d\n", p.year, p.power)
}

func main() {
	car1 := &Bmw{
		year:  2000,
		color: "red",
	}
	car2 := &Porsche{
		year:  2000,
		power: 500,
	}
	demoCar(car1)
	demoCar(car2)

	fmt.Println("\nCloning...")
	car3 := car1.clone()
	car4 := car2.clone()
	demoCar(car3)
	demoCar(car4)
}

func demoCar(c Car) {
	fmt.Printf("Variable memory address: %v. ", &c)
	c.print()
}
