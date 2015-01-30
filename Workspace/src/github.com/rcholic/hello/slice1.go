package main
import (
"fmt"
"math"
)

func main() {
	p := []int{2, 3, 4, 5}
	fmt.Println("p == " , p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, math.Pow(float64(p[i]), 2))
	}
}
