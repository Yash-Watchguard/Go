package main

import (
	"fmt"
	"runtime"
)
type person struct{
	Name string 
	Age int
}
type human interface{
   speak()
}
func(p1 *person)speak(){
  fmt.Printf("hi ,i am %s",p1.Name)
}
func SaySomthing(p1 * person){
 fmt.Printf("hi ,i am %s",p1.Name)
}

func main(){
	p1:=person{
		"yash",
		18,
	}
	(&p1).speak()// allowed
	p1.speak()//also allowed
	SaySomthing(&p1) // only this allowed
	// SaySomthing(p1) not allowed
	
}