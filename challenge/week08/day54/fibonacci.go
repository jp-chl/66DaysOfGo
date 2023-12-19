package main

func RecursiveFibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}

func SequentialFibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	var n1 uint = 1
	var n2 uint = 0
	for i := uint(2); i < n; i++ {
		temp := n1 + n2
		n2 = n1
		n1 = temp
	}
	return n2 + n1
}
