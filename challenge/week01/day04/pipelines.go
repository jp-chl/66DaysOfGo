package main

import (
	"fmt"
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

func cube(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num * num
		}
		close(out)
	}()

	return out
}

func average(in <-chan int) float64 {
	sum := 0
	elements := 0
	for num := range in {
		sum += num
		elements++
	}
	return (float64)(sum / elements)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	listAsChannel := listToChannel(nums...)
	cubesChannel := cube(listAsChannel)
	result := average(cubesChannel)

	fmt.Println("Average of the cubes is", result)
}
