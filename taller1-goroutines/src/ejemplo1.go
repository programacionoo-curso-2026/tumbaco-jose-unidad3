package main

import (
	"fmt"
	"time"
)

func ShowGoroutine(id int) {
	fmt.Printf("Goroutine #%d\n", id)
}

func main() {
	go ShowGoroutine(1)

	time.Sleep(1 * time.Second)
}
