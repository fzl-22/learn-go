package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)

	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	// the other value is nothing
	return initials[0], "_"
}

func main() {
	firstName1, secondName1 := getInitials("tifa lockhart")
	fmt.Println(firstName1, secondName1)

	firstName2, secondName2 := getInitials("cloud strife")
	fmt.Println(firstName2, secondName2)

	firstName3, secondName3 := getInitials("barret")
	fmt.Println(firstName3, secondName3)
}
