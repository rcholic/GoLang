package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
v1 = Vertex{1, 2}
p = &v1
)

func main() {
	fmt.Println(p)
	fmt.Printf("v1.X = %d\n", v1.X)
	fmt.Println(p.X)
}
