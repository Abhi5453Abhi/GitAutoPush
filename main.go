package main

func PrintPing(ping chan bool, pong chan bool, done chan bool) {
	// defer wg.Done()
	// for i := 0; i < 10; i++ {
	// 	<-ping
	// 	fmt.Println("Ping")
	// 	pong <- true
	// }
	for {
		s
	}
}

func PrintPong(ping chan bool, pong chan bool, done chan bool) {
	// defer wg.Done()
	// for i := 0; i < 10; i++ {
	// 	<-pong
	// 	fmt.Println("Pong")
	// 	ping <- true
	// }
}

func main() {
	// wg := sync.WaitGroup{}

	ping := make(chan bool, 10)
	pong := make(chan bool, 10)
	done := make(chan bool)

	// wg.Add(2)
	go PrintPing(ping, pong, done)
	go PrintPong(ping, pong, done)
	ping <- true

	for i := 0; i <= 10; i++ {

	}

	// wg.Wait()
	close(done)

}
