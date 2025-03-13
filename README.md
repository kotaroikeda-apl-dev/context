## **概要**

## **実行方法**

```sh
go run cmd/WithValue/main.go # context.WithValueの基本実装
go run cmd/workerPool/main.go # 複数のジョブを並行して実行し、結果を集約
go run cmd/withCancel/main.go # キャンセルを実装
```

## **学習ポイント**

1.  **`context.WithValue`** を使うと、関数間で安全にデータ（リクエスト ID など）を受け渡せる。
2.  カスタム型 **`contextKey`** を使うことで、キーの衝突を防ぎ、安全に **`context.Value()`** を管理できる。
3.  **`go func(ctx context.Context, id int) { ... }(ctx, i)`** を使うことで、ループの **`i`** の値が変わる影響を防げる。
4.  各 Goroutine に **`context.WithValue`** で異なるデータ（リクエスト ID, 優先度）を渡すことで、Worker に個別のタスクを持たせる設計の基本を学べる。
5.  **`context.WithCancel`** を使うと、キャンセルを実装できる。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
