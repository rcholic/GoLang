package main

import ("fmt")

func main() {

    var a [5]uint

    for i := 0; i < 5; i++ {
        a[i] = uint(i)
        fmt.Println(i)
    }

}
