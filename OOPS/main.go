

// type Calculator interface{
// 	calculate(int,int)int
// }

// type adder struct{}
// func(add adder) calculate(x,y int)int{
// 	return x+y
// }

// type subtractor struct{}
// func(sub subtractor)calculate(x,y int)int{
// 	return x-y
// }
// func perform(a Calculator,x int,y int)int{
// 	return a.calculate(x,y)
// }
// func main(){
//      add:=adder{}
// 	 sub:=subtractor{}

// 	 fmt.Println(perform(add,2,3))
// 	 fmt.Println(perform(sub,2,3))
// }

package main

import "fmt"

// Struct Definitions
type Rectangle struct {
	length float64
	wi  float64
}

type Cuboid struct {
	Rectangle
	height float64
}

// Interface
type AreaCalculator interface {
	area() float64
}

// Method for Rectangle
func (r Rectangle) area() float64 {
	return r.length * r.wi
}

// Method for Cuboid
func (c Cuboid) area() float64 {
	return 2*(c.length*c.wi + c.wi*c.height + c.height*c.length)
}

// Function to Calculate Area using Interface
func Calculate_Area(obj AreaCalculator) float64 {
	return obj.area()
}

// Main Function
func main() {
	rect := Rectangle{
		length: 4,
		wi:  5,
	}

	cuboid := Cuboid{
		Rectangle: Rectangle{
			length: 4,
			wi:  5,
		},
		height: 9,
	}

	// Using Interface to Calculate Area
	fmt.Println("Area of Rectangle:", Calculate_Area(rect))
	fmt.Println("Area of Cuboid:", Calculate_Area(cuboid))
}
