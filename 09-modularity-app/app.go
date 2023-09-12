package main

import (
	"fmt"

	"github.com/tkmagesh/cisco-go-sep-2023/09-modularity-app/calculator"
	"github.com/tkmagesh/cisco-go-sep-2023/09-modularity-app/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red("app invoked")
	fmt.Println(greet("Magesh"))
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOpCount())
	fmt.Println(utils.IsPrime(79))
}
