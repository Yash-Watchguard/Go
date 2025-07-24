package main

import "fmt"

// func add(a ...int) int {
// 	sum := 0

// 	for v := range a {

// 		sum += a[v]
// 	}

// 	return sum
// }

// func fact(num int) int {
// 	if num == 0 || num == 1 {
// 		return 1
// 	}
// 	fact := num * fact(num-1)
// 	return fact
// }

// func Outer(x int) func(int) int {
// 	return func(y int) int {
// 		return x * y
// 	}
// }

type person struct {
	name string
	last string
}

type child struct {
	person
	father string
}

func (c *child) born() {
	fmt.Println(c.name, c.last, c.father)
}

func (p person) add(x, y int) int {
	return x + y
}

func main() {

	p := person{
		"shyam",
		"pratap",
	}

	c := child{
		p,
		"Trivendra",
	}

	(&c).born()
	fmt.Println(c.add(1, 2))

}
