package main
import "fmt"


//could aso be X, Y int
type Rect struct {
	X int
	Y int
}


func main() {
	var myNum int = 100;

	var myRect  = Rect{10, 20}
	fmt.Println("myRect.x = ", myRect.X)
	fmt.Println("myNum is: ", myNum)
}
