package main

import (
	"fmt"
	"reflect"
	_ "time" // would like to use for future .. or would like to call intenal init functs.
	"unsafe"
)

// in text/code or RO Data segment
const PI float32 = 3.14
const (
	MAX  = 9999
	MIN  = 0
	MINS = 60 * 60
	HOUR = 60 * MINS // expression
)

var (
	GlobalInt1 = 1111 // Data Segment, initialised data
)
var GlobalInt2 int // BSS, this is uninitialised data

const GLOBALCONST = 100*HOUR + MIN + MAX/20 // what ever it does but possible

func main() {
	var global int = 100             // stack memory
	var str1 = "Hello VolksWagon >🧩" // The str is in stack but the original data is in RO Data Segment
	println(global, str1)

	var str2 = "Hello Universe!, how are you doing >🧩"
	var str3 = str1 + str2
	fmt.Println("Size of str1:", unsafe.Sizeof(str1), "Len of str1 in bytes:", len(str1))
	fmt.Println("Size of str2:", unsafe.Sizeof(str2), "Len of str2 bytes:", len(str2))
	fmt.Println("Size of str3:", unsafe.Sizeof(str3), "Len of str3 bytes:", len(str3))

	var str4 string // Ptr:nil Len:0
	//if str4 == nil { // not a valid operation, on strings dont check nil

	//}
	str4 = "Hello World"    // Ptr: 0x1000a2 Len:11
	str4 = "Hello Universe" // Ptr: 0x000f4  Len: 14
	str5 := "Hello World"   // Ptr: 0x1000a2 Len:11
	str6 := "Hello World"
	str6 = "Hello Another World"
	//	str7 := "" //  Ptr:nil Len:0
	// int8 := 10 // int (8 bytes on 64bit)
	// var num1 uint8
	// num1 = 255
	println(str4, str5, str6) // ""
	println("----")

	//var any1 interface{} = 100 // interface{}
	var any1 any // a special type still it is statically typed
	// what is the zero value of any -- nil
	fmt.Println("Value of any", any1, "Type of any:", reflect.TypeOf(any1)) //
	any1 = true
	fmt.Println("Value of any", any1, "Type of any:", reflect.TypeOf(any1)) //
	any1 = "Hello World"
	fmt.Println("Value of any", any1, "Type of any:", reflect.TypeOf(any1)) //

	any1 = MAX
	fmt.Println("Value of any", any1, "Type of any:", reflect.TypeOf(any1)) //

	any1 = 5645234.324
	fmt.Println("Value of any", any1, "Type of any:", reflect.TypeOf(any1)) //

	any1 = nil
	fmt.Println("Value of any", any1, "Type of any:", reflect.TypeOf(any1)) //

	c1 := 123.23 + 1.4i // complex128
	fmt.Println(c1)
	c2 := complex(123.23, 1.4)
	var r1, im1 float32 = 123.23, 1.4
	c3 := complex(r1, im1)

	c4 := c1 + c2
	c5 := c1 - c2
	c6 := c1 * c2
	c7 := c1 / c2

	r2 := real(c6)
	im2 := imag(c6)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, r2, im2)

	a, b := 10, 20
	// t := a
	// b = t
	// a = b
	a, b = b, a // can be done in Go
	println(a, b)
	a, b, c := 10, 20, 30
	a, b, c = c, a, b // can swap more than one variable
	println(a, b, c)

	_, d, _, _ := 10+10-10/1, 20+20-20, "", true && true || false // _ blank identifier, the right hand side is an expression
	println(d)
	_ = 500 // does nothing
}

// CSP by Tony Hoare in late 1970 or early 80
// null for c++
// The billion dollar mistake I have done was inventing null

/* String - Str = "Hello World"
-------------------------------
String header
Ptr --> Pointer of the original data , the size of a pointer on 64bit is 8 bytes
Len --> Length of that data, the size of the integer on 64bit is 8 bytes
*/

// where ever you see a pointer, directly or indirectly(in the header), can check whether the variable
// is nil or not, this rule is not for string

/*any Header
DataPtr : Pointer to the actual data // 8 bytes
TypePtr : Pointer to the type of the data // 8 bytes
*/
