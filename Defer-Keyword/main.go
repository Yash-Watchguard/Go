package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	hour := time.Now().Hour()

	defer fmt.Println(today)
	fmt.Println(hour)
}
