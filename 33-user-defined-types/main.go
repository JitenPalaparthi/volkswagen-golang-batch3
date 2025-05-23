package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a1 := rect{}.Area(10.12, 12.34)
	a2 := square{}.Area(23.45)
	fmt.Println("Area of rect:", a1)
	fmt.Println("Area of square:", a2)

	var r1, r2 rect
	var s1, s2 square

	fmt.Printf("Size of r1:%d address of r1:%p\n", unsafe.Sizeof(r1), &r1)
	fmt.Printf("Size of s1:%d address of s1:%p\n", unsafe.Sizeof(s1), &s1)
	fmt.Printf("Size of r2:%d address of r2:%p\n", unsafe.Sizeof(r2), &r2)
	fmt.Printf("Size of s2:%d address of s2:%p\n", unsafe.Sizeof(s2), &s2)

}

// empty struct
type rect struct{} // the size os 0 bytes, still it has a valid address

type square struct{}

func (rect) Area(l, b float32) float32 {
	return l * b
}

func (square) Area(s float32) float32 {
	return s * s
}
