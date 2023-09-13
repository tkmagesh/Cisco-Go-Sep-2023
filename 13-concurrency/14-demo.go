/* using channels for streaming data  */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
	}()
	return ch
}
