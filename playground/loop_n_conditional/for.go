package loop_n_conditional

import "fmt"

func ForLoop() {
	// Go does not support while loop, but try emulating
	fmt.Println("----- mimicking while loop -----")
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	fmt.Println("----- for loop -----")
	for i := 0; i < 5; i++ {
		fmt.Println("value of x is:", x)
	}

	fmt.Println("----- Every element loop -----")
	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println("value of i is:", names[i])
	}

	for idx, val := range names {
		fmt.Printf("The position %v is %v\n", idx, val)
	}

	for _, val := range names {
		fmt.Printf("The value is %v\n", val)
	}
}