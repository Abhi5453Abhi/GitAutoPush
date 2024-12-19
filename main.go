package main

import "fmt"

func getNumbers(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func PrintSquares(ch chan int) {
	for i := range ch {
		fmt.Println(i * i)
	}
}

func main() {
	wg := s
	ch := make(chan int)

	go getNumbers(ch)

	PrintSquares(ch)
}
