package main

func main() {
	even := make(chan bool)
	odd := make(chan bool)

	go PrintEven()
	go PrintOdd()

	for i := 0; i <= 10; i++ {
		select{
			cas
		}
	}
}
