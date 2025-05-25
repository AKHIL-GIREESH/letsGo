package main

import "fmt"

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	slice := numbers[1:3]

	slice = append(slice, 6, 7, 8, 9, 10)

	fmt.Println(cap(slice), len(slice))
	fmt.Println(cap(numbers), len(numbers), numbers[4])

}
