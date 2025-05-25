package main

import "fmt"

func main() {

	// If-else example
	age := 18
	fmt.Println("If-else example:")
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Switch example
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 80:
		fmt.Println("Good!")
	case score >= 70:
		fmt.Println("Average")
	default:
		fmt.Println("Need improvement")
	}

	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// While-like for loop
	fmt.Println("\nWhile-like for loop:")
	count := 1
	for count <= 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}
}
