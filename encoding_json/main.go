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