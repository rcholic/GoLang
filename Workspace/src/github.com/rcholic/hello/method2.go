package main
import (
	"fmt"
	)

type Person struct {
	name, location string
	age int
}

func (p *Person) toString() (string, string, int) {
	
	return p.name, p.location, p.age
	
	//return ("name is %s, location is %s, and age is %d\", p.name, p.location, p.age)
}

func main() {
	var p = Person{"Tony", "Atlanta", 12}
	
	fmt.Println(p.toString())
	
	//fmt.Printf("name is %s, location is %s, and age is %d\n", p.name, p.location, p.age)

}
