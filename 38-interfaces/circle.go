package main

const PI = 3.17

type Circle float32

func (c Circle) Area() float32 {
	return PI * float32(c*c)
}

func (c Circle) Perimeter() float32 {
	return 2 * PI * float32(c)
}

func (Circle) What() string {
	return "Circle"
}
