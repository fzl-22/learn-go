package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// receiver function
// format the bill (can only be called from bill object)
// pass a pointer to prevent the creation of object's copy
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25s ...$%v\n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25s ...$%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25s ...$%0.2f\n", "total:", total+b.tip)
	return fs
}

// update tip (to change object value, we must pass a pointer)
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip
	// can be automatically dereferenced
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bills was saved to file")
}
