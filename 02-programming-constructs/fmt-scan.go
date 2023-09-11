package main

import "fmt"

func main() {
	var s string
	fmt.Println("Enter the text:")
	fmt.Scanln(&s)
	fmt.Println(s)

	var x, y int
	fmt.Println("Enter the two numbers:")
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)
}
