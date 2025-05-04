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