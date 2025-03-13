package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
	priorityKey  contextKey = "priority"
)

func worker(ctx context.Context, id int) {
	reqID := ctx.Value(requestIDKey)
	priority := ctx.Value(priorityKey)

	fmt.Printf("Worker %d: リクエスト ID=%v, 優先度=%v\n", id, reqID, priority)
	time.Sleep(2 * time.Second)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		ctx := context.WithValue(context.Background(), requestIDKey, fmt.Sprintf("req-%d", i))
		ctx = context.WithValue(ctx, priorityKey, i%2) // 0 or 1 (偶数優先)

		go func(ctx context.Context, id int) {
			defer wg.Done()
			worker(ctx, id)
		}(ctx, i)
	}

	wg.Wait()
	fmt.Println("すべてのジョブ完了")
}
