package fn

import (
	"fmt"
	"os"
)

type Bill struct {
	Name  string
	items map[string]float64
	tip   float64
}

// make new bills
func NewBill(name string) Bill {
	b := Bill{
		Name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b Bill) Format() string {
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
func (b *Bill) UpdateTip(tip float64) {
	(*b).tip = tip
}

// Add an item to the bill
func (b *Bill) AddItem(name string, price float64) {
	(*b).items[name] = price
}

// Save bill
func (b *Bill) Save() {
	dir := "bills"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	data := []byte(b.Format())
	filePath := dir + "/" + b.Name + ".txt"

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file:", filePath)
}