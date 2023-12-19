package main

import (
	"fmt"
	"math/rand"
	"time"
)

func slowFunc() {
	fmt.Println("Function that takes time...")
	randomDuration := time.Duration(1+rand.Intn(3)) * time.Second
	time.Sleep(randomDuration)
}

func logFuncDuration(theFunc func()) {
	defer func(t time.Time) {
		fmt.Printf("[LOG] Function took %v\n", time.Since(t))
	}(time.Now())
	theFunc()
}

func main() {
	logFuncDuration(slowFunc)
}
