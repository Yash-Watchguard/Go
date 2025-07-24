package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("enter x under 1 to 4\n")
	fmt.Scanln(&x)

	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("y")
	}
}
