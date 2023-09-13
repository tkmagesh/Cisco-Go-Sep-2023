package main

import (
	"fmt"
)

// share memory by communicating

// consumer
func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
