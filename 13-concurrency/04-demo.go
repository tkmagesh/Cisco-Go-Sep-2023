package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "number of goroutines to spawn")
	flag.Parse()
	fmt.Printf("Spawning %d goroutines. Hit ENTER to start...\n", count)
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	for i := 1; i <= count; i++ {
		wg.Add(1)    // increment the waitgroup counter by 1
		go f1(wg, i) // scheduling this function to be executed by the scheduler
	}
	wg.Wait() // blocks the execution of the main() until the counter in the waitgroup is 0
	fmt.Println("All goroutines completed. Hit ENTER to exit...")
	fmt.Scanln()
}

func f1(wg *sync.WaitGroup, id int) {
	fmt.Printf("f1-[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("f1-[%d] completed\n", id)
	wg.Done() // decrement the counter by 1
}
