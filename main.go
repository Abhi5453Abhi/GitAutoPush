package main

import (
	"fmt"
	"sync"
)

func PrintEven(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	// defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		<-even // Wait for the signal to print an even number
		fmt.Println(i)
		odd <- true // Signal the odd goroutine
	}
}

func PrintOdd(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	// defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		<-odd // Wait for the signal to print an odd number
		fmt.Println(i)
		even <- true // Signal the even goroutine
	}
}

func main() {
	wg := sync.WaitGroup{}

	even := make(chan bool)
	odd := make(chan bool)
	done := make(chan bool)

	// wg.Add(2)
	go PrintEven(even, odd, &wg)
	go PrintOdd(even, odd, &wg)

	even <- true // Start the process by signaling the even goroutine
	// wg.Wait()    // Wait for both goroutines to complete
}
