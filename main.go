package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	thirdPartyApi(ctx)
}

func thirdPartyApi(ctx context.Context) {
	ctxWithTo, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Work is done")
	case <-ctxWithTo.Done():
		fmt.Println("Deadline exceeded")
	}

}
