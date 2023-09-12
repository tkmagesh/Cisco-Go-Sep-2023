/*
create a divide function that returns quotient & remainder
if the divisor is 0, return an error and handle it in the main function
*/

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == DivideByZeroError {
		fmt.Println("please do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("unknown error :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)

}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
