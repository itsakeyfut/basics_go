package user_input

import (
	"bufio"
	"fmt"
	"os"
	"playground/fn"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func CreateBill() fn.Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := fn.NewBill(name)
	fmt.Println("Created the bill - ", b.Name)

	return b
}

func PromptOptions(b fn.Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}