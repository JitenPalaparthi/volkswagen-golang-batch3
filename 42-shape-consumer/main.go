package main

import (
	"fmt"

	"github.com/JitenPalaparthi/volkswagen-golang-batch3-shapes/shape"
	"github.com/JitenPalaparthi/volkswagen-golang-batch3-shapes/shape/circle"
	"github.com/JitenPalaparthi/volkswagen-golang-batch3-shapes/shape/greet"
	"github.com/JitenPalaparthi/volkswagen-golang-batch3-shapes/shape/rect"
	"github.com/JitenPalaparthi/volkswagen-golang-batch3-shapes/shape/square"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Hello World, start of my application")
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
	//shape.GreetFromShape()
	greet.Greet() // Greet is Upper case , so unrestricted
	//	greet.greet() // this cannot be called since greet() func is restricted

	var p1 greet.Public
	fmt.Println(p1.PublicField) // unrestricted bcz starts with UpperCase
	//fmt.Println(p1.privateField) // restricted stats with lowerCase
	log.Info().Msg(" end of my application")
	// UUID

	id := uuid.New()
	fmt.Println("UUID:", id.String())
	shape.SayHi()
}

// a package must have a directory
// the name of the package is the name of the directory
