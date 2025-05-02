package main

import "fmt"

func main() {
	// strings
	var firstName string = "Lewis"
	var lastName string = "Arsenault"
	var middleName string
	fmt.Println(firstName + " " + lastName)

	middleName = "PlayStation"
	fmt.Println(firstName + " " + middleName + " " + lastName)

	position := "lead engineer"
	fmt.Println(firstName + " " + middleName + " " + lastName + "\nPosition: " + position)

	// ints
	var num1 int = 20
	var num2 = 30
	num3 := 40
	fmt.Println("Sum: ", num1 + num2 + num3)

	// bits & memory
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