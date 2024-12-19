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
}

func PrintOdd(even chan bool, odd chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-odd
		fmt.Println(i)
		even <- true
	}
}

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
