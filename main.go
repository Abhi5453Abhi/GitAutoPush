package main

import (
	"fmt"
	"sync"
)

func PrintPing(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	for i := 0; i <= 10; i += 2 {
		<-ping
		fmt.Println("ping")
		pong <- true
	}
}

func PrintPong(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	for i := 0; i <= 10; i += 2 {
		<-pong
		fmt.Println("pong")
		ping <- true
	}
	close(ping)
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
