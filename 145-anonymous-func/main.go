package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("hi")
	}()

	func(s string) {
		fmt.Println("my name is:", s)
	}("yash")
	// Assigning an anonymous function to a variable
	greet := func(s string) {
		fmt.Println("hi", s)
	}
	(greet("yash"))
	greet("bob")

	count := 0

	increment := func() int {
		count++
		return count
	}

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

/*
 Anonymous function assigned to a variable
When you assign an anonymous function to a variable, like:

go
Copy code
x := func() {
    fmt.Println("Hello")
}
You’re only defining the function here.

There is no () at the end because you’re not calling it immediately — you’re just storing the function in x.
*/
