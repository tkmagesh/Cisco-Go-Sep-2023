package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divide by zero error")

func main() {
	defer func() {
		fmt.Println("[main] - deferred function")
		if err := recover(); err != nil {
			fmt.Println("[main] - recovered from panic, err :", err)
			return
		}
		fmt.Println("[main] - executed successfully")
	}()
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideClient(100, divisor)
		if err != nil {
			fmt.Println("[main] - error occurred, err :", err)
			fmt.Println("try again")
			continue
		}
		fmt.Printf("[main] dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// intended to convert the panic into an error
func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party library (can result in a panic)
func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[divide] - calculating quotient")
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient = x / y
	fmt.Println("[divide] - calculating remainder")
	remainder = x % y
	return
}
