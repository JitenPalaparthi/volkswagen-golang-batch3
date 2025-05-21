package main

import (
	"fmt"
	"reflect"
)

type integer = int // integer is alias for int

func main() {
	var num1 int = 100
	// size of int is based on arch .. on 32bit it is 4 bytes and on 64 bits it is 8 bytes
	var num2 uint64 = 1231231
	var num3 uint8 = 255
	var num4 integer = 99843456867

	var float1 float32 = 21.12323      // float3.4 * 10 power 38 (~ 6-7 decimals)
	var float2 float64 = 2123123.12313 // ~ 15-16 decimals

	var ok1 bool = true
	var str1 string = "Hello World!"

	var char1 rune = 'a'
	var char2 rune = '📦'
	// rune is an alias for int32

	var char3 int32 = 'A'
	var char4 uint8 = 'B'

	var char5 = char3 + char2 + char1 // arithmetic operations

	var byte1 byte = 255 // no special type called byte, it is just a uint8

	println(num1, num2, num3, num4, float1, float2, ok1, char1, char2, char3, char4, char5, str1, byte1)
	fmt.Println(num1, num2, num3, num4, float1, float2, ok1, char1, char2, char3, char4, char5, str1, byte1)
	fmt.Printf("float1: %.2f\nfloat2: %.3f\n", float1, float2)
	fmt.Printf("Type of float1:%T\n", float1)
	fmt.Println("Type of float2:", reflect.TypeOf(float2))

	// Group of variables like this
	var (
		a, b, c int32  // zero value of numbers is 0
		ok2     bool   // zero value of boolean is false
		str2    string // zero value of string is ""
	)

	fmt.Println(a, b, c, ok2, str2)

	var a1 = 987                  // the type is automatically inferred based on the balue
	var f1 = 123.23               // type is inferred as float64
	var ok3 = true                // type is inferred as bool
	var str3 = "Hello VolksWagon" // type is inferred as string

	var age1 uint8 = 42
	var age2 = 42 // age should be a simple uint8 but here it is int

	var pi1 float32 = 3.14
	var pi2 = 3.14 // float64

	// shorthand declaration
	a2 := 123 /// there is a type inference and also no need to use var keyword
	ok4 := true
	str4 := "Hello World"
	fmt.Println(a1, f1, ok3, str3, a2, ok4, str4, pi1, pi2, age1, age2)
}

// type inference --> type is inferred by the compiler
// zero values --? zero value is a value that is automatically given when no value is assigned
