package main

import("fmt")
type youngin interface{
	modify()
}
type mystring string
func(st *mystring) modify(){
    *st= "yash"
}

func main(){
  var i youngin

  st := mystring("chaaal")
   fmt.Println(st)
  i=&st
  i.modify()

  fmt.Println(st)

}