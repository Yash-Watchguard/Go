package main

import (
	"fmt"
)

type value struct {
	x int
	y int
}

func (v *value) modify() {
	v.x = v.x * 10
	v.y = v.y * 10
}
func main() {
	v1 := value{2, 2}
	v1.modify()

	fmt.Println(v1)

}
