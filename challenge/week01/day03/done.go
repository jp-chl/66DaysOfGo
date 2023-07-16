package main

import (
	"fmt"
	"time"
)

func working(done <-chan bool) {
	for {
		select {
		default:
			fmt.Println("working...")
			time.Sleep(time.Second * 1)
		case <-done:
			fmt.Println("stopping")
			return
		}
	}
}

func main() {
	done := make(chan bool)

	go working(done)

	time.Sleep(time.Second * 3)

	close(done)

	// To print "stopping", you can add a delay after close
	//time.Sleep(time.Millisecond * 5)
}
