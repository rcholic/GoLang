package main
import "fmt"

func main() {
	
	container := make([]int, 102)

	for i := 0; i < 100; i++ {
		container[i] = i * i
	
	}

	fmt.Println(container)

}
