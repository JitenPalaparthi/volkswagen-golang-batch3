package main

import "fmt"

func main() {

	r1 := Rect{L: 98.3, B: 78.45}
	a1, p1 := Area(r1), Perimeter(r1)
	fmt.Printf("Area of Rect r1:%.3f\n", a1)
	fmt.Printf("Perimeter of Rect r1:%.3f\n", p1)

	a2, p2 := r1.Area(), (&r1).Perimeter() // no need to use (&r1).Perimeter
	fmt.Printf("Area of Rect r1:%.3f\n", a2)
	fmt.Printf("Perimeter of Rect r1:%.3f\n", p2)

	fmt.Printf("Area of Rect r1:%.3f\n", r1.A)
	fmt.Printf("Perimeter of Rect r1:%.3f\n", r1.P)

	r2 := &Rect{L: 98.3, B: 78.45}
	a3 := (*r2).Area() // no need to call as it is a pointer receiver
	p3 := r2.Perimeter()
	fmt.Printf("Area of Rect r2:%.3f\n", a3)
	fmt.Printf("Perimeter of Rect r2:%.3f\n", p3)
	fmt.Printf("Area of Rect r2:%.3f\n", r2.A)
	fmt.Printf("Perimeter of Rect r2:%.3f\n", r2.P)

	r3 := NewRect(19.45, 67.8)
	fmt.Println(r3)

	var r4 *Rect = &Rect{}
	r4.L = 78.56
	r4.B = 56.76
	fmt.Println(r4)

	r5 := new(Rect)
	r5.L = 78.56
	r5.B = 56.76
	fmt.Println(r5)

}

type Rect struct {
	L, B float32
	A, P float32
}

func Area(r Rect) float32 {
	r.A = r.L * r.B
	return r.A
}

func Perimeter(r Rect) float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

// func (r Rect) Area() float32 {
// 	r.A = r.L * r.B
// 	return r.A
// }

// func (rectangle Rect) Perimeter() float32 {
// 	rectangle.P = 2 * (rectangle.L + rectangle.B)
// 	return rectangle.P
// }

// pointer receivers,call by reference

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

// idiomatic approach is always use a single letter receiver
func (rectangle *Rect) Perimeter() float32 {
	rectangle.P = 2 * (rectangle.L + rectangle.B)
	return rectangle.P
}
