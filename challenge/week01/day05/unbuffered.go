package main

import "fmt"

// This will raise "fatal error: all goroutines are asleep - deadlock!"
func deadlockSample1() {
	myChannel := make(chan int)

	myChannel <- 2 // This line will raise the error

	number := <-myChannel

	fmt.Println("number = ", number)
}

// This will raise "fatal error: all goroutines are asleep - deadlock!"
func deadlockSample2() {
	myChannel := make(chan int)

	number := <-myChannel // This line will raise the error

	fmt.Println("number = ", number)

}

// This will raise "fatal error: all goroutines are asleep - deadlock!"
func deadlockSample3() {
	myChannel := make(chan int)

	myChannel <- 2 // This line will raise the error
}

// This will execute correctly
func channelSyncOkSample1() {
	myChannel := make(chan int)

	go func() {
		myChannel <- 2
	}()

	number := <-myChannel

	fmt.Println("number = ", number)
}

// This will raise "fatal error: all goroutines are asleep - deadlock!"
func deadlockSample4() {
	myChannel := make(chan int, 1) // buffer capacity: 1

	myChannel <- 1
	myChannel <- 2 // This line will raise the error

	number := <-myChannel

	fmt.Println("number = ", number)
}

// This will execute correctly.
// You don't even need a goroutine with the right buffered channel and its respective consumer(s)
func channelSyncOkSample2() {
	myChannel := make(chan int, 2)

	myChannel <- 1
	myChannel <- 2

	number := <-myChannel
	fmt.Println("number = ", number)

	number = <-myChannel
	fmt.Println("number = ", number)
}

func main() {
	deadlockSample4()
}
