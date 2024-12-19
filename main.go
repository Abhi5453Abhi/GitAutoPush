package main

import "fmt"

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	evenInt := make(chan int)
	oddInt := make(chan int)

	go PrintEven()
	go PrintOdd()

	for i := 0; i <= 10; i++ {
		select {
		case <-even:
			fmt.Println(evenInt)
		case <-odd:
			fmt.Println(oddInt)
		}
	}
}
