package main

import (
	"fmt"
	"sync"
)

func getNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func PrintSquares(ch chan int) {
	for i := range ch {
		fmt.Println(i * i)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go getNumbers(ch, &wg)
	wg.Wait()

	PrintSquares(ch)
}
