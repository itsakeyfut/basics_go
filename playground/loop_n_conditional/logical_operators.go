package loop_n_conditional

import "fmt"

func LogicalOperators() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less thyan 45")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for idx, val := range names {
		if idx == 1 {
			fmt.Println("continuing at pos", idx)
			continue
		}
		if idx > 2 {
			fmt.Println("breaking at pos", idx)
			break
		}

		fmt.Printf("The value of pos %v is %v\n", idx, val)
	}
}