package main

import "fmt"

func main() {
    var myArr [5]int

    for i := 0; i < 5; i++ {
        myArr[i] = i
        fmt.Printf("i is %d\n", i)
    }
}
