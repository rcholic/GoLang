package main
import "fmt"

func main() {
	myName := "UnitedStates"

	for i := 0; i < len(myName); i++ {
		fmt.Printf("character is %c\n", myName[i])
	}

	for _,c := range myName {
	
		fmt.Printf("character is %c\n", c)
	}

}
