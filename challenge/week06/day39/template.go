package main

import "fmt"

type EaterType interface {
	pick()
	prepare()
	ingest()
}

type Eater struct {
	eaterType EaterType
}

func (e *Eater) eat() {
	e.eaterType.pick()
	e.eaterType.prepare()
	e.eaterType.ingest()
}
func (e *Eater) setEaterType(eaterType EaterType) {
	e.eaterType = eaterType
}

type MeatEater struct{}
type Vegetarian struct{}
type Vegan struct{}

func (m *MeatEater) pick() {
	fmt.Println("Selecting meat with high protein")
}
func (m *MeatEater) prepare() {
	fmt.Println("Adding vegetables, and roasting the meat")
}
func (m *MeatEater) ingest() {
	fmt.Println("Enjoying meat + vegetables")
}

func (v *Vegetarian) pick() {
	fmt.Println("Selecting beans, eggs and vegetables")
}
func (v *Vegetarian) prepare() {
	fmt.Println("Adding pepper, salt and spicy stuff")
}
func (v *Vegetarian) ingest() {
	fmt.Println("Enjoying a healthy food made of eggs, beans and vegetables")
}

func (v *Vegan) pick() {
	fmt.Println("Selecting only non-animal food")
}
func (v *Vegan) prepare() {
	fmt.Println("Preparing a healthy food")
}
func (v *Vegan) ingest() {
	fmt.Println("Enjoying a super healthy food made of plants")
}

func main() {
	meatEater := &MeatEater{}
	vegetarian := &Vegetarian{}
	vegan := &Vegan{}

	eater := &Eater{
		eaterType: meatEater,
	}
	eater.eat()
	fmt.Println("==================================")

	eater.setEaterType(vegetarian)
	eater.eat()
	fmt.Println("==================================")

	eater.setEaterType(vegan)
	eater.eat()
}
