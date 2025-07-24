package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	Name    string
	Age     int
	Address // Embedded struct (no field name = anonymous field)
}

func main() {
	p := Person{
		Name: "Yash",
		Age:  22,
		Address: Address{
			City:  "Delhi",
			State: "Delhi",
		},
	}

	// Accessing embedded fields directly
	fmt.Println(p.Name)  // Yash
	fmt.Println(p.City)  // Delhi
	fmt.Println(p.State) // Delhi
}
