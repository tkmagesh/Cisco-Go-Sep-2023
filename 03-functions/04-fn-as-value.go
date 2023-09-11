package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	var operation func(int, int) int
	operation = add
	fmt.Println(operation(10, 20))

	operation = subtract
	fmt.Println(operation(10, 20))
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
