package main

import "fmt"

// without generics
func integerSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// without generics
func floatSum(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// with generics
func numbersSum[T int | float64](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	intArray := []int{1, 2, 3}
	floatArray := []float64{1.0, 2.0, 3.0}

	fmt.Printf("integer sum: %d\n", integerSum(intArray))
	fmt.Printf("float sum: %.2f\n", floatSum(floatArray))

	fmt.Printf("(generic) integer sum: %d\n", numbersSum(intArray))
	fmt.Printf("(generic) float sum: %.2f\n", numbersSum(floatArray))
}
