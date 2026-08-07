package main

import "fmt"

func ShowGoroutine(id int) {
	fmt.Printf("Goroutine #%d\n", id)
}

func main() {
	go ShowGoroutine(1)
}
