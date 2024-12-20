/**

Worker pool pattern

**/

package main

import "sync"

func GetSquare(val int, wg *s)

func main() {
	ch := make(chan int, 1000)
	wg := sync.WaitGroup{}

	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go GetSquare(i)
	}
}
