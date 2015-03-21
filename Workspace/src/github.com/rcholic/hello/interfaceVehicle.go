package main

import "fmt"
import "time"

type Vehicle interface {
	Run()
	getMileage()
}

//a Car Type
type Car struct {
	vehicleType string
	mileage     int
}

//implement the Run method
func (c Car) Run() string {
	return fmt.Sprintf("%v is running\n", c.vehicleType)
}

//implement the getMileage() method
func (c Car) getMileage() int {
	return c.mileage
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("msg is ", msg1)
	case msg2 := <-c2:
		fmt.Println("msg2 is: ", msg2)
	}

	aCar := Car{"Car", 1000}
	fmt.Print(aCar.Run())
	fmt.Println(aCar.getMileage())
}
