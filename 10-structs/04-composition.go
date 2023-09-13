package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product // composed / inherited
	/*
		Dummy
		Price  float64 //overriding the Price attribute from the base (Product)
	*/
	Expiry string
}

type SuperPerishableProduct struct {
	PerishableProduct
	Discount float32
}

func main() {
	// var grapes PerishableProduct
	// var grapes PerishableProduct = PerishableProduct{Product{101, "Grapes", 50}, "2 Days"}
	var grapes PerishableProduct = PerishableProduct{
		Product: Product{Id: 101, Name: "Grapes", Price: 50},
		Expiry:  "2 Days",
	}
	fmt.Printf("%#v\n", grapes)
	fmt.Println(grapes.Expiry)
	// fmt.Println(grapes.Product.Price)
	fmt.Println(grapes.Price)

	fmt.Println(grapes.Product.Name)
	// fmt.Println(grapes.Dummy.Name)

	milk := SuperPerishableProduct{
		PerishableProduct: PerishableProduct{
			Product: Product{
				Id:    102,
				Name:  "Milk",
				Price: 30,
			},
			Expiry: "1 Day",
		},
		Discount: 20,
	}
	fmt.Println(milk.Name)
}
