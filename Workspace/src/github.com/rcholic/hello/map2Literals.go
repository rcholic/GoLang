package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Micro International": {100.2934, -129.9923},
	"Google" : {36.4222, -122.0867},
}

func main() {

	/*
	var m = map[string]Vertex{
		"Micro International": Vertex{100.2934, -129.9923,},
		"Google" : Vertex{36.4222, -122.0867,},
	}
	*/


	fmt.Println(m)

}

