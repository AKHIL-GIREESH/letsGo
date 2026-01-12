package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v *Vertex) MethodScale(magnitude int) {
	v.X *= magnitude
	v.Y *= magnitude
}

func FuncScale(v *Vertex, magnitude int) {
	v.X *= magnitude
	v.Y *= magnitude
}

func MethodsVFunctions() {
	v1 := Vertex{
		5, 10,
	}

	v1.MethodScale(3)
	fmt.Println(v1)

	v2 := Vertex{
		5, 10,
	}

	FuncScale(&v2, 3)
	fmt.Println(v2)

}
