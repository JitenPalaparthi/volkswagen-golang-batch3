package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 uint = 9878798 //1001011010111101 0000 1110
	fmt.Printf("num1:%b\n", num1)
	var num2 uint8 = uint8(num1) //  lose some bits here if type cast

	var num3 uint8 = 0b00001110 // can give a value in binary format
	var num4 uint8 = 0x11       // give the value in xexa format

	fmt.Println(num2, num3)

	var num5 uint64 = uint64(num3) // there is not implicit type casting
	fmt.Println(num4, num5)

	var float1 float32 = 1231.12312
	var num6 int = int(float1) // data is lost.. the data after . is lose
	println(num6)

	//var ok1 = true
	//var num7 uint8 = uint8(ok1) // a bool cannot be casted to uint8

	var char1 int32 = 65
	println(string(char1))

	var char2 rune = '🎯'
	fmt.Println(string(char2))

	fmt.Println(string(int(float1))) // cannot directly cast float32/64 to string

	var str1 = "1234"
	//var num8 int = int(str1) // casting does not work
	num8, err := strconv.Atoi(str1)
	if err != nil {
		println(err.Error())
	} else {
		println(num8)
	}

	a1, b1, c1, d1 := Calc(10, 20)
	println("Addition:", a1, "Substraction:", b1, "Multiplication:", c1, "Division:", d1)
	_, _, c2, _ := Calc(10, 20)
	println("Multiplication:", c2)

	var num9 int = 12312312

	str2 := strconv.Itoa(num9)
	fmt.Println("Value of str2:", str2, "Type of str2:", reflect.TypeOf(str2))
	str3 := fmt.Sprint(num9, true, 12321.12321, "Another strings")
	fmt.Println("Value of str3:", str3, "Type of str3:", reflect.TypeOf(str3))

	str4 := "2312312.123123"

	float2, err := strconv.ParseFloat(str4, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%.6f\n", float2)
	}

	str5 := "Yes"
	ok1, _ := strconv.ParseBool(str5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		println(ok1)
	}
	//  It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	// Any other value returns an error.

	str6 := "ffff"
	num10, err := strconv.ParseInt(str6, 16, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num10)
	}

	c3 := 123.23 + 9i // complex128
	var r1, img1 float32 = 123.23, 9
	c4 := complex(r1, img1) // complex64
	c5 := complex64(c3) + c4
	fmt.Println(c5)
}

func Calc(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// Type Casting
// Implicit and Explicit, Go does not have implicit type casting
// Type Conversion
// Type Assertion

// all number types can be type casted among them
// yes, there might be loss of data

// any number type can be type casted among themself
// any int can be casted to string
// a bool cannot be type casted to any other data type
// a string cannot be type casted to int or float or bool, but type converted

func BoolToInt(k bool) uint8 {
	if k {
		return 1
	}
	return 0
}
