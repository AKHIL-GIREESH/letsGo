package main

import "fmt"

func Slice2() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = append(s, 8, 9)
	fmt.Println(s)

	s1 := make([]int, 0, 5)
	fmt.Println(s1, len(s1), cap(s1))

}
