package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings
	strs()

	// ints
	ints()

	// bits & memory
	bits_n_memory()

	// print
	prints()

	// array & slices
	array_n_slices()

	// Standard Library
	stl()
}	

func strs() {
	var firstName string = "Lewis"
	var lastName string = "Arsenault"
	var middleName string
	fmt.Println(firstName + " " + lastName)

	middleName = "PlayStation"
	fmt.Println(firstName + " " + middleName + " " + lastName)

	position := "lead engineer"
	fmt.Println(firstName + " " + middleName + " " + lastName + "\nPosition: " + position)
}
func ints() {
	var num1 int = 20
	var num2 = 30
	num3 := 40
	fmt.Println("Sum: ", num1 + num2 + num3)
}
func bits_n_memory() {
	var num4 int8 = 0x01
	var num5 int64 = 0xFFFF
	var num6 uint32 = 127
	var num7 int32 = -128
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)

	var fnum1 float32 = 25.98
	var fnum2 float64 = 6555455544.44444
	fmt.Println(fnum1)
	fmt.Println(fnum2)

	var cnum1 complex64 = complex(1.5, 2.5)
	var cnum2 complex128 = complex(3.1, -4.2)
	fmt.Println(cnum1)
	fmt.Println(cnum2)

	c := complex(2.0, 3.0)
	fmt.Println("Real:", real(c))
	fmt.Println("Real:", imag(c))

	var bnum byte = 0x11
	var rnum rune = -65536
	fmt.Println(bnum)
	fmt.Println(rnum)
}
func prints() {
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
func array_n_slices() {
	var ageArr1 [3]int = [3]int{20, 25, 30}
	var ageArr2 = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ageArr1, len(ageArr1))
	fmt.Println(ageArr2, len(ageArr2))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 1999
	scores = append(scores, 3294)
	fmt.Println(scores)

	// slice ranges
	r1 := names[1:3]
	r2 := names[2:]
	r3 := names[:3]
	fmt.Println(r1, r2, r3)
}
func stl() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Index(greeting, "t"))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))


	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	ages := []int{46, 43, 65, 23, 2, 34, 12, 56, 27}

	sort.Ints(ages)
	fmt.Println(ages)

	idx := sort.SearchInts(ages, 43)
	fmt.Println(idx)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}