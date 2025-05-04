package fn

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func NewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) Format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k + ":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

	return fs
}

// Update tip
func (b *bill) UpdateTip(tip float64) {
	(*b).tip = tip
}

// Add an item to the bill
func (b *bill) AddItem(name string, price float64) {
	(*b).items[name] = price
}