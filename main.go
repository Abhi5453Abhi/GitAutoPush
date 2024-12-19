package main

func getNumbers(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	clos
}

func main() {
	ch := make(chan int)

	go getNumbers(ch)

	PrintSquares(ch)
}
