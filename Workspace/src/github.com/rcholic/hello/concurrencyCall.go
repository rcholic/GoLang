package main

import "fmt"

func f(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	go fmt.Println("type an input below: ")

	fmt.Scanln(&input)
	fmt.Printf("You have input this: %s\n", input)
	fmt.Println("done")


}
