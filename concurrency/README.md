# ✅ 1. Goroutine（軽量スレッド）

## 概要:

`go` キーワードを使うだけで関数を非同期に実行できます（並行処理の開始）。

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {
	go sayHello() // 新しい goroutine として実行
	time.Sleep(1 * time.Second) // main関数が先に終わらないように待つ
}
```

📌 `goroutine` は非常に軽量で、数万個同時に扱えます。

# ✅ 2. Channel（goroutine 間の通信）

## 概要:

chan 型を使って、gorountine 間でデータをやり取りします。

```go
package main

import "fmt"

func worker(ch chan string) {
	ch <- "Finished work" // chに値を送る
}

func main() {
	ch := make(chan string)
	go worker(ch)

	msg := <-ch // ch から値を受け取る（ブロッキング）
	fmt.Println(msg)
}
```

📌 `chan T` は「T 型の値を送受信するチャネル」

# ✅ 3. Select 文（チャネルの multiplex）

## 概要:

複数のチャネルからの受信を待つときに使用。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
```

📌 最初に値が来たチャネルが選ばれる

# ✅ 4. Context（キャンセル・タイムアウトの制御）

## 概要:

外部から goroutine を安全にキャンセルしたいときに使います。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longTask(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task canceled")
			return
		default:
			time.Sleep(500 * time.Millisecond)
			ch <- "Working..."
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	ch := make(chan string)
	go longTask(ctx, ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
```

📌 `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline` などが使えます。

# 📦 実用で意識すべきこと

| 概念      | 使い所                               |
| --------- | ------------------------------------ |
| goroutine | 並列処理、複数 API 呼び出し等        |
| channel   | 処理結果や状態の通知                 |
| select    | 複数チャネルの待機、タイムアウト処理 |
| context   | タイムアウト制御、キャンセル伝播     |

# Defer

## ✅ 基本構文

```go
package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("This runs at the end")
	fmt.Println("End")
}
```

## 実行結果：

```sh
Start
End
This runs at the end
```

📌 `defer` した関数はその関数が終了するときに実行される。

## ✅ 複数の defer は「スタック形式」で実行される

```go
func main() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}
```

## 実行結果：

```sh
third
second
first
```

📌 `defer` は LIFO（後入れ先出し） の順序で実行されます。

## ✅ よくある用途

### 1. ファイルのクローズ

```go
file, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
// 読み込み処理など
```

### 2. ロックの解除

```go
var mu sync.Mutex

func criticalSection() {
    mu.Lock()
    defer mu.Unlock() // パニックやリターンでも必ず解除される
    // 処理...
}
```

### 3. Panic と Recover の組み合わせ

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Something went wrong")
}
```

## ✅ defer の注意点

1. 即時評価されるのは引数だけ
   呼び出しは遅延するが、引数の評価は `defer` 宣言時に行われる。

```go
func main() {
    x := 10
    defer fmt.Println("x =", x) // ここで x = 10 と評価される
    x = 20
}
```

出力は `x = 10`

2. パフォーマンスに若干の影響

高速ループ内では使いすぎない方がよい（不要なコストになることもある）

# 🔑 まとめ

| 特徴                 | 説明                                       |
| -------------------- | ------------------------------------------ |
| 実行タイミング       | 関数終了時                                 |
| 実行順               | スタック（後入れ先出し）                   |
| 主な用途             | リソース解放、ロック解除、ログ、Panic 処理 |
| 引数の評価タイミング | defer 宣言時                               |
