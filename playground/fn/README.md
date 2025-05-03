## 関数の型

### 引数なし、戻り値なし

```go
func Hello() {
    fmt.Println("Hello")
}
```

### 引数あり、戻り値なし

```go
func Greet(name string) {
    fmt.Println("Hello,", name)
}
```

### 引数なし、戻り値あり

```go
func GetNumber() int {
    return 42
}
```

### 引数あり、戻り値あり

```go
func Add(a int, b int) int {
    return a + b
}
```

### 複数戻り値

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divide by zero")
    }
    return a / b, nil
}
```

## 高度な型宣言パターン

### 戻り値に名前を付ける（named return values）

```go
func Split(sum int) (x, y int) {
    x = sum * 4 /9
    y = sum - x
    return // x, y が返る
}
```

### 可変長引数（variadic parameter）

```go
func Sum(num ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### 関数を引数に取る

```go
func Apply(f func(int) int, x int) int {
    return f(x)
}
```

### 関数を返す

```go
func MakeAddr(base int) func(int) int {
    return func(x int) int {
        return base + x
    }
}
```

### 無名関数

```go
func main() {
    func(msg string) {
        fmt.Println(msg)
    }{"Hello"}
}
```

## 構造体やメソッドと組み合わせた関数

### メソッドの型宣言

```go
type User struct {
    Name string
}

func (u User) Greet() string {
    return "Hello, " + u.Name
}
```

### ポインタレシーバのメソッド

```go
func (u *User) Rename(newName string) {
    u.Name = newName
}
```

## 関数型

```go
type Operator func(int, int) int

func Add(a, b int) int {
    return a + b
}
```
