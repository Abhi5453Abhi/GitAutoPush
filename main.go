package main

func main() {
	ping := make(chan bool)
	pong := make(chan bool)

	wg := s

	go PrintPing()
	go PrintPong()
}
