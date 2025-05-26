package shape

import (
	"fmt"
	"shapes/shape/greet"
)

type IShape interface {
	//Area() float32
	//Perimeter() float32
	IArea
	IPerimter
	IWhat
}

type IArea interface {
	Area() float32
}

type IPerimter interface {
	Perimeter() float32
}
type IWhat interface {
	What() string
}

func Shape(ishape IShape) {
	fmt.Printf("Area of %s: %.2f\n", ishape.What(), ishape.Area())
	fmt.Printf("Perimeter of %s: %.2f\n", ishape.What(), ishape.Perimeter())
}
func GreetFromShape() {
	greet.Greet()
}
