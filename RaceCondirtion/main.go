package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{0}

	// func() {
	// 	array = append(array, 1)
	// }()
	// func() {
	// 	array = append(array, 2)
	// }()
	// func() {
	// 	array = append(array, 3)
	// }()
	// fmt.Println(array)

	// Problem is occure when we create the goroutine first make sure each and every goroutine run (but this code have rave condition)
    /*var wg sync.WaitGroup
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		array = append(array, 1)
		defer wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		array = append(array, 2)
		defer wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		array = append(array, 3)
		defer wg.Done()
	}(&wg)
	wg.Wait()
	fmt.Println(array)*/

	//  now i am using Mutax to remove race condition
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(3)
	go func(wg *sync.WaitGroup,mu *sync.Mutex) {
		mu.Lock()
		array = append(array, 1)
		mu.Unlock()
		defer wg.Done()
	}(&wg,&mu)
	go func(wg *sync.WaitGroup,mu *sync.Mutex) {
		mu.Lock()
		array = append(array, 2)
		mu.Unlock()
		defer wg.Done()
	}(&wg,&mu)
	go func(wg *sync.WaitGroup,mu *sync.Mutex) {
		mu.Lock()
		array = append(array, 3)
		mu.Unlock()
		defer wg.Done()
	}(&wg,&mu)
	wg.Wait()
	fmt.Println(array)

}