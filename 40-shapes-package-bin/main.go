package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/circle"
	"shapes/shape/greet"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {

	r1 := rect.NewRect(12.3, 14.5)
	s1 := square.NewSquare(23.4)
	c1 := circle.Circle(12.3)

	r2 := rect.NewRect(14.3, 16.5)
	s2 := square.NewSquare(33.4)
	c2 := circle.Circle(15.3)

	shapeSlice := append([]shape.IShape{}, r1, s1, c1, r2, s2, c2, rect.NewRect(56.67, 45.67), square.NewSquare(89.78), circle.Circle(123.4))

	for _, v := range shapeSlice {
		shape.Shape(v)
		println()
	}
	shape.GreetFromShape()
	greet.Greet() // Greet is Upper case , so unrestricted
	//	greet.greet() // this cannot be called since greet() func is restricted

	var p1 greet.Public
	fmt.Println(p1.PublicField) // unrestricted bcz starts with UpperCase
	//fmt.Println(p1.privateField) // restricted stats with lowerCase
}

// a package must have a directory
// the name of the package is the name of the directory
