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
}

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// with 2 parameter & return result
func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// with 2 parameters & 2 return results
// named results
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
