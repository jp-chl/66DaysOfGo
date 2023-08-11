package main

import (
	"fmt"
	"sync"
)

type MyCustomData struct {
	data string
}

type SingletonProvider struct {
	once              sync.Once
	singletonInstance *MyCustomData
}

func (sp *SingletonProvider) GetInstance() *MyCustomData {
	sp.once.Do(func() {
		sp.singletonInstance = &MyCustomData{
			data: "test data",
		}
	})
	return sp.singletonInstance
}

func main() {
	provider := &SingletonProvider{}

	// several goroutines
	for i := 0; i < 10; i++ {
		go func(index int) {
			uniqueInstance := provider.GetInstance()
			demo(index, uniqueInstance)
		}(i)
	}

	fmt.Scanln() // wait for input before exiting
}

func demo(index int, variable interface{}) {
	fmt.Printf("index = %d, address of (%v) is: %p\n", index, variable, variable)
}
