package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("[main] - deferred function")
		if err := recover(); err != nil {
			fmt.Println("[main] - recovered from panic, err :", err)
			return
		}
		fmt.Println("[main] - executed successfully")
	}()
	divisor := 7
	q, r := divide(100, divisor)
	fmt.Printf("[main] dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[divide] - calculating quotient")
	quotient = x / y // results in a panic if y == 0
	fmt.Println("[divide] - calculating remainder")
	remainder = x % y
	return
}
