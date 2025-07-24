package main

import (
	"fmt"
)

func main() {
	array := [4]int{1, 2, 34, 4}

	// we can make a slice by using a array

	x := array[1:3]
	fmt.Println(x)
	/*
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	*/
}
