package main

import (
	"fmt"
)

func main() {
	type person struct {
		first string
		last  string
		age   int
	}

	// var p1 person
	// var p2 person

	// p1.first = "yash"
	// p1.last = "goyal"
	// p1.age = 20

	// p2.first = "yash"
	// p2.last = "goyal"
	// p2.age = 20

	p1 := person{
		first: "yash",
		last:  "goyal",
		age:   18,
	}
	p2 := person{
		first: "rekha",
		last:  "goyal",
		age:   25,
	}
	fmt.Println(p1, p2)

	fmt.Printf("%T \t %#v", p1, p1)
}
