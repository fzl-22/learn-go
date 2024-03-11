package main

import "fmt"

func updateName(x *string) {
	fmt.Println("memory address of x is:", &x)
	*x = "wedge"
}

func main() {
	// Non-Pointer Values -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	// fmt.Println("memory address of name is:", &name)

	// var nameAddress *string = &name
	nameAddress := &name
	fmt.Printf("memory address\t\t: %p\n", nameAddress)
	fmt.Printf("value at address\t: %s\n", *nameAddress) // dereference

	fmt.Println(name)

	updateName(nameAddress)

	fmt.Println(name)
}
