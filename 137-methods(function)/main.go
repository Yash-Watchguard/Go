package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) happybirthday() {
	p.age++

}
func main() {
	p1 := person{name: "yash", age: 20}

	p1.happybirthday()
	fmt.Println(p1.age)
}
