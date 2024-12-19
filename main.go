package main

import (
	"fmt"
	"sync"
)

func PrintEven(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		<-even
		fmt.Println(i)
		odd <- true
	}
	close(even)
}

func PrintOdd(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-odd
		fmt.Println(i)
		even <- true
	}
	close(odd)
}

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
