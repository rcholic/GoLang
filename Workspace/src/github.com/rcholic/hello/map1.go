package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	//var m map[string]Vertex

	var m = make(map[string]Vertex)
	m["International"] = Vertex{40.6843, -75.39987}

	fmt.Println(m["International"].Lat, m["International"].Long)
}

