package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		fmt.Println("Operation started")
		add(100, 200)
		fmt.Println("Operation Completed!")

		fmt.Println("Operation started")
		subtract(100, 200)
		fmt.Println("Operation Completed!")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	}, 100, 200)
}

func logOperation(operation func(int, int), x, y int) {
	fmt.Println("Operation started")
	operation(100, 200)
	fmt.Println("Operation Completed!")
}

func logAdd(x, y int) {
	fmt.Println("Operation started")
	add(100, 200)
	fmt.Println("Operation Completed!")
}

func logSubtract(x, y int) {
	fmt.Println("Operation started")
	subtract(100, 200)
	fmt.Println("Operation Completed!")
}

// 3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
