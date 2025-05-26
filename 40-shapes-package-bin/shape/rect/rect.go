package rect

type Rect struct {
	L, B float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (*Rect) What() string {
	return "Rect"
}
