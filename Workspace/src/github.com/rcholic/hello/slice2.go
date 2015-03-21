package main
import "fmt"

func main() {
    
    container := make([]string, 5)

    for i := 0; i< 5; i++ {
        container[i] = "hello world"
    }

    container2 := make([]string, 2)
    container2 = append(container[1:3])

    fmt.Println("container is: ", container)
    fmt.Println("container2 is: ", container2)
}
