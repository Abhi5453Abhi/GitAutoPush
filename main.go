package main

import (
	"fmt"
	"sync"
)

func PrintPing(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ping
		fmt.Println("Ping")
		pong <- true
	}
}

func PrintPong(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ping
		fmt.Println("Ping")
		pong <- true
	}
}

func main() {
	wg := sync.WaitGroup{}

	ping := make(chan bool, 10)
	pong := make(chan bool, 10)

	wg.Add(2)
	go PrintPing()
	go PrintPong()
	go func() {

	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-pong
			fmt.Println("Pong")
			ping <- true
		}
	}()
	ping <- true
	wg.Wait()

}
