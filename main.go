/**

Worker pool pattern

**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func GetSquare(val int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- val
}

func WorkerFunction(ch chan int, wg *sync.WaitGroup, res chan int) {
	defer wg.Done() //execute once all the rest of things are done in this function
	for i := range ch {
		res <- i * i
	}

}

func GetResult(val int, result chan int) {
	result <- val * val
}

func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	ch := make(chan int, 100)
	res := make(chan int, 100)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go WorkerFunction(ch, &wg, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()
	for i := range res {
		fmt.Println(i)
	}

}
