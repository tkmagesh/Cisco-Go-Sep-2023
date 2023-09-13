package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genFib(stopCh)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				fmt.Println("stop instruction received")
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
