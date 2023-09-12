package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
	*/
	// var productRanks map[string]int = make(map[string]int)
	// var productRanks = make(map[string]int)
	// productRanks := make(map[string]int)
	/*
		productRanks := map[string]int{}
		productRanks["pen"] = 3
	*/
	// productRanks := map[string]int{"pen": 3, "pencil": 2, "notebook": 1}
	productRanks := map[string]int{
		"pen":      3,
		"pencil":   2,
		"notebook": 1,
	}
	fmt.Println(productRanks)
	productRanks["marker"] = 4
	fmt.Println(productRanks)

	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	//check the presence of a key
	keyToCheck := "marker"
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("productRanks[%q] = %d\n", keyToCheck, val)
	} else {
		fmt.Printf("productRanks[%q] does not exist\n", keyToCheck)
	}

	//delete a key (using the delete fn)
	delete(productRanks, "marker")
	fmt.Println(productRanks)

}
