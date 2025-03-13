## **概要**

## **実行方法**

```sh
go run cmd/WithValue/main.go # context.WithValueの基本実装
```

## **学習ポイント**

1.  **`context.WithValue`** を使うと、関数間で安全にデータ（リクエスト ID など）を受け渡せる。
2.  カスタム型 **`contextKey`** を使うことで、キーの衝突を防ぎ、安全に **`context.Value()`** を管理できる。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
