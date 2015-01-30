package main
import "fmt"
import "strconv"


func main() {
	i, err := strconv.Atoi("42")

	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	}

	fmt.Printf("Converted integer: %d\n", i)

}
