package main

import "fmt"

func main() {
	// for (in while-style)
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	// for (in regular for-style)
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Printf("%s\n", names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %d is %s\n", index, value)
	}

	// replace with _ if not needed
	for _, value := range names {
		fmt.Printf("the value is %s\n", value)
		// doesn't modify the original slice,
		// because it's a local variable
		value = "new string"
		// this will modify the original slice
		// names[index] = fmt.Sprintf("new string %d", index)
	}

	fmt.Println(names)
}
