package main

import (
	"fmt"
	"time"
)

func syHello(){
	fmt.Println("first")
	time.Sleep(3000*time.Millisecond)
	fmt.Println("third")
}
func sayhi(){
    fmt.Println("second")
}
func main(){
	fmt.Println("learning go routines")

	go syHello()
	sayhi()

	time.Sleep(2000 * time.Millisecond)

}