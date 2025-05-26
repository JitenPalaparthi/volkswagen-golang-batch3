package square

type Square struct {
	S float32
}

func NewSquare(s float32) *Square {
	return &Square{s}
}

func (r *Square) Area() float32 {
	return r.S * r.S
}

func (s *Square) Perimeter() float32 {
	return 4 * s.S
}

func (*Square) What() string {
	return "Square"
}
