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

type PerishableProduct struct {
	Product
	Expiry string
}

// overriding Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	grapes := PerishableProduct{
		Product: Product{
			Id:    1,
			Name:  "Grapes",
			Price: 100,
		},
		Expiry: "2 days",
	}
	fmt.Println(grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}
