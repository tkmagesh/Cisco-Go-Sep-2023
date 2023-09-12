package calculator

import "fmt"

func init() {
	fmt.Println("init [sub.go] invoked")

}

func Subtract(x, y int) int {
	opCount["Subtract"]++
	return x - y
}
