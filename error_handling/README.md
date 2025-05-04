# Erro Handling

Go では「関数から戻る値のひとつ」としてエラーが扱われます。

例：ファイルを開く

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("hello.txt")
    if err != nil {
        // エラーがある場合の処理
        fmt.Println("エラー:", err)
        return
    }

    // ファイルが正常に開けたとき
    fmt.Println("ファイルを開けました:", file.Name())
    file.Close()
}
```

## ✅ ポイント

- op.Open の戻り値は (file \*os.File, err error)
- err != nil でエラーが起きたか確認
- 明示的に return または log.Fatal(err) などで中断するのが一般的

## 関数でエラーを返すパターン

```go
package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
```

## 🔧 よく使う error の作成方法

| 方法                          | 用途                             |
| ----------------------------- | -------------------------------- |
| `errors.New("メッセージ")`    | シンプルなエラーを作る           |
| `fmt.Errorf("詳細: %w", err)` | エラーラップ（元のエラーを内包） |
| `errors.Is(err, targetErr)`   | エラーの種類を比較               |
| `errors.As(err, &targetType)` | エラー型を抽出（型アサーション） |

```go
package main

import (
    "errors"
    "fmt"
)

var ErrCustom = errors.New("Custom Error")

func run() error {
    return fmt.Errorf("Failed while running: %w", ErrCustom)
}

func main() {
    err := run()
    if errors.Is(err, ErrCustom) {
        fmt.Println("Custom Error occurred:", err)
    }
}
```

## 🧼 エラーの取り扱いのベストプラクティス

1. err != nil は常にチェックする（省略しない）
2. わかりやすいエラーメッセージを返す
3. ラップをして元のエラー情報を保持する
4. パッケージをまたぐときはエラー文言を調整する

# ✅ 1. errors.Is / errors.As の活用

Go では、エラーをラップしても「元のエラー」を比較・抽出したい場面が多いため、標準パッケージで errors.Is や errors.As が提供されています。

## ◾ errors.Is: エラーの「中身の一致」を確認

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func fetchData() error {
	return fmt.Errorf("fetch failed: %w", ErrNotFound)
}

func main() {
	err := fetchData()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Not found any data.")
	} else {
		fmt.Println("Unknown Error:", err)
	}
}
```

🔍 fmt.Errorf("...: %w", err) で「ラップされたエラー」でも errors.Is が内部までたどって一致するか判定できます。

## ◾ errors.As: 特定の型のエラーにキャスト（型アサーション）

```go
package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
}

func run() error {
	return fmt.Errorf("failed while running: %w", MyError{Code: 404, Msg: "not found"})
}

func main() {
	err := run()

	var myErr MyError
	if errors.As(err, &myErr) {
		fmt.Println("Detected custom error:", myErr.Code, myErr.Msg)
	} else {
		fmt.Println("Unexpected Error:", err)
	}
}
```

# ✅ 2. カスタムエラー型（type MyError struct {}）

## 基本形：

```go
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("[Code %d] %s", e.Code, e.Msg)
}
```

この型は `error` インターフェースを満たすため、どこでも `error` として使えます。

## よくある用途：

- HTTP エラー（404, 500 など）
- データバリデーションエラー
- アプリ固有のエラー

```go
func validate(input string) error {
	if input == "" {
		return MyError{Code: 400, Msg: "Empty input"}
	}
	return nil
}
```

# ✅ 3. サードパーティエラー処理ライブラリ：github.com/pkg/errors（※今は非推奨気味）

Go 1.13 以降、`fmt.Errorf("%w")`, `errors.Is`, `errors.As` が標準に加わったため、`pkg/errors` はあまり使われなくなりましたが、古いコードでは見かけます。

## 代表的な使い方（旧式）：

```go
import "github.com/pkg/errors"

func f() error {
	return errors.Wrap(io.EOF, "failed reading")
}

func main() {
	err := f()
	fmt.Printf("%+v\n", err) // スタックトレース付きで表示
}
```

現在は Go の標準機能だけで十分 なケースが多く、`pkg/errors` は主にスタックトレース付きエラーログ用途で一部で使用される程度です。
