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

	x, y := 100, 200
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)
}

func increment(x *int) {
	fmt.Println("[increment] address of x:", x)
	*x += 1
}

func swap(n1, n2 *int) /* no return values */ {
	*n1, *n2 = *n2, *n1
}
