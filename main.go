package main

import (
	"fmt"
)

func main() {
	// Create channels for communication
	even := make(chan bool)
	odd := make(chan bool)

	// Launch a goroutine to print even numbers
	go func() {
		for i := 0; i <= 10; i += 2 {
			<-even // Wait for signal from the main routine
			fmt.Print(i)
			odd <- true // Signal the odd goroutine
		}
	}()

	// Launch a goroutine to print odd numbers
	go func() {
		for i := 1; i <= 10; i += 2 {
			<-odd // Wait for signal from the even goroutine
			fmt.Print(i)
			even <- true // Signal the even goroutine
		}
	}()

	// Start the printing process
	even <- true // Trigger the even goroutine
	<-odd        // Wait for the last odd number to complete

	// End of main function
	fmt.Println()
}
