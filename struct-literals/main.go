package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	x, y := fmt.Println(v1, p, v2, v3)
	fmt.Printf("%T", p)
	fmt.Println(x, y)
	z := []int{1, 2, 3, 4, 5}
	n, err := fmt.Println(z) // print the slice and capture return values
	fmt.Println(n, err)
}
