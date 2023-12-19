package main

import "fmt"

type motor struct {
	cylinders int
	power     int
}

type shift struct {
	speeds int
}

type iMotor interface {
	setCylinders(cylinders int)
	setPower(power int)
	getCylinders() int
	getPower() int
}

type iShift interface {
	setSpeeds(speeds int)
	getSpeeds() int
}

func (m *motor) setCylinders(cylinders int) {
	m.cylinders = cylinders
}
func (m *motor) setPower(power int) {
	m.power = power
}
func (m *motor) getCylinders() int {
	return m.cylinders
}
func (m *motor) getPower() int {
	return m.power
}

func (s *shift) setSpeeds(speeds int) {
	s.speeds = speeds
}
func (s *shift) getSpeeds() int {
	return s.speeds
}

type iCarFactory interface {
	makeMotor() iMotor
	makeShift() iShift
}

type bmwMotor struct {
	motor
}
type bmwShift struct {
	shift
}

type porscheMotor struct {
	motor
}
type porscheShift struct {
	shift
}

type bmw struct {
}

type porsche struct {
}

func (b *bmw) makeMotor() iMotor {
	return &bmwMotor{
		motor: motor{
			cylinders: 6,
			power:     300,
		},
	}
}
func (b *bmw) makeShift() iShift {
	return &bmwShift{
		shift: shift{
			speeds: 6,
		},
	}
}

func (p *porsche) makeMotor() iMotor {
	return &porscheMotor{
		motor: motor{
			cylinders: 7,
			power:     350,
		},
	}
}
func (p *porsche) makeShift() iShift {
	return &porscheShift{
		shift: shift{
			speeds: 7,
		},
	}
}

func getCarFactory(brand string) (iCarFactory, error) {
	if brand == "bmw" {
		return &bmw{}, nil
	}
	if brand == "porsche" {
		return &porsche{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

func main() {
	bmwFactory, _ := getCarFactory("bmw")
	porscheFactory, _ := getCarFactory("porsche")

	bmwMotorInstance := bmwFactory.makeMotor()
	bmwShiftInstance := bmwFactory.makeShift()

	porscheMotorInstance := porscheFactory.makeMotor()
	porscheShiftInstance := porscheFactory.makeShift()

	fmt.Println("Brand Bmw")
	demoMotor(bmwMotorInstance)
	demoShift(bmwShiftInstance)

	fmt.Println()
	fmt.Println("Brand Porsche")
	demoMotor(porscheMotorInstance)
	demoShift(porscheShiftInstance)
}

func demoMotor(m iMotor) {
	fmt.Printf("Motor Cylinders: %d", m.getCylinders())
	fmt.Println()
	fmt.Printf("Motor Power: %d", m.getPower())
	fmt.Println()
}

func demoShift(s iShift) {
	fmt.Printf("Shift Speeds: %d", s.getSpeeds())
	fmt.Println()
}
