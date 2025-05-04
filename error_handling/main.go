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