package main

import (
	"context"
	"fmt"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func processRequest(ctx context.Context) {
	reqID := ctx.Value(requestIDKey)
	fmt.Println("リクエスト ID:", reqID)
}

func main() {
	ctx := context.WithValue(context.Background(), requestIDKey, "abc-123")
	processRequest(ctx)
}
