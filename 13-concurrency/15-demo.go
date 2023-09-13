/* using channels for streaming data  */

package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	count := 20
	ch := genNos(count)
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

// producer
func genNos(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
	}()
	return ch
}
