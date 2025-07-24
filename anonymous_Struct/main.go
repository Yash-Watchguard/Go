package main

import "fmt"

func main() {
	person := struct {
		first string
		last  string
		age   int
	}{
		first: "yash",
		last:  "goyal",
		age:   20,
	}

	fmt.Println(person.first)
}
