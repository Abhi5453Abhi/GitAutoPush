package main

import (
	"context"
	"fmt"
	"time"
)

func thirdPartyApi(ctx context.Context) error {
	time.Sleep(400 * time.Millisecond)
	return nil
}

func main() {
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	err := thirdPartyApi(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
