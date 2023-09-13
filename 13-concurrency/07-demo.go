// shared state

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		// fn(wg)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("Counter :", counter)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)
}
