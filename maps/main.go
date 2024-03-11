package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            5.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		2675831341: "mario",
		8734810871: "luigi",
		3668314379: "peach",
	}

	fmt.Println(phonebook[2675831341])

	phonebook[3668314379] = "bowser"
	phonebook[8713697973] = "zielisme"

	for key, value := range phonebook {
		fmt.Printf("%d - %s\n", key, value)
	}
}
