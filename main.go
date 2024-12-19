package main

import "fmt"

func PrintEven(even chan bool, odd chan bool, evenInt chan int, oddInt chan int) {
	for i := 0; i <= 10; i += 2 {
		<-even       // Wait for permission to print even
		evenInt <- i // Send even number to the evenInt channel
		odd <- true  // Notify that the odd goroutine can print
	}
	close(evenInt) // Close evenInt channel after finishing
}

func PrintOdd(even chan bool, odd chan bool, evenInt chan int, oddInt chan int) {
	for i := 1; i <= 10; i += 2 {
		<-odd        // Wait for permission to print odd
		oddInt <- i  // Send odd number to the oddInt channel
		even <- true // Notify that the even goroutine can print
	}
	close(oddInt) // Close oddInt channel after finishing
}

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	evenInt := make(chan int)
	oddInt := make(chan int)

	go PrintEven(even, odd, evenInt, oddInt)
	go PrintOdd(even, odd, evenInt, oddInt)

	// Start by allowing the even goroutine to print first
	even <- true

	// Print the numbers in order
	for {
		select {
		case evenNum, ok := <-evenInt:
			if ok {
				fmt.Print(evenNum)
			}
		case oddNum, ok := <-oddInt:
			if ok {
				fmt.Print(oddNum)
			}
		}
		// Exit condition when both channels are closed
		if len(evenInt) == 0 && len(oddInt) == 0 {
			break
		}
	}
}
