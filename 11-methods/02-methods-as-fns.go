package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Price = %0.2f", product.Id, product.Name, product.Price)
}

func ApplyDiscount(product *Product, discount float64) {
	product.Price = product.Price * ((100 - discount) / 100)
}

func main() {
	pen := Product{
		Id:    1,
		Name:  "Pen",
		Price: 2.5,
	}
	fmt.Println(Format(pen))
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))

	/*
		pen := &Product{
			Id:    1,
			Name:  "Pen",
			Price: 2.5,
		}
		fmt.Println(Format(*pen))
		ApplyDiscount(pen, 10)
		fmt.Println(Format(*pen))
	*/
}
