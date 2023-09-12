package main

import (
	"errors"
	"fmt"
)

var DummyError error = errors.New("dummy error")

func main() {
	pass := true
	err := errorFn(pass)
	if err == DummyError {
		fmt.Println("fn invocation resulted in a dummy error")
		return
	}
	if err != nil {
		fmt.Println("fn invocation resulted in an unknown :", err)
		return
	}
	fmt.Println("fn invoked successfully")
}

func errorFn(pass bool) error {
	if !pass {
		// returning a known (Dummy) error
		/*
			err := DummyError
			return err
		*/

		// returning an unknown error
		return errors.New("unknown error")
	}
	return nil
}
