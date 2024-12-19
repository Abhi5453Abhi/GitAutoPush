package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	ping := make(chan bool)
	pong := make(chan bool)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ping
			fmt.Println("Ping")
			pong <- true
		}
	}()
}
