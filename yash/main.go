package main

import(
	"fmt"
)
type types interface{
	int | float64 
}

func add[T types](x, y T)T{
   return x+y
}
func main(){
    fmt.Println(add(3,5))
}