package main

import "fmt"

func main() {
	names := [4]string{
		"A",
		"B",
		"C",
		"D",
	}

	fmt.Println(names)

	a := names[0:2] // length=2, capacity=4, pointer="A"
	b := names[1:3] // length=2, capacity=3, pointer="B"
	fmt.Println(a, b)

	b[0] = "X"
	fmt.Println(a, b, names)

}
