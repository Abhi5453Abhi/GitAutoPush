/**

Worker pool pattern

**/

package main

import (
	"sync"
)

func GetSquare(val int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- val
}

func WorkerFunction(ch chan int, wg *sync.WaitGroup, res chan int) {
	wg.Done()
	val := <-ch
	res <- val * val
}

func GetResult(val int, result chan int) {
	result <- val * val
}

func main() {
	ch := make(chan int, 100)
	res := make(chan int, 100)
	wg := sync.WaitGroup{}

	for i := 0; i <= 3; i++ {
		go WorkerFunction(ch, &wg, res)
	}

	for i := 0; i <= 100; i++ {
		ch <- i
	}

	for i := range res {

	}

	close(ch)

}
