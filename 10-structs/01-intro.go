package main

import "fmt"

func main() {
	/*
		var product struct {
			id    int
			name  string
			price float64
		}
		product.id = 100
		product.name = "Pen"
		product.price = 10.5
	*/

	var product struct {
		id    int
		name  string
		price float64
	} = struct {
		id    int
		name  string
		price float64
	}{
		id:    100,
		name:  "Pen",
		price: 10.5,
	}
	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	printProduct(product)
}

func printProduct(product struct {
	id    int
	name  string
	price float64
}) {
	fmt.Printf("id = %d, name = %q, price = %0.2f\n", product.id, product.name, product.price)
}
