package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 120
	s := strconv.Itoa(x)
	fmt.Printf("%s\n", s)

	y, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("Inc %d\n", y+1)
	}
}
