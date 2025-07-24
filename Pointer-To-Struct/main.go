package main

import (
	"fmt"
)

type person struct {
	name    string
	sirname string
	age     int
}

func main() {
	p1 := person{
		"yash",
		"goyal",
		18,
	}

	p2 := &p1

	fmt.Println(p2.name)
	fmt.Println((*p2).age)
}
