package main

import (
	"context"
	"time"
)

func thirdPartyApi(ctx context.Context) error {
	time.Sleep(400 * time.Millisecond)
	return nil
}

func main() {
	ctx := context.Background()
	ctxWithTo := ctx.contextWithTimeOut(ctx, 200*time.Millisecond)

	err := thirdPartyApi(ctx)
}
