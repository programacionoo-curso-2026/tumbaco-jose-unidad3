package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ShowGoroutine(id int) {
	delay := rand.Intn(500)

	fmt.Printf("Goroutine #%d\nwith %dms\n", id, delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func main() {
	for i := 0; i < 10; i++ {
		go ShowGoroutine(i)
	}

	time.Sleep(1 * time.Second)
}
