package main

import "fmt"

type Struct struct {
	X int
	Y int
}

func StructsnPointers() {
	v := Struct{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
