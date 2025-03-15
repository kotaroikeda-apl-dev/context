package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func processJob(ctx context.Context, jobID int, errCh chan<- error) {
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("ジョブ %d: キャンセルされました\n", jobID)
			errCh <- ctx.Err()
			return
		default:
			if rand.Intn(10) < 5 { // 50%の確率でエラー
				fmt.Printf("ジョブ %d: 失敗、リトライ %d 回目\n", jobID, i+1)
				time.Sleep(time.Duration(i+1) * time.Second)
				continue
			}
			fmt.Printf("ジョブ %d: 成功\n", jobID)
			errCh <- nil
			return
		}
	}
	errCh <- errors.New("最大リトライ回数超過")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	errCh := make(chan error, 5) // 5つのジョブ分のバッファ

	for i := 1; i <= 5; i++ {
		go processJob(ctx, i, errCh)
	}

	// すべてのジョブの結果を取得
	for i := 1; i <= 5; i++ {
		err := <-errCh
		if err != nil {
			fmt.Println("エラー:", err)
		}
	}

	time.Sleep(12 * time.Second)
	fmt.Println("メイン処理終了")
}
