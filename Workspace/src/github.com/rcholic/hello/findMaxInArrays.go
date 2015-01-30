package main
import "fmt"

func Max(slice []int) int {
	max := slice[0] //take the first element as max for now

	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}

	return max
}

func main() {
	A1 := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 10, 100}
	A2 := [5]int {-100, 99, -99, 98, 101}
	A3 := [3]int {56, 89, 103}
	A4 := A3[:]
	A4 = append(A1[:], A2[len(A2)-1])

	A5 := make([]int, 2, 2)
	A5 = append(A5, 1, 2, 3, 4, 5)
	fmt.Println("A5 is: ", A5) // why this output [0, 0, 1, 2, 3, 4, 5] ?

	maxNum := Max(A4)
	fmt.Println("max number in A4 is ", maxNum)
	fmt.Println("A4 is ", A4)

	//not using index, so replace it with underscore _
	for _, val := range A1 {
		fmt.Println("A1 is: ", val)
	}

	for index, _ := range A2 {
		fmt.Printf("A2[%d] = %d\n", index, A2[index])
	}

}
