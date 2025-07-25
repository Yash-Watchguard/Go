package main

import ("fmt")

type Calculator interface{
	calculate(int,int)int
}

type adder struct{}
func(add adder) calculate(x,y int)int{
	return x+y
}

type subtractor struct{}
func(sub subtractor)calculate(x,y int)int{
	return x-y
}
func perform(a Calculator,x int,y int)int{
	return a.calculate(x,y)
}
func main(){
     add:=adder{}
	 sub:=subtractor{}

	 fmt.Println(perform(add,2,3))
	 fmt.Println(perform(sub,2,3))
}