package main

import (
	"fmt"
	"math"
)

type Animal struct {
	name string
}

type Human struct {
	*Animal
}

func (a *Animal) speak() string {
	return fmt.Sprintf("%s speakings %.3f\n", a.name, math.Pi)
}

func (h *Human) speak() string {
	return fmt.Sprintf("human speaking\n")
}

func main() {
	aCat := &Animal{"Cat"}
	fmt.Println(aCat.speak())

	aHuman := &Human{aCat}
	fmt.Println(aHuman.speak())
}
