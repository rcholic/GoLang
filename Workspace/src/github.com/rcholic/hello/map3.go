package main

import "fmt"

func main() {

    myMap := make(map[int]string)
    myMap[1] = "hello world"
    myMap[2] = "hello world again"
    
    fmt.Println("length of map is: ", len(myMap))

    fmt.Println("map is: ", myMap[1])
    _, prs := myMap[3]  //not existing key

    fmt.Println("prs is: ", prs)
}
