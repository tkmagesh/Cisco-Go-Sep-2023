package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // increment the waitgroup counter by 1
		go f1()   // scheduling this function to be executed by the scheduler
	}
	f2()
	wg.Wait() // blocks the execution of the main() until the counter in the waitgroup is 0
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
