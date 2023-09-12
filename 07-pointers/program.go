package main

import "fmt"

func main() {
	var no int = 10

	var noPtr *int
	noPtr = &no // address from a value
	fmt.Println(noPtr)

	val := *noPtr // value from an address (deferencing)
	fmt.Println(val)

	fmt.Println("Before increment:", no)
	increment(&no)
	fmt.Println("After increment:", no)

}

func increment(x *int) {
	fmt.Println("[increment] address of x:", x)
	*x += 1
}

func swap( /*  */ ) /* no return values */ {

}
