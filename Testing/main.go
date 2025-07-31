// package main

// import (
// 	"fmt"
	
// )
// func sum(x,y int)int{
// 	var s int
	
// 		s=x+y
		
	
// 	return s
// }
// func multiply(x,y int)int{
// 	return x*y
// }
// func main(){
//      x:=sum(1,3)
// 	 fmt.Print(x)
     
// 	 y:=multiply(2,3)
// 	 fmt.Println(y)
	 
// }
package main

import (
	"fmt"
	
)
func sum(x ...int)int{
	var s int
	
		for _,v:=range x{
			s+=v
		}
		
	
	return s+1
}

func main(){
	
     x:=sum(1,3)
	 fmt.Printf("%d","x")
	 fmt.Print(x)
}