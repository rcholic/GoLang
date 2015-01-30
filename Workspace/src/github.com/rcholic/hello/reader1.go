package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello, reader!")

	//bytes representation of letters (e.g. 104 for h)
	b := make([]byte, 8)

	for {
		n, err := r.Read(b) //read the string in r to the slice 'b'
		fmt.Printf("n = %v err = %v b=%v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
