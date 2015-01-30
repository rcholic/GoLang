package main
import "fmt"

var c, python, java bool
var i, j = 1, -10

func main() {
	k, m := 100, -20
	fmt.Println(i, c, python, java)

	if !c {
		fmt.Println("c is false")
	}

	fmt.Printf("i = %d, j = %d\n", i, j)
	fmt.Printf("k = %d, m = %d\n", k, m)

}
