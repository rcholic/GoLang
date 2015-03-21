package main

import (
	"fmt"
)

func main() {
	seq := make([]string, 8)

	for i := 0; i < 8; i++ {
		seq[i] = "a"
	}

	fmt.Println(seq)

}
