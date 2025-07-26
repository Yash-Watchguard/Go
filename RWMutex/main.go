package main

import ("sync"
"time"
"fmt"
)

type safemap struct {
	data map[string]string
	mu sync.RWMutex
}
func(mp *safemap)Writer(key string, value string){
	(*mp).mu.Lock()
	(*mp).data[key]=value
	defer (*mp).mu.Unlock()
}
func(mp *safemap)reader(key string){
	(*mp).mu.RLock()
	defer (*mp).mu.RUnlock()
	fmt.Println((*mp).data[key])
}

func main() {
	mp:=safemap{data:make(map[string]string)}
	var wg sync.WaitGroup

	for i:=1;i<3;i++{
		wg.Add(1)
		func(i int){
			defer wg.Done()
			key := "name"
			value := "yash"

			mp.Writer(key,value)
		}(i)

	}
	for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            key := fmt.Sprintf("key%d", i)
            time.Sleep(100 * time.Millisecond) // Wait for writer to finish
             mp.reader(key)
            
        }(i)
    }
	fmt.Println(mp)

}