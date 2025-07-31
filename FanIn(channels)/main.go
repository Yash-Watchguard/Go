package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendvalue(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	wg.Done()
}

func reveivevalue(even <-chan int, odd <-chan int, other chan<- int) {
	for val := range even {
		other <- val
	}
	for val := range odd {
		other <- val
	}
	close(other)
	wg.Done()
}

func Printother(other <-chan int) {
	defer wg.Done() // ✅ Use defer so it always runs
	for val := range other {
		fmt.Println(val)
	}
}

func main() {
	even := make(chan int)
	odd := make(chan int)
	other := make(chan int)

	wg.Add(3)

	go sendvalue(even, odd)
	go reveivevalue(even, odd, other)
	go Printother(other) //  MUST run in a goroutine

	wg.Wait()
}