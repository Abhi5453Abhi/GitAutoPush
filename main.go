package main

import "sync"

func PrintPing(ping chan bool, pong chan bool, wg *s)

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
