package main

import (
	"fmt"
	"sync"
)

func listToChannel(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func successor(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num + 1
		}
		close(out)
	}()

	return out
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	listAsChannel1 := listToChannel(1, 2, 3, 4, 5)
	listAsChannel2 := listToChannel(10, 11, 12, 13, 14)

	c1 := successor(listAsChannel1)
	c2 := successor(listAsChannel2)

	for num := range merge(c1, c2) {
		fmt.Println(num)
	}
}
