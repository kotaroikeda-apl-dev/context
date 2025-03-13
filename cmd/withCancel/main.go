package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: キャンセルされました\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker %d: ジョブ %d を処理中\n", id, job)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	jobs := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	time.Sleep(5 * time.Second)
	cancel() // 途中でキャンセル

	close(jobs)
	wg.Wait()
	fmt.Println("すべてのジョブ完了")
}
