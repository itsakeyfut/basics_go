package main

import (
	"fmt"
	"playground/fn"
)

func main() {
	// // strings
	// variables_n_string.Strs()

	// // ints
	// variables_n_string.Ints()

	// // bits & memory
	// variables_n_string.BitAndMemory()

	// // print
	// variables_n_string.Prints()

	// // array & slices
	// variables_n_string.ArrayAndSlices()

	// // Standard Library
	// variables_n_string.Stl()

	// // Loop
	// loop_n_conditional.ForLoop()

	// Logical Operators and continue and break keyword in for looping
	// loop_n_conditional.LogicalOperators()

	// // Function
	// fn.CycleNames([]string{"cloud", "tifa", "barret"}, fn.SayGreeting)
	// fn.CycleNames([]string{"cloud", "tifa", "barret"}, fn.SayBye)

	// a1 := fn.CircleArea(10.5)
	// a2 := fn.CircleArea(15)
	// fmt.Println(a1, a2)
	// fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

	// // Multiple Functions
	// org, upper := fn.GetInitials("Akari Akaza")
	// fmt.Printf("%v is upper to %v \n", org, upper)

	// org2, upper2 := fn.GetInitials("Mirakurun")
	// fmt.Printf("%v is upper to %v \n", org2, upper2)

	// // Map
	// fn.Maps()

	// // Pass by value
	// name := "Akari"
	// fmt.Println(name)
	// fn.UpdateName(name)
	// fmt.Println(name)

	// // Pointer
	// name := "Akari"
	// fn.UpdateRefName(&name)
	// fmt.Println("memory address of name is:", &name)

	// refName := &name
	// fmt.Println("memory address:", refName)
	// fmt.Println("value at memory address", *refName)

	// fmt.Println(name)

	// // Structs
	// myBill := structs.NewBill("mario")
	// fmt.Println(myBill)

	// Receiver Functions
	myBill := fn.NewBill("mario")
	myBill.AddItem("curry", 4.25)
	myBill.AddItem("toffee pudding", 6.99)
	myBill.UpdateTip(10)

	fmt.Println(myBill.Format())
}