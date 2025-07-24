package main

import "fmt"

func main() {
	array := []int{1, 3, 4, 5, 5}
	// x := sum(1, 2,3,4,5)
	x := sum(array...)
	fmt.Println(x)
}

func sum(arary ...int) int {
	var total int
	for _, k := range arary {
		total += k
	}
	return total
}
