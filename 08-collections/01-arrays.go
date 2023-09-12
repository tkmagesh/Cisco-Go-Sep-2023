package main

import "fmt"

func main() {
	// var nos [5]int // memory is allocated and initialized
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}

	nos := [5]int{3, 1, 4, 2, 5}

	fmt.Println("Iterating an array using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array using range")
	for i, no := range nos {
		fmt.Printf("nos[%d] = %d\n", i, no)
	}

	x := [3]int{10, 20, 30}
	fmt.Printf("%p\n", &x)

	y := [3]int{10, 20, 30}
	fmt.Printf("%p\n", &y)

	fmt.Println(x == y) // everything is a value in go

	//accessing the values using the pointer
	var nosPtr *[5]int
	nosPtr = &nos
	fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0]) // NO NEED to deference the pointer to an array to access the individual elements

	sort(&nos) // sort the nos array
	fmt.Println(nos)

	nos2 := nos // COPY the nos array
	nos2[0] = 100
	fmt.Println(nos)
	fmt.Println(nos2)
}

func sort(values *[5]int) /* no return values */ {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
