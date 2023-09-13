package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	ch := make(chan int)

	go func() {
		x, y := 0, 1
		timeoutCh := time.After(4 * time.Second)
	LOOP:
		for {
			select {
			case <-timeoutCh:
				fmt.Println("timed out!")
				break LOOP
			default:
				ch <- x
				time.Sleep(300 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeOut(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
} */
