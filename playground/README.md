# 実行方法

```sh
go run main.go
```

# プロジェクト全体を実行

**Go モジュールが必要**

## Go モジュール初期化

```sh
go mod init <module_name>
```

## 実行

```sh
go run .
```

# Tips

## 公開関数

外部から関数呼び出す場合は、関数名のキャピタルを**大文字**にする必要がある

```go
// public
func GetUser() {}

// private
func validateUserInfo() {}
```
