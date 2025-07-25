package main

import(
	"fmt"
	"os"
	"io"
)

func main(){
	// first create a text file
	file,err:=os.Create("temp.txt")

	if err!=nil{
		fmt.Println("file is not created successfully",err)
        return
	}
	defer file.Close()

	// add the content in file

	content:="hi , i am yash"

	_,_ = io.WriteString(file,content)

	// i wnat to read the data from file
	// 1. by os.Readflie

	data,error:=os.ReadFile("temp.txt")
	if error!=nil{
		fmt.Println("error reading file:",error)
	}
	defer file.Close()
    _,_=io.WriteString(file,"hjinkkjikljkm,mk,")
	fmt.Println(string(data))

	





}
