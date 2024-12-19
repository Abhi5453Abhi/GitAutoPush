package main

import "fmt"

func PrintEven(even chan bool, odd chan bool, evenInt chan int, oddInt chan int) {
	for i := 0; i <= 10; i += 2 {
		<-even
		evenInt <- i
		odd <- true
	}
}

func PrintOdd(even chan bool, odd chan bool, evenInt chan int, oddInt chan int) {
	for i := 1; i <= 10; i += 2 {
		<-odd
		oddInt <- i
		odd <- true
	}
}

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	evenInt := make(chan int)
	oddInt := make(chan int)

	go PrintEven(even, odd, evenInt, oddInt)
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
