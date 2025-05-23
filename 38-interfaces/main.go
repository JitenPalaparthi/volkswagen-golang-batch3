package main

func main() {
	r1 := NewRect(12.3, 14.5)
	s1 := NewSquare(23.4)
	c1 := Circle(12.3)

	r2 := NewRect(14.3, 16.5)
	s2 := NewSquare(33.4)
	c2 := Circle(15.3)

	shapeSlice := append([]IShape{}, r1, s1, c1, r2, s2, c2, NewRect(56.67, 45.67), NewSquare(89.78), Circle(123.4))

	for _, v := range shapeSlice {
		Shape(v)
		println()
	}

}
