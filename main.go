/**

Worker pool pattern

**/

package main

import (
	"fmt"
	"sync"
)

func GetSquare(val int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- val * val
}

func main() {
	ch := make(chan int, 10000)
	wg := sync.WaitGroup{}

	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go GetSquare(i, &wg, ch)
	}

	wg.Wait()

	for i := range ch {
		fmt.Println(i)
	}
}
