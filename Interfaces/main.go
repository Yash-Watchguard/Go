package main

import (
	"fmt"
	
)
type person struct{
	name string
	age int
} 
type school interface{
	leave_school()
	come_school()
}

func(p person)leave_school(){
	fmt.Println("Bye-Bye",p.name)
}
func(p person)come_school(){
	fmt.Println("Good Morning",p.name)
}

func message(v school){
    v.come_school()
}
func main(){
	student:=person{"yash",20}
	teacher:=person{"ss",40}

	var aaa school
	aaa=student
	// aaa.leave_school()
	// aaa.come_school()
	message(aaa)

	aaa=teacher
	// aaa.leave_school()
	// aaa.come_school()
	message(aaa)
}
