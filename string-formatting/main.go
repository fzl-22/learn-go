package main

import "fmt"

func main() {
	age := 35
	name := "shaun"
	// Print
	fmt.Print("hello, ")
	fmt.Print("world!\n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goobye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings with format specifiers)
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	// note: %q must be used on a string
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %f points!\n", 225.55)
	fmt.Printf("you scored %0.2f points!\n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("The saved string is:", str)
}
