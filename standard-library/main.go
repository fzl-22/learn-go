package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings package
	greeting := "hello there friend!"

	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("original string value =", greeting)

	// sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 65)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)

	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
