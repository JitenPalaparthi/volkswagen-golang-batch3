package main

import "fmt"

func main() {

	var num1 int = 100
	var num2 myInt1 = 200
	var num3 myInt2 = 300
	var num4 myInt3 = 400

	var float1 float32 = 12.3

	fmt.Println(myInt1(num1).ToString())
	fmt.Println(num2.ToString())
	fmt.Println(myInt1(num3).ToString())
	fmt.Println(myInt1(num4).ToString())

	fmt.Println(myInt2(num1).Sq())
	fmt.Println(myInt2(num2).Sq())
	fmt.Println(num3.Sq())
	fmt.Println(myInt2(num4).Sq())

	fmt.Println(myInt3(num1).Cube())
	fmt.Println(myInt3(num2).Cube())
	fmt.Println(myInt3(num3).Cube())
	fmt.Println(num4.Cube())

	// float1
	fmt.Println(myInt1(float1).ToString())
	fmt.Println(myInt2(float1).Sq())
	fmt.Println(myInt3(float1).Cube())

	//var b bool
	//myInt1(b).ToString() // cannot convert from bool to myInt1 or myInt2 or myInt3
}

type myint = int // alias

type myInt1 int // This is not inheritance.. myInt1 is a new type

func (mi myInt1) ToString() string {
	return fmt.Sprint(mi)
}

type myInt2 myInt1

func (mi myInt2) Sq() int {
	return int(mi * mi)
}

type myInt3 myInt2

func (mi myInt3) Cube() int {
	return int(mi * mi * mi)
}
