package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func thirdPartyApi(ctx context.Context) error {
	time.Sleep(400 * time.Millisecond)
	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("TLE")
	}
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
