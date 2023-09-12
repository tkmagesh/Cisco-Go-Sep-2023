package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	cisco := Organization{"Cisco", "San Jose"}
	e1 := Employee{1, "John", &cisco}
	e2 := Employee{2, "Jane", &cisco}

	cisco.City = "Bengaluru"

	fmt.Println(e1.Org.City)
	fmt.Println(e2.Org.City)

}
