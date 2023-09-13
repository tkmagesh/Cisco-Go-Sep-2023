package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling this function to be executed by the scheduler
	f2()
	//block the execution of main fn so that the scheduler can execute the other goroutines that are scheduled
	time.Sleep(1 * time.Second)
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
