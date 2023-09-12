package main

import (
	"errors"
	"fmt"
)

var F1Error error = errors.New("f1 error")
var F2Error error = errors.New("f2 error")

func main() {
	err := f1()
	if errors.Is(err, F1Error) {
		fmt.Println("f1 error occurred")
	}
	if errors.Is(err, F2Error) {
		fmt.Println("f2 error occurred")
	}

}

func f1() error {
	/*
		err := f2()
		combinedErr := errors.Join(F1Error, err)
		return combinedErr
	*/
	return fmt.Errorf("combined error: %w", f2())
	// return F1Error
}

func f2() error {
	return F2Error
}
