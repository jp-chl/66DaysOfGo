package main

import (
	"fmt"
	"time"
)

func myFunc(num int) {
	fmt.Println(num)
}

func main() {
	go myFunc(1) // each go call will spawn an independent child process
	go myFunc(2)
	go myFunc(3)
	go myFunc(4)

	time.Sleep(time.Second * 1)

	fmt.Println("end...")
}
