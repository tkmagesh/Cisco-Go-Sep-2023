package main

import (
	"fmt"
	"math"
)

// v1.0
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

// v2.0
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

// checking if the x is a Circle or Rectangle, to confirm that x has the Area() method

/*
func PrintArea(x interface{}) {
	switch obj := x.(type) {
	case Circle:
		fmt.Println("Area : ", obj.Area())
	case Rectangle:
		fmt.Println("Area : ", obj.Area())
	default:
		fmt.Println("Argument does not have Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch obj := x.(type) {
	case interface{ Area() float32 }:
		fmt.Println("Area : ", obj.Area())
	default:
		fmt.Println("Argument does not have Area() method")
	}
}
*/

/*
func PrintArea(x interface{ Area() float32 }) {
	fmt.Println("Area : ", x.Area())
}
*/

type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area : ", x.Area())
}

// v3.0
// Implement Perimeter behavior to Circle & Rectangle

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

/*
	func PrintPerimeter(x interface{ Perimeter() float32 }) {
		fmt.Println("Perimeter : ", x.Perimeter())
	}
*/
type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter : ", x.Perimeter())
}

type StatsFinder interface {
	AreaFinder
	PerimeterFinder
}

// v4.0
func PrintShape(x StatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area : ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 12, Width: 12}
	// fmt.Println("Area : ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
