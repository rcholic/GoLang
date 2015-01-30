package main

import "fmt"

func main() {
	slice1 := make([]int, 10, 20)

	for i := 0; i < 20; i++ {
		slice1 = append(slice1, i)
	}

	fmt.Println("slice1 is: ", slice1)

	slice2 := slice1[10:]
	fmt.Println("slice2 is: ", slice2)
}
