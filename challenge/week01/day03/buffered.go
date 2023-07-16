package main

import "fmt"

func wrong() {
	charChannel := make(chan string, 3)

	mySlice := []string{"a", "b", "c"}

	// go func() {
	// 	for _, letter := range mySlice {
	// 		charChannel <- letter
	// 	}
	// }()

	for _, letter := range mySlice {
		charChannel <- letter
	}

	// close(charChannel)

	for message := range charChannel {
		fmt.Println(message)
	}
}

func right() {
	intChannel := make(chan int, 3)

	mySlice := []int{1, 2, 3}

	// for _, number := range mySlice {
	// 	intChannel <- number
	// }
	for _, number := range mySlice {
		select {
		case intChannel <- number:
		}
	}

	close(intChannel)

	for message := range intChannel {
		fmt.Println(message)
	}
}

func other() {
	myChannel1 := make(chan int, 3)
	myChannel2 := make(chan int, 5)

	mySlice := []int{1, 2, 3, 4, 5, 6}

	for _, number := range mySlice {
		select {
		case myChannel1 <- number:
			fmt.Println("Sending to Channel 1")
		case myChannel2 <- number:
			fmt.Println("Sending to Channel 2")
		}
	}

	close(myChannel1)
	close(myChannel2)

	for message := range myChannel2 {
		fmt.Println("Receiving number", message)
	}

	for message := range myChannel1 {
		fmt.Println("Receiving number", message)
	}

}

func main() {
	other()
}
