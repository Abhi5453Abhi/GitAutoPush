package main

import (
	"fmt"
	"sync"
)

func PrintPing(ping chan bool, pong chan bool, done chan bool, wg *sync.WaitGroup) {
	// defer wg.Done()
	// for i := 0; i < 10; i++ {
	// 	<-ping
	// 	fmt.Println("Ping")
	// 	pong <- true
	// }
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
	// defer wg.Done()
	// for i := 0; i < 10; i++ {
	// 	<-pong
	// 	fmt.Println("Pong")
	// 	ping <- true
	// }
	for {
		select {
		case <-pong:
			fmt.Println("Ping")
			ping <- true
		case <-done:
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	ping := make(chan bool, 10)
	pong := make(chan bool, 10)
	done := make(chan bool)

	wg.Add(2)
	go PrintPing(ping, pong, done, &wg)
	go PrintPong(ping, pong, done, &wg)
	ping <- true

	for i := 0; i <= 10; i++ {

	}

	wg.Wait()
	close(done)

}
