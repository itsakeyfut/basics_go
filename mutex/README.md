# 🔒 1. sync.Mutex（排他ロック）

複数の goroutine が同じ変数にアクセスする時に、**同時書き込み・読み書き競合**を防ぐために使います。

```go
package main

import (
	"fmt"
	"sync"
)

var cnt int
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()
	cnt++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final Count:", cnt)
}
```

# 🔄 2. sync.RWMutex（読み書き分離のロック）

- `RLock`：読み取り専用ロック（複数 goroutine 可）
- `Lock`：書き込み排他ロック（他の読み書きと共存不可）

```go
var rwMu sync.RWMutex
var data = make(map[string]string)

func read(key string) string {
    rwMu.RLock()
    defer rwMu.RUnlock()
    return data[key]
}

func write(key, value string) {
    rwMu.Lock()
    defer rwMu.Unlock()
    data[key] = value
}
```

🔑 読み込みが頻繁で、書き込みが少ない時に有効。

# ⏳ 3. sync.WaitGroup（goroutine の完了を待つ）

`Add(n)`、`Done()`、`Wait()` の 3 セットで使います。

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done()
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
	fmt.Println("Add done")
}
```

✅ `Done()` を忘れると `Wait()` でずっと待機してしまうので注意！

## Output

```sh
$ go run .
Worker 3 done
Worker 1 done
Worker 2 done
Add done

$ go run .
Worker 3 done
Worker 2 done
Worker 1 done
Add done
```

# ☝️ 4. sync.Once（一度だけ実行）

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initSomething() {
	fmt.Println("This runs only once")
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(initSomething)
		}()
	}

	time.Sleep(1 * time.Second) // 実行完了待ち
}
```

🔄 `once.Do(func) `は並行実行中でも一度だけ関数を呼び出す。初期化処理などで重宝。

# ✅ まとめ

| 機能             | 説明                                       |
| ---------------- | ------------------------------------------ |
| `sync.Mutex`     | シンプルな排他ロック                       |
| `sync.RWMutex`   | 読み取り専用と書き込み排他を分離して使える |
| `sync.WaitGroup` | 複数の goroutine の終了を待機できる        |
| `sync.Once`      | 一度だけ実行したい初期化などに便利         |
