package main

import (
	"fmt"
	"time"
)

func main() {
	// log
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// profile
	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)
	*/

	// log + profile
	/*
		logAdd := getLogOperation(add)
		profiledLogAdd := getProfiledOperation(logAdd)
		profiledLogAdd(100, 200)
	*/

	profiledLogAdd := getProfiledOperation(getLogOperation(add))
	profiledLogAdd(100, 200)

}

func getProfiledOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time elapsed :", elapsed)
	}
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Operation started")
		operation(100, 200)
		fmt.Println("Operation Completed!")
	}
}

// 3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
