package main

import (
    "fmt"
    "time"
)

func Postcard(from string) {
    fmt.Println("hello from ", from)
}

func main() {
    Postcard("main")

    go Postcard("A goroutine")

    go func() {
        fmt.Println("A closure")
    }()

    time.Sleep(1 * time.Second)
}
