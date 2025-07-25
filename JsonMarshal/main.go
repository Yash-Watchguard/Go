package main 

import(
	"fmt"
	"encoding/json"
)
type person struct{
	Name string `json:"name"`
	Age int `jsn:"age"`
}
func main(){
 p1:=person{Name:"yash",Age: 20}
 p2:=person{Name:"rekha",Age:40}

 c:= []person{p1,p2}

 jasondata,err:= json.Marshal(c)
 if err==nil{
	fmt.Println(string(jasondata))
 }else{
	fmt.Println(err)
 }
}