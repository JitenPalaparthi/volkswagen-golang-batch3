package main

import (
	"fmt"
	"strconv"
)

func main() {

	var any1 any // nil, nil

	if any1 == nil {
		println("Yes nil")
	}

	any1 = 100        // what is the type --> int
	any1 = uint8(100) // what is the type --> uint8

	// var num1 uint8 = uint8(any1)

	var num1 uint8 = any1.(uint8) // type assertion any.(type)
	println(num1)

	var num2 int = int(any1.(uint8))
	println(num2)

	// var str1 string = any1.(string) // it returns an run time error
	// println(str1)

	var str1 string
	str1, ok := any1.(string) // it returns an run time error
	if ok {
		println(str1)
	} else {
		println("unable to assert to string, so trying with uint8")
		num3, ok := any1.(uint8)
		if ok {
			println(num3)
		} else {
			println("unable to assert to string")
		}
	}

	var (
		num10 int     = 12321          // ok
		num11 uint8   = 32             //ok
		num12 float32 = 123.123        // ok
		num13 float64 = 1232312.123123 //ok
		str4  string  = "123123"       //ok
		ok2   bool    = true           //ok
		any11 any     = 123123.123     //ok
		any2  any     = "123123.12312" //ok
		any3  any     = 98883          //ok
		num14 uint16  = 0b00110011     // binary
		num15 uint64  = 0x112233ffbb   // xexa
	)

	sum := 0.0 // float64

	sum = float64(num10) + float64(num11) + float64(num12) + num13 + any11.(float64) + float64(any3.(int))

	f4, _ := strconv.ParseFloat(any2.(string), 64)
	sum += f4 // even error it addes only 0

	i1, _ := strconv.Atoi(str4)
	sum += float64(i1) // even error it addes only 0

	if ok2 {
		sum += float64(1)
	}

	sum = sum + float64(num14) + float64(num15)
	fmt.Printf("Sum:%.5f\n", sum)
}
