package main

import (
	"fmt"
	"math/rand"
)

type WeatherProvider interface {
	GetCelsius() int
}

type Provider1 struct{}
type Provider2 struct{}

func (p Provider1) GetCelsius() int {
	return rand.Intn(45)
}

func (p Provider2) GetFahrenheit() int {
	return rand.Intn(115)
}

type AdapterProvider2 struct {
	provider Provider2
}

func (p AdapterProvider2) GetCelsius() int {
	return int((float64)(p.provider.GetFahrenheit()-32) / 1.8)
}

func main() {
	weatherInCity1 := Provider1{}
	fmt.Print("City 1. ")
	demoTemperature(weatherInCity1)

	weatherInCity2 := AdapterProvider2{
		provider: Provider2{},
	}
	fmt.Print("City 2. ")
	demoTemperature(weatherInCity2)
}

func demoTemperature(provider WeatherProvider) {
	fmt.Printf("Weather in celsius: %v\n", provider.GetCelsius())
}
