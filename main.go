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
	ctx := context.Background()
	ctxWithTo, cancel := ctx.contextWithTimeOut(ctx, 200*time.Millisecond)
	defer cancel()

	err := thirdPartyApi(ctxWithTo)

	if err != nil {
		fmt.Println(err)
	}
}
