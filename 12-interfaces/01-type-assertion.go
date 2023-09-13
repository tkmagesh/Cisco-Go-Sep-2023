package main

import "fmt"

func main() {
	var x interface{}
	x = 200
	x = "asdfasfd"
	x = 19.99
	x = true
	x = struct{}{}
	fmt.Println(x)

	x = 100
	// y := x + 200 // compilation error
	x = "asfsadfasd"
	// y := x.(int) + 200 // results in a runtime error if x has a non integer value
	// check if the x can be converted to int (type assertion)
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an integer, val = ", val)
	}

	// var y interface{}
	var y any
	// y = 100
	// y = true
	// y = 19.99
	y = struct{}{}
	/*
		y = struct {
			id int
		}{
			id: 100,
		}
	*/
	switch val := y.(type) {
	case int:
		fmt.Println("y is an integer, y * 2 = ", val*2)
	case string:
		fmt.Println("y is a string, len(y) = ", len(val))
	case bool:
		fmt.Println("y is a bool, !y = ", !val)
	case float64:
		fmt.Println("y is a float64, y * 0.95 = ", val*0.95)
	case struct{}:
		fmt.Println("y is a struct")
	case nil:
		fmt.Println("y is not initialized")
	default:
		fmt.Println("y is of unknown type")
	}

}
