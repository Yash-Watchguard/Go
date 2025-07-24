package main

import (
	"fmt"
)

func main() {
	foo()
	bar("yash")
	fmt.Println(aloha("yash"))

	x, y := 2, 3
	add, div := sum(x, y)
	fmt.Println(add, div)

}
func foo() {
	fmt.Println("foo called")
}
func bar(s string) {
	fmt.Println("my name is", s)
}
func aloha(s string) string {
	return fmt.Sprint("my name is", s)
}

func sum(x, y int) (int, int) {
	return x + y, x - y
}
