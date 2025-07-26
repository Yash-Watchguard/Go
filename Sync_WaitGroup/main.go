package main

// import "sync"
import (
	"fmt"
	"runtime"
	"sync"
)

func Worker(i int, wg *sync.WaitGroup){
  defer wg.Done()
fmt.Printf("worker %d started his task\n",i)
fmt.Printf("worker %d completed his task\n",i)

}
func main() {
	var wg sync.WaitGroup
   for i:=1;i<=3;i++{
	wg.Add(1)
	 go Worker(i,&wg)
   }
fmt.Println(runtime.NumGoroutine())
   wg.Wait()
   fmt.Println("work is completed")
   fmt.Println(runtime.GOOS)
   fmt.Println(runtime.NumCPU())
   fmt.Println(runtime.NumGoroutine())
}