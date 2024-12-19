package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create channels for communication
	even := make(chan bool)
	odd := make(chan bool)
	var wg sync.WaitGroup

	// Add 2 to WaitGroup for the two goroutines
	wg.Add(2)

	// Launch a goroutine to print even numbers
	go func() {
		defer wg.Done() // Mark this goroutine as done when finished
		for i := 0; i <= 10; i += 2 {
			<-even // Wait for signal from main routine or odd goroutine
			fmt.Print(i)
			odd <- true // Signal the odd goroutine
		}
	}()

	// Launch a goroutine to print odd numbers
	go func() {
		defer wg.Done() // Mark this goroutine as done when finished
		for i := 1; i <= 10; i += 2 {
			<-odd // Wait for signal from the even goroutine
			fmt.Print(i)
			even <- true // Signal the even goroutine
		}
	}()

	// Start the printing process
	even <- true // Trigger the even goroutine
	wg.Wait()    // Wait for all goroutines to finish
	<-odd

	// End of main function
	fmt.Println()
}
