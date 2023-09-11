package main

import "fmt"

func main() {
	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/

	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/

	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red   = iota + 2 // 0 + 2
			green            // 1 + 2
			blue             // 2 + 2
		)
	*/

	/*
		const (
			red   = iota*2 + 1 // 0 * 2 + 1
			green              // 1 * 2 + 1
			blue               // 2 * 2 + 1
		)
	*/

	const (
		red   = iota*2 + 1 // 0 * 2 + 1
		green              // 1 * 2 + 1
		_
		blue // 3 * 2 + 1
	)
	fmt.Printf("red = %d, green = %d, blue = %d\n", red, green, blue)

	fmt.Println("iota applied")
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
	// fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)

	fmt.Printf("%T\n", VERBOSE)
}
