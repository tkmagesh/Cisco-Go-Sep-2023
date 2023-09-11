package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	greetUser("Magesh", "Kuppan")
	fmt.Print(getGreetMsg("Magesh", "Kuppan"))

	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// Use only quotient
	// q, r := divide(100, 7)
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d \n", q)
	*/
}

func sayHi() {
	fmt.Println("Hi there!")
}

// with 1 parameter
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

// with 2 parameters
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// with 2 parameter & return result
func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// with 2 parameters & 2 return results
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// named results
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
