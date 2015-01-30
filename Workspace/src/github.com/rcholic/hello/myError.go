package main
import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

//define the Error method on the struct 'MyError'
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError {
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

}

