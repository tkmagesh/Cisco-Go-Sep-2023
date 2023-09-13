package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

func NewProduct(id int, name string, price float64) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
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

func NewPerishableProduct(id int, name string, price float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:    id,
			Name:  name,
			Price: price,
		},
		Expiry: expiry,
	}
}

type SuperPerishableProduct struct {
	PerishableProduct
	Discount float32
}

func NewSuperPerishableProduct(id int, name string, price float64, expiry string, discount float32) *SuperPerishableProduct {
	return &SuperPerishableProduct{
		PerishableProduct: PerishableProduct{
			Product: Product{
				Id:    id,
				Name:  name,
				Price: price,
			},
			Expiry: expiry,
		},
		Discount: discount,
	}
}

func main() {
	// var grapes *PerishableProduct = NewPerishableProduct(101, "Grapes", 50, "2 Days")
	// var grapes  = NewPerishableProduct(101, "Grapes", 50, "2 Days")
	grapes := NewPerishableProduct(101, "Grapes", 50, "2 Days")
	fmt.Printf("%#v\n", grapes)
	fmt.Println(grapes.Expiry)
	// fmt.Println(grapes.Product.Price)
	fmt.Println(grapes.Price)

	fmt.Println(grapes.Product.Name)
	// fmt.Println(grapes.Dummy.Name)

	milk := NewSuperPerishableProduct(102, "Milk", 20, "1 Month", 10)
	fmt.Println(milk.Name)
}
