package main

import (
	"fmt"
)

func main() {
	i, j := 22, 23

	p, q := &i, &j

	fmt.Println(p, q)
	fmt.Println(*p, *q)
	*q++
	*p++
	fmt.Println(i, j)
	fmt.Println(*, *p)
	

}
