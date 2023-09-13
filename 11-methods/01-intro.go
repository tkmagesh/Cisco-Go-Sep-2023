package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Price = %0.2f", product.Id, product.Name, product.Price)
}

func (product *Product) ApplyDiscount(discount float64) {
	product.Price = product.Price * ((100 - discount) / 100)
}

func main() {
	pen := Product{
		Id:    1,
		Name:  "Pen",
		Price: 2.5,
	}
	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println("After applying 10% discount")
	fmt.Println(pen.Format())
}
