package main

import "fmt"
import "time"

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
        for i := 0; i < 4; i++ {
            f("goroutine-inside")
        }
	}("going")
    
    time.Sleep(2 * time.Second)    

	var input string
	go fmt.Println("type an input below: ")

	fmt.Scanln(&input)
	fmt.Printf("You have input this: %s\n", input)
	fmt.Println("done")
}
