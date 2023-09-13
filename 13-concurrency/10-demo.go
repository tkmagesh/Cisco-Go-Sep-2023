/* channel behaviors */
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	/*
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	/*
		ch <- 100
		go func() {
			data := <-ch
			fmt.Println(data)
		}()
	*/
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		data := <-ch
		fmt.Println(data)
		wg.Done()
	}()
	ch <- 100
	wg.Wait()
}
