// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// 	// "time"
// )

// func main(){
//     c:=make(chan int)
//     go func(){
//       for i:=0;i<10;i++{
//         c<-i
//       }
//       close(c)
//     }()

//        for val:=range c{
//         fmt.Println(val)
//        }

// //    time.Sleep(time.Second)
// }

// n to 1 (means n goroutine sending the value in one channels)

// func main(){
//     c:=make(chan int)
//     var wg sync.WaitGroup
//     wg.Add(3)
//     go func(){
        
//         for i:=1;i<=10;i++{
//             c<-i
//         }
//         wg.Done()
//     }()
    
//     go func(){
//         for i:=0;i<10;i++{
//             c<-i
//         }
//        wg.Done()
//     }()
//     go func(){
//       wg.Wait()
//       close(c)
//     }()

//     for val:=range c{
//         fmt.Println(val)
//     }
//     time.Sleep(time.Second)
// }


// semaphore
package main

import (
    "fmt"
    "time"
)

func main() {
    c := make(chan int)
    done := make(chan struct{}, 2) // semaphore with buffer size = number of tasks

    // First sender
    go func() {
        for i := 1; i <= 10; i++ {
            c <- i
        }
        done <- struct{}{} // signal done
    }()

    // Second sender
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        done <- struct{}{} // signal done
    }()

    // Closer goroutine (waits for both senders)
    go func() {
        <-done
        <-done
        close(c)
    }()

    for val := range c {
        fmt.Println(val)
    }

    time.Sleep(time.Second)
}
