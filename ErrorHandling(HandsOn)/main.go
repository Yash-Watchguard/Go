package main

import (
	"fmt"
)

type CustErr struct{
	mess string
	code int
}
func(e CustErr)Error()string{
fmt.Sprintf("Error %d is %s",e.code,e.mess)
}

func main(){
  
}