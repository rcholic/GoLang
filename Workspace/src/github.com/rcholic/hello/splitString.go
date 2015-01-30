package main
import "fmt"
import "strings"

func main() {

	var aSentence string= "Hello World, how are   you??"
	words := strings.Fields(aSentence)

	newSlice := words[2:4] //can slice the array

	for _, word := range words {

		fmt.Printf("the word is %s\n", word)
		fmt.Printf("length of the newSlice is: %d\n", len(newSlice))
	}


}
