package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
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

func main() {
	// ping := make(chan bool, 2)
	// pong := make(chan bool, 2)

	// wg := sync.WaitGroup{}

	// wg.Add(2)
	// go PrintPing(ping, pong, &wg)
	// go PrintPong(ping, pong, &wg)

	// ping <- true
	// wg.Wait()

	r := gin.Default()

	r.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		newId, err := strconv.Atoi(id)
		ctx.JSON(http.StatusAccepted, gin.H{
			"id":      id,
			"Message": "Successfull",
		})
	})

	r.Run(":3000")

}
