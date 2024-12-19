package main

func main() {
	ch := make(chan int)

	go getNumbers(ch)

	PrintSquares(ch)
}
