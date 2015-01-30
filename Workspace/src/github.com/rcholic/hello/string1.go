package main
import (
	"fmt"
	
)

var myName = "Ivytony"

func main() {
	
	for i, v := range myName {
	
		fmt.Printf("at index of %d, the value is: %c\n", i, v)
	}

}
