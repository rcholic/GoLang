package main
import "fmt"

func main() {

	p := []int{2, 3, 5, 10}
	fmt.Println("p = ", p)
	fmt.Println("p[1:3]=", p[1:3])

	//missing low index implies 0
	fmt.Println("p[:2] = ", p[:2])

	//missing high index implies len(s)
	fmt.Println("p[0:] = ", p[0:])
}
