package main

import (
	"fmt"
)

type human interface {
	speak()
}

type person struct {
	name string
}

type seceretagent struct {
	person
	work string
}

func (obj person) speak() {
	fmt.Println("hi i am ", obj.name)
}
func (obj seceretagent) speak() {
	fmt.Println("hi i am ", obj.name)
}

func saysomethng(obj human) {
	obj.speak()
}
func main() {
	p1 := person{
		name: "yash",
	}

	secret_agent := seceretagent{
		person: person{name: "zaid"},
		work:   "spy",
	}
	// concreate implementation and Polymorphism
	// p1.speak()
	// secret_agent.speak()

	// \by client code using interface

	saysomethng(p1)
	saysomethng(secret_agent)

}
