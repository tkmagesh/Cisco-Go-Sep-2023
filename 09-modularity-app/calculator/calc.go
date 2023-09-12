package calculator

import "fmt"

/*
var opCount int

func GetOpCount() int {
	return opCount
}
*/

var opCount map[string]int

func init() {
	fmt.Println("init [calc.go] invoked")
	opCount = make(map[string]int)
}

func GetOpCount() map[string]int {
	return opCount
}
