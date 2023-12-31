package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating a slice using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating a slice using range")
	for i, no := range nos {
		fmt.Printf("nos[%d] = %d\n", i, no)
	}

	// appending new values
	// nos = append(nos, 6)
	// nos = append(nos, 7, 8, 9)
	tens := []int{10, 20, 30}
	nos = append(nos, tens...)
	fmt.Println(nos)

	sort(nos)
	fmt.Println(nos)

	nos2 := nos[2:4]
	fmt.Println(nos2)

	// change nos2
	nos2[0] = 999
	fmt.Println("nos :", nos)
	fmt.Println("nos2 :", nos2)

	fmt.Println("nos[2:5] :", nos[2:5])
	fmt.Println("nos[:5] :", nos[:5])
	fmt.Println("nos[2:] :", nos[2:])

	// appending the copy of the slice
	nos2 = append(nos2, 1000)
	fmt.Println("After appending to nos2")
	fmt.Println("nos :", nos)
	fmt.Println("nos2 :", nos2)
}

func sort(values []int) /* no return values */ {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
