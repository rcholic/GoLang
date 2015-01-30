package main
import (
"fmt"
)


func myFunc(x, y int) (z, g int) { 
	z = x + 1
	g = y + 10

	return
}

func singleFunc(x int) (int) {
	z := x * x
	return z
}

func main() {
	sum1, sum2 := myFunc(1, 2)
	powerr := singleFunc(10)
	fmt.Println("myFunc returns: ", sum1)
	fmt.Println("myFunc returns: ", sum2)
	fmt.Println(myFunc(2, 4))
	fmt.Println(powerr)


}
