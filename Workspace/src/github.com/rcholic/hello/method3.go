package main

import "fmt"

type Population int

type Context struct {
    header string
    route string
    status int
}

func (c Context) toString() (string, string, int) {

    return c.header, c.route, c.status
}

func (population Population) howMany() {
    fmt.Println("population is ", population)
}

func main() {
    pop := Population(100)
    pop.howMany()

    context := Context{"header1", "route1", 200}
    fmt.Println(context.toString())
}
