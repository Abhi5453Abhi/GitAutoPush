package main

import (
	"fmt"
	"sync"
)

func PrintPing(ping chan bool, pong chan bool, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ping:
			fmt.Println("Ping")
			pong <- true
		case <-done:
			return
		}
	}
}

func PrintPong(ping chan bool, pong chan bool, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-pong:
			fmt.Println("Pong")
			ping <- true
		case <-done:
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	ping := make(chan bool, 1) // Buffer of 1 is enough
	pong := make(chan bool, 1)
	done := make(chan bool)

	wg.Add(2)
	go PrintPing(ping, pong, done, &wg)
	go PrintPong(ping, pong, done, &wg)

	// Start the ping-pong sequence
	ping <- true

	// Allow 10 iterations
	for i := 0; i < 10; i++ {
		// Let the ping-pong goroutines handle the work
	}

	// Signal the goroutines to stop
	close(done)

	// Wait for the goroutines to finish
	wg.Wait()
}
