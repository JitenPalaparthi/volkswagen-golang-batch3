package main

import "fmt"

func main() {

	a, b := 10, 20

	r1 := Calc(a, b, func(i1, i2 int) int {
		return i1 + i2
	})
	r2 := Calc(a, b, sub)

	fn1 := func(a, b int) int {
		return a * b
	}

	r3 := Calc(a, b, fn1)

	var fn2 Fn = func(i1, i2 int) int {
		return i1 / i2
	}
	//r4 := Calc(a, b, (func(int, int) int)(fn2))
	r4 := Calc(a, b, fn2)

	fmt.Println(fn2.Sq(4))

	fmt.Println(r1, r2, r3, r4)
}

type Fn func(int, int) int

func (f Fn) Sq(n int) int {
	return n * n
}

// func (f Fn) Add(int, int) Fn

func Calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func sub(a, b int) int {
	return a - b
}
