/* using channels for streaming data  */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {

	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("Done")
}

// producer
func genNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- i * 10
		}
		fmt.Println("closing the channel")
		close(ch)
	}()
	return ch
}
