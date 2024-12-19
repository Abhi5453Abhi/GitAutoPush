package main

func main() {
	even := make(chan bool, 10)
	odd := make(chan bool, 10)
	wg := s

	go PrintEven()
	go PrintOdd()

}
