package fn

import "fmt"

func Maps() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string {
		32525234: "mario",
		32543583: "luigi",
		92847534: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[32525234])

	phonebook[498326495] = "bowser"
	fmt.Println(phonebook)

	phonebook[92847534] = "yoshi"
	fmt.Println(phonebook)
}