package main

import "sync"

func main() {
	wg := sync.WaitGroup{}

	even := make(chan bool)
	odd := make(chan bool)

	wg.Add(2)
	go PrintEven(even, odd, &wg)
	go PrintOdd(even, odd, &wg)

	even <- true
	wg.Wait()
}
