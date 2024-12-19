package main

import (
	"fmt"
	"sync"
)

func PrintPing(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		<-ping
		fmt.Println("Ping")
		pong <- true
	}

}

func PrintPong(ping chan bool, pong chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		<-pong
		fmt.Println("Pong")
		ping <- true
	}
}

func Divide(a, b int) error {
	dividend := a % b

	fmt.Println(dividend)
	if dividend == 0 {
		return fmt.Errorf("Divisor is 0")
	}
	return nil
}

func main() {
	// ping := make(chan bool, 2)
	// pong := make(chan bool, 2)

	// wg := sync.WaitGroup{}

	// wg.Add(2)
	// go PrintPing(ping, pong, &wg)
	// go PrintPong(ping, pong, &wg)

	// ping <- true
	// wg.Wait()
	err := Divide(4, 2)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
