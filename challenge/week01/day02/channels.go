package main

import "fmt"

func main() {
	myChannel := make(chan string)

	// anonymous function
	go func() {
		myChannel <- "some data"
	}()

	message := <-myChannel

	fmt.Println(message)
}
