package main

import (
	"fmt"
	"sync"
)

func PrintPing(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		select {
		case <-ping:
			fmt.Println("ping")
			pong <- true
		}
	}
}

func PrintPong(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		select {
		case <-pong:
			fmt.Println("pong")
			ping <- true
		default:
			fmt.Println("Ping CLOSSED")
			close(ping)
		}
	}
}

func main() {
	ping := make(chan bool)
	pong := make(chan bool)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go PrintPing(ping, pong, &wg)
	go PrintPong(ping, pong, &wg)
	ping <- true
	wg.Wait()
}
