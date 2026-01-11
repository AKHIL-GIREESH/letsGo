package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1
	fmt.Println(m["a"])

	m["a"] = 2
	fmt.Println(m["a"])

	delete(m, "a")
	fmt.Println(m["a"])

	v, ok := m["a"]
	fmt.Println(v, ok)

}
