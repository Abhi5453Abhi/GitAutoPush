package main

import (
	"sync"
)

func PrintEven(ch chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		ch <- i
	}
}

func PrintOdd(ch chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		ch <- i
	}
}

func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int)
	done := make(chan bool)
	// odd := make(chan bool)

	wg.Add(2)
	go PrintEven(ch, done, &wg)
	go PrintOdd(ch, done, &wg)

}
