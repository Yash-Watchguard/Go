package main

import (
	"fmt"
	"runtime"
	// "runtime"
	"sync"
)
func SomeOperation(wg *sync.WaitGroup){
	fmt.Println("hi i am done")
	defer wg.Done()
}
func main() {
    var counter int32
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
			runtime.Gosched()
            counter++

            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
