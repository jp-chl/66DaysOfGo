package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "data for channel 1"
	}()

	go func() {
		channel2 <- "data for channel 2"
	}()

	select {
	case msgFromChannel1 := <-channel1:
		fmt.Println(msgFromChannel1)
	case msgFromChannel2 := <-channel2:
		fmt.Println(msgFromChannel2)
	}
}
