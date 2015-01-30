package main

import "fmt"

func main() {
	var a [5]int

	for i := 0; i < 5; i++ {
		a[i] = i
	}

	fmt.Printf("this is an array %g\n", a)

	//declare a slice
	//var letters []byte

	//create the slice with value and capacity
	//letters := make([]byte, 5, 5)

	letters := []byte{'a', 'b', 'c'}

	fmt.Printf("length is: %d\n", len(letters))
}
