package main

import "fmt"
import "time"

//channels can take a buffer of values
func main() {
	messages := make(chan string, 3)

	messages <- "buffered"
	messages <- "channel"
	messages <- "thirdvalue"
	fmt.Println("lengthis is: ", len(messages))

	fmt.Println(<-messages)
	fmt.Println("lengthis is: ", len(messages))
	fmt.Println(<-messages)

	time.Sleep(2* time.Second)
	fmt.Println(<-messages)

	fmt.Println("lengthis is: ", len(messages))
}
