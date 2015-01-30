package main
import "fmt"

func main() {

var a [5]int

for i := 0; i < 5; i++ {
	a[i] = i
	fmt.Println(i)
}

	fmt.Printf("length of a is: %d\n", len(a))
	fmt.Println(a)
}
