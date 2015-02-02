package main

import "fmt"

type node struct {
    parent *node
    left *node
    right *node
    val string
}

func main() {
    
    firstNode := &node{nil, nil, nil, "first"}
    parentNode := &node{firstNode, nil, nil, "parent"}
    firstNode.parent = parentNode

    fmt.Printf("does it work?, parentNode value is %v\n", parentNode.val)
}
