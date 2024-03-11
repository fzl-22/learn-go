package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning, %s\n", name)
}

func sayBye(name string) {
	fmt.Printf("Good bye, %s\n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, value := range names {
		f(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	names := []string{"cloud", "tifa", "baret", "sephiroth"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	area1 := circleArea(10.5)
	area2 := circleArea(15)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", area1, area2)
}
