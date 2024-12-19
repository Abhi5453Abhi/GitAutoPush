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
	go PrintOdd(even, odd, evenInt, oddInt)

	for i := 0; i <= 10; i++ {
		select {
		case val1, ok := <-even:
			if ok {
				fmt.Println(val1)
			}
		case val2, ok := <-odd:
			if ok {
				fmt.Println(val2)
			}
		}
	}
	even <- true
}
