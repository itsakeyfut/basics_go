package user_input

import (
	"bufio"
	"fmt"
	"os"
	"playground/fn"
	"strconv"
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
	
	for {
		opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

		switch opt {
		case "a":
			for {
				name, _ := getInput("Item name: ", reader)
				price, _ := getInput("Item price: ", reader)
		
				p, err := strconv.ParseFloat(price, 64)
				if err != nil {
					fmt.Println("The price must be a number")
					continue
				}
				b.AddItem(name, p)
		
				fmt.Println("Item added - ", name, price)
				break
			}

		case "t":
			tip, _ := getInput("Enter tip amount ($): ", reader)
	
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				continue
			}
			b.UpdateTip(t)
	
			fmt.Println("Tip added - ", tip)

		case "s":
			fmt.Println("You chose to save the bill", b)
			b.Save()
			return

		default:
			fmt.Println("That was not a valid option...")

		}
	}
}