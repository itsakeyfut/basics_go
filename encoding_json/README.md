# ✅ 1. JSON を構造体に読み込む（Unmarshal）

## ◾ JSON から Go 構造体へ変換

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	jsonData := `{"id": 1, "name": "Tim", "email": "eim123@example.com"}`

	var user User
	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		panic(err)
	}
	fmt.Println("Struct:", user)
}
```

- `json.Unmarshal` は `[]byte` からデコード。
- `json:"name"` のタグでフィールド名を JSON と一致させる。

## ✅ 2. Go 構造体を JSON に変換（Marshal）

```go
func main() {
	jsonData := `{"id": 1, "name": "Tim", "email": "eim123@example.com"}`

	var user User
	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		panic(err)
	}
	fmt.Println("Struct:", user)
}
```

- `json.Marshal` は構造体 → `[]byte` に変換（= JSON）

## ◾ 見やすい整形された JSON（MarshalIndent）

```go
formatted, _ := json.MarshalIndent(user, "", " ")
fmt.Println(string(formatted))
```

出力：

```json
{
  "id": 1,
  "name": "Tim",
  "email": "tim123@example.com"
}
```

## ✅ 3. ネストした JSON

```go
type Profile struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type Account struct {
    ID      int     `json:"id"`
    Profile Profile `json:"profile"`
}
```

ネストも自動的にパースされます。

## ✅ 4. マップ形式で動的にパース

フィールドが決まっていないときや動的データには `map[string]interface{}` を使います。

```go
func main() {
    jsonData := `{"id": 1, "name": "Tim", "tags": ["go", "json"]}`

    var data map[string]interface{}
    json.Unmarshal([]byte(jsonData), &data)

    fmt.Println("ID:", data["id"])
    fmt.Println("Tag:", data["tags"])
}
```

- 数値は `float64`、文字列は `string`、配列は `[]interface{}` になるのでキャストが必要。

## ✅ 5. ファイル読み書きでの JSON 操作

### ◾ 読み込み

```go
bytes, _ := os.ReadFile("user.json")
json.Unmarshal(bytes, &user)
```

### ◾ 書き込み

```go
bytes, _ := json.MarshalIndent(user, "", " ")
os.WriteFile("user.json", bytes, 0644)
```

## ✅ 6. JSON 配列の扱い

```go
jsonArray := `[{"id":1,"name":"A"}, {"id":2,"name":'B"}]`

var user []User
json.Unmarshal([]byte(jsonArray), &users)

for _, u := range users {
    fmt.Println(u.Name)
}
```

## ✅ エラー処理のポイント

- `json.Unmarshal` や `json.Marshal` はエラーを返す → `if err != nil` チェック必須
- 構造体タグが間違ってるとパースされない

## ✅ まとめ

| 処理                 | メソッド                 | 使い方                           |
| -------------------- | ------------------------ | -------------------------------- |
| JSON → 構造体        | `json.Unmarshal`         | `json.Unmarshal([]byte, &v)`     |
| 構造体 → JSON        | `json.Marshal`           | `json.Marshal(v)`                |
| 整形された JSON 出力 | `json.MarshalIndent`     | `json.MarshalIndent(v, "", " ")` |
| 動的データ読み込み   | `map[string]interface{}` | 構造が定まらないときに便利       |
