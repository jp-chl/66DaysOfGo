package main

import (
	"fmt"
	"sync"
)

func listToChannel(done <-chan bool, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case out <- num:
			case <-done:
				return
			}
		}
	}()

	return out
}

func successor(done <-chan bool, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-done:
				return
			}

		}
	}()

	return out
}

func merge(done <-chan bool, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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
	done := make(chan bool)
	defer close(done)

	listAsChannel1 := listToChannel(done, 1, 2, 3, 4, 5)
	listAsChannel2 := listToChannel(done, 10, 11, 12, 13, 14)

	c1 := successor(done, listAsChannel1)
	c2 := successor(done, listAsChannel2)

	for num := range merge(done, c1, c2) {
		fmt.Println(num)
	}
}
