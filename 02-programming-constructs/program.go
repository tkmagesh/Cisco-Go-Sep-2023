/* programming constructs */

package main

import "fmt"

func main() {
	/* if else */
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	// the scope of "no" variable is limited to if-else block
	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	/* switch case */
	/*
			rank by score
			0 - 3 => "Very Bad"
			4 - 7 => "Poor"
		 	8 - 9 => "Good"
			10 => "Excellent"
			otherwise => "Invalid score"
	*/
	/*
		score := 6
		switch score {
		case 0:
			fmt.Println("Very Bad")
		case 1:
			fmt.Println("Very Bad")
		case 2:
			fmt.Println("Very Bad")
		case 3:
			fmt.Println("Very Bad")
		case 4:
			fmt.Println("Poor")
		case 5:
			fmt.Println("Poor")
		case 6:
			fmt.Println("Poor")
		case 7:
			fmt.Println("Poor")
		case 8:
			fmt.Println("Good")
		case 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/
	/*
		score := 6
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Very Bad")
		case 4, 5, 6, 7:
			fmt.Println("Poor")
		case 8, 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/

	switch score := 6; score {
	case 0, 1, 2, 3:
		fmt.Println("Very Bad")
	case 4, 5, 6, 7:
		fmt.Println("Poor")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	// fallthrough
	switch no := 4; no {
	case 0:
		fmt.Println("no = 0")
		fallthrough
	case 1:
		fmt.Println("no <= 1")
		fallthrough
	case 2:
		fmt.Println("no <= 2")
		fallthrough
	case 3:
		fmt.Println("no <= 3")
		fallthrough
	case 4:
		fmt.Println("no <= 4")
		fallthrough
	case 5:
		fmt.Println("no <= 5")
		fallthrough
	case 6:
		fmt.Println("no <= 6")
		// fallthrough
	case 7:
		fmt.Println("no <= 7")
		fallthrough
	case 8:
		fmt.Println("no <= 8")
	}

	fmt.Println("fallthrough applied")
	switch plan := "Premium"; plan {
	case "Super":
		fmt.Println("[Super] License for a family of 3")
		fallthrough
	case "Premium":
		fmt.Println("[Premium] Download for offline experience")
		fallthrough
	case "Pro":
		fmt.Println("[Pro] Create private playlist")
		fallthrough
	case "Free":
		fmt.Println("[Free] Listen to music")
	}

	fmt.Println("Using switch like if-else-if")
	switch no := 31; {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%2 != 0:
		fmt.Printf("%d is an odd number\n", no)
	}

	// for
	fmt.Println()
	fmt.Println("For")

	fmt.Println("v1.0")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("v2.0 [while]")
	x := 1
	for x < 100 {
		x += x
	}
	fmt.Println("x =", x)

	fmt.Println("v3.0 [infinite]")
	sum := 1
	for {
		sum += sum
		if sum > 100 {
			break
		}
	}
	fmt.Println("sum =", sum)

	// using labels
OUTER_LOOP:
	for x := 0; x <= 10; x++ {
		for y := 0; y <= 10; y++ {
			fmt.Printf("x = %d, y = %d\n", x, y)
			if x == y {
				fmt.Println("=========================")
				continue OUTER_LOOP
				// break OUTER_LOOP
			}
		}
	}

}
