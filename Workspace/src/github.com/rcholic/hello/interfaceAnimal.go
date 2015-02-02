package main

import "fmt"

type Animal interface {
	Speak() string
}

//implement Animal interface
type Cat struct {
	sound string
}

//implement Animal interface
type Dog struct {
	sound string
}

func (c Cat) Speak() string {
	//	c.sound = "meow";
	return c.sound
	//	return "woof"
}

func (d Dog) Speak() string {
	//	d.sound = "bark";
	return d.sound
	//	return "meow"
}

/*
func (a Animal) Speak() string {
	fmt.Println(a.sound)
}
*/

func main() {

	var aCat = Cat{"meow"}
	var aDog = Dog{"bark"}

	fmt.Println("aCat.Speak(): ", aCat.Speak())
	fmt.Println("aDog.Speak(): ", aDog.Speak())
}
