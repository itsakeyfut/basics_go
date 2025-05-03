package variables_n_string

import "fmt"

func Strs() {
	var firstName string = "Lewis"
	var lastName string = "Arsenault"
	var middleName string
	fmt.Println(firstName + " " + lastName)

	middleName = "PlayStation"
	fmt.Println(firstName + " " + middleName + " " + lastName)

	position := "lead engineer"
	fmt.Println(firstName + " " + middleName + " " + lastName + "\nPosition: " + position)
}