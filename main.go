package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	// Even goroutine
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i += 2 {
			<-even // Wait for signal
			fmt.Print(i)
			if i < 10 {
				odd <- true // Signal the odd goroutine
			}
		}
	}()

	// Odd goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			<-odd // Wait for signal
			fmt.Print(i)
			even <- true // Signal the even goroutine
		}
	}()

	// Start the process
	even <- true // Trigger the even goroutine
	wg.Wait()    // Wait for both goroutines to complete
	fmt.Println()
}
