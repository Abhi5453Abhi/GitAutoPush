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

	go func() {
		for i := 0; i <= 100; i++ {
			ch <- i
		}
	}()

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go WorkerFunction(ch, &wg, res)
	}

	set.Timeout(time.Second * 4)

	go func() {
		wg.Wait()
		close(res)
	}()
	for i := range res {
		fmt.Println(i)
	}

}
