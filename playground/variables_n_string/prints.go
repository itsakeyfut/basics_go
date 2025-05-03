package variables_n_string

import "fmt"

func Prints() {
	fmt.Print("hello, ")
	fmt.Print("world")
	fmt.Print("new line \n")

	fmt.Println("Hi")
	fmt.Println("Catch you later")

	age := 27
	name := "lewis"
	fmt.Println("my age is", age, "and my name is", name)

	your_age := "20"
	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v \n", your_age, name)
	fmt.Printf("my age is %q and my name is %q \n", your_age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("age is of type %T \n", name)
	fmt.Printf("you scored %f points! \n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)
}