package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}() //go routine on a function closure

	time.Sleep(time.Millisecond * 1500) //main() sleeps 1500 ms
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
