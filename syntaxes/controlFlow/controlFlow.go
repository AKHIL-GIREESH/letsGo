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

	// If with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Switch case example
	fmt.Println("\nSwitch case example:")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Mid week")
	}

	// Switch without expression (like if-else)
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

	// For loop examples
	fmt.Println("\nFor loop examples:")

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

	// For range loop with slice
	fmt.Println("\nFor range loop:")
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// For range loop with map
	fmt.Println("\nFor range loop with map:")
	person := map[string]string{
		"name": "John",
		"age":  "25",
		"city": "New York",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
