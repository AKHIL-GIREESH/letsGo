package main

import "fmt"

func main() {
	var firstname string
	var lastname string
	fmt.Print("Enter your name : ")
	fmt.Scanf("%s %s", &firstname, &lastname)
	fmt.Printf("Hello %s %s\n", firstname, lastname)
}
