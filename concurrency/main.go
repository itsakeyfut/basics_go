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