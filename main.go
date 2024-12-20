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

func WorkerFunction(val int) {

}

func GetResult(val int, result chan int) {
	result <- val * val
}

func main() {
	ch := make(chan int, 100)
	res := make(chan int, 100)
	wg := sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		ch <- i
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		go WorkerFunction(i)
	}

}
