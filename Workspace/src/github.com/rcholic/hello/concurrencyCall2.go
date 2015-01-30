package main

import "fmt"

func sayName(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("at ", i, " :", name)
	}
}


func main() {
	var name1, name2 string = "John", "Tony"
	sayName("begin")
	fmt.Println("now printing names: ", name1, name2)
	go sayName("John")
	sayName("tony")
}
