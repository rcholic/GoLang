package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	p := &Vertex{5, 6}
	fmt.Println(p.Abs())

	//can also define a Vertex struct directly
	var m = Vertex{3, 4}
	fmt.Println(m.Abs())
}
