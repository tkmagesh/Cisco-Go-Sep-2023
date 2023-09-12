package main

import "fmt"

type Product struct {
	id    int
	name  string
	price float64
}

func main() {
	/*
		var product Product
		product.id = 100
		product.name = "Pen"
		product.price = 10.5
	*/

	// preferred
	var product = Product{
		id:    100,
		name:  "Pen",
		price: 10.5,
	}

	// var product = Product{100, "Pen", 10.5} // not advisable
	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	printProduct(product)

	pencilPtr := &Product{102, "Pencil", 2.9}
	// fmt.Println(pencilPtr)
	// fmt.Println((*pencilPtr).id, (*pencilPtr).name, (*pencilPtr).price)
	fmt.Println(pencilPtr.id, pencilPtr.name, pencilPtr.price)

	markerPtr := new(Product)
	fmt.Println(markerPtr)

}

func printProduct(product Product) {
	fmt.Printf("id = %d, name = %q, price = %0.2f\n", product.id, product.name, product.price)
}
