package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(10)
	go f1()
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
