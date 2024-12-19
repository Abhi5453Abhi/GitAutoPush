package main

import "sync"

func PrintEven(even chan bool, odd chan bool, wg *s)

func main() {
	even := make(chan bool, 10)
	odd := make(chan bool, 10)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go PrintEven(even, odd, &wg)
	go PrintOdd(even, odd, &wg)

	even <- true
	wg.Wait()

}
