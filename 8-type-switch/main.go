package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	any1 := any("Hello World")
	ok := IsNumber(any1)

	if ok {
		fmt.Println(any1, "is a number type")
	} else {
		fmt.Println(any1, "is not a number type")
	}

	var a, b float32 = 12.34, 98.67
	if r, err := SumOf(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum Of:%.3f\n", r)
	}

	var a1, b1 int = 12, 98
	if r, err := SumOf(a1, b1); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum Of:%.3f\n", r)
	}

	if r, err := SumOf(int8(123), int8(3)); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum Of:%.3f\n", r)
	}

	if r, err := SumOf(int8(123), uint8(3)); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum Of:%.3f\n", r)
	}

	if r, err := SumOf(67.45, 67); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum Of:%.3f\n", r)
	}

	if r, err := SumOf(true, true); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Sum Of:%.3f\n", r)
	}

	r := SumOfG(12, 13) // The code is generated for int
	fmt.Printf("Sum Of:%.3f\n", r)

	r = SumOfG(12.123, 13.123) // the code is generated for float64
	fmt.Printf("Sum Of:%.3f\n", r)

	// r = SumOfG(true, false)
	// fmt.Printf("Sum Of:%.3f\n", r)

}

func IsNumber(a any) bool {
	switch a.(type) { // type switch
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func SumOf(any1, any2 any) (float64, error) {
	if reflect.TypeOf(any1) != reflect.TypeOf(any2) {
		return 0, fmt.Errorf("both of the variables are not of same type")
	}

	if !IsNumber(any1) {
		return 0, errors.New("input arguments are not of a number type")
	}

	sum := 0.0 // float64(0)

	switch v1 := any1.(type) {
	case int:
		sum = float64(v1 + any2.(int))
	case uint:
		sum = float64(v1 + any2.(uint))
	case int8:
		sum = float64(v1 + any2.(int8))
	case uint8:
		sum = float64(v1 + any2.(uint8))
	case int16:
		sum = float64(v1 + any2.(int16))
	case uint16:
		sum = float64(v1 + any2.(uint16))
	case int32:
		sum = float64(v1 + any2.(int32))
	case uint32:
		sum = float64(v1 + any2.(uint32))
	case int64:
		sum = float64(v1 + any2.(int64))
	case uint64:
		sum = float64(v1 + any2.(uint64))
	case float32:
		sum = float64(v1 + any2.(float32))
	case float64:
		sum = v1 + any2.(float64)
	}
	return sum, nil
}

func SumOfG[T int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a, b T) float64 {
	return float64(a + b)
}
