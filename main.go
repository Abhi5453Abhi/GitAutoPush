package main

import (
	"fmt"
	"sync"
)

func PrintEven(even chan bool, odd chan bool, wg *sync.WaitGroup, done chan bool) {
	// defer wg.Done()
	// defer close(done)
	for i := 0; i <= 10; i += 2 {
		<-even // Wait for the signal to print an even number
		fmt.Println(i)
		odd <- true // Signal the odd goroutine
	}
	done <- true
}

func PrintOdd(even chan bool, odd chan bool, wg *sync.WaitGroup, done chan bool) {
	// defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		<-odd // Wait for the signal to print an odd number
		fmt.Println(i)
		even <- true // Signal the even goroutine
	}
	close(even)
}

func main() {
	wg := sync.WaitGroup{}

	even := make(chan bool)
	odd := make(chan bool)
	done := make(chan bool)

	// wg.Add(2)
	go PrintEven(even, odd, &wg, done)
	go PrintOdd(even, odd, &wg, done)

	even <- true // Start the process by signaling the even goroutine
	// wg.Wait()    // Wait for both goroutines to complete
	<-done
}
