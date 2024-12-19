package main

import "sync"

func PrintPing(ping chan bool, pong chan bool, wg *s)

func main() {
	ping := make(chan bool, 2)
	pong := make(chan bool, 2)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go PrintPing()
	go PrintPong()

	ping <- true
	wg.Wait()
}
