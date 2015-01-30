package main
import "fmt"

type myType struct {
	mySlice []int
	length int
}

func main() {

	aType := myType{make([]int, 10), 10}

	fmt.Printf("aType.mySlice.cap is: %d, and aType.length = %d\n", cap(aType.mySlice), aType.length)

}
