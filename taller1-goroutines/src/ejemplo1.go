package main

import (
	"fmt"
	"time"
)

func ShowGoroutine(id int) {
	fmt.Printf("Goroutine #%d\n", id)
}

func main() {
	for i := 0; i < 10; i++ {
		go ShowGoroutine(i)
	}

	time.Sleep(1 * time.Second)
}
