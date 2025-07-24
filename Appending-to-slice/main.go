package main

import (
	"fmt"
)

func main() {

	// if capacity is less then it double the capacity
	x := make([]int, 3, 5)
	x[0], x[1], x[2] = 1, 2, 3

	x = append(x, 4, 5, 6,7,2,23,4,5,33,4,67,7)

	fmt.Println(cap(x))
}
