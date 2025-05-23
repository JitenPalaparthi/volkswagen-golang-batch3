package main

import "fmt"

func main() {
	//db.Select().Where().Filter().GorupBy().OrderBy().Query
	//Calc.Add(10).Sub(20).Mul(2).Get()

	c := NewCalc(100)
	r := c.Add(10).Add(20).Sub(10).Mul(2).Div(2).Add(200).Get()
	fmt.Println("result:", r)
}

type ICalc interface {
	Add(int) ICalc
	Sub(int) ICalc
	Mul(int) ICalc
	Div(int) ICalc
	Get() int
}

type Calc struct {
	data int
}

func NewCalc(d int) ICalc {
	return &Calc{d}
}

func (c *Calc) Add(d int) ICalc {
	c.data += d
	return c
}

func (c *Calc) Sub(d int) ICalc {
	c.data -= d
	return c
}

func (c *Calc) Mul(d int) ICalc {
	c.data *= d
	return c
}
func (c *Calc) Div(d int) ICalc {
	c.data /= d
	return c
}

func (c *Calc) Get() int {
	return c.data
}
