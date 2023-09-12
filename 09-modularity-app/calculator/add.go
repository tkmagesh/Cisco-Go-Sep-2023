package calculator

import "fmt"

func init() {
	fmt.Println("init [add.go] invoked")

}
func Add(x, y int) int {
	opCount["Add"]++
	return x + y
}
