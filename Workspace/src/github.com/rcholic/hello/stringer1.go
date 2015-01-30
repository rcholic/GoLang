package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) toString() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 32}
	z := Person{"Zaphod Beeblebrox", 90}
	fmt.Println(a, z)
}
