package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"pages"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, err := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Println(err)

	intB, _ := json.Marshal(10)
	fmt.Println(intB)
}
