package main

import("fmt")
type person struct{
	name string
	age int
}

func( p person)String() string{
	
	s:=fmt.Sprintf("name:%s and age :%d",p.name,p.age)
	return s
}
func main(){
    p1:=person{"yash",20}
	fmt.Println(p1)
}