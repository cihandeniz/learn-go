package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}       // has type Vertex
	v2 = Vertex{X: 1}       // Y:0 is implicit
	v3 = Vertex{}           // X:0 and Y:0
	v4 = Vertex{Y: 1, X: 2} // reverse order
	p  = &Vertex{1, 2}      // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, v4, p)
}
