/*
Build an interactive calculator

Display the following menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit

Accept the user choice

if the user choice == 5
	exit the application

if the user choice is 1 - 4
	accept 2 numbers
	perform the corresponding operation
	print the result
	IMPORTANT : display the menu again

if the user input is not (1-5)
	print "Invalid choice"
	IMPORTANT : display the menu again

Organize your logic using functions and make sure that no function has more than one responsibility (SRP)
*/

package main

import "fmt"

func main() {
	var x, y, result, userChoice int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice. Try again..!")
			continue
		}
		x, y = getOperands()
		result = processUserChoice(userChoice, x, y)
		fmt.Println("Result :", result)
	}
}

func processUserChoice(userChoice, x, y int) int {
	var result int
	switch userChoice {
	case 1:
		result = add(x, y)
	case 2:
		result = subtract(x, y)
	case 3:
		result = multiply(x, y)
	case 4:
		result = divide(x, y)
	}
	return result
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (x, y int) {
	fmt.Println("Enter the operands :")
	fmt.Scanln(&x, &y)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x + y
}

func divide(x, y int) int {
	return x + y
}
