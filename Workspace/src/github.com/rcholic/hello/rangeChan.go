package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "tow"
	close(queue) //close the channel, but it can still send the values

	//closed channel can still send the data
	/*
		for elem := range queue {
			fmt.Println(elem)
		}
	*/

	for len(queue) > 0 {
		fmt.Println(<-queue)
	}

}
