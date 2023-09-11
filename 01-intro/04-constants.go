package main

import "fmt"

func main() {
	// const pi float32 = 3.14
	// const pi = 3.14
	// fmt.Println(pi)

	/*
		const (
			pi          float32 = 3.14
			app_version string  = "2.0.0"
		)
	*/

	/* type inference */
	const (
		pi          = 3.14
		app_version = "2.0.0"
	)
	fmt.Println(pi, app_version)
}
