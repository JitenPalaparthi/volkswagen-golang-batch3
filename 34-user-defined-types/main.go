package main

import (
	"fmt"
	"unsafe"
)

// order or fields in a struct is very imp
// have to reduce the padding so that the size of the struct can be in a limit
// if the order is not properly taken care, the size may be high

func main() {
	p1 := Person{}
	fmt.Println(p1, "size of p1:", unsafe.Sizeof(p1))

	e1 := Employee{}
	fmt.Println(e1, "size of e1:", unsafe.Sizeof(e1))
	t1 := T1{}
	fmt.Println(t1, "size of t1:", unsafe.Sizeof(t1))

	t2 := T2{}
	fmt.Println(t2, "size of t2:", unsafe.Sizeof(t2))

}

type Person struct {
	No         uint64 //0 // 8
	IsMarried  bool   //false // 8 -- 7 bytes are padded
	Name       string //"" // 16
	Email      string //"" // 16
	IsEmployed bool   // false // 8 -- 7 bytes are padded
} // 42 bytes

type Employee struct {
	No         uint64 //0 // 8
	IsMarried  bool   //false // 1 byte -- 7 bytes are padded
	IsEmployed bool   // false // 1 byte -- 6 bytes area padded
	B1         bool
	B2         bool
	B3         bool
	B4         byte
	num2       int16
	Name       string //"" // 16
	Email      string //"" // 16
} // 42 bytes

//     No                IsMarried
// ||||||||            	 | |||||||
//						   padding bytes

type T1 struct {
	B1 bool  // 4
	I2 int32 // 4
	B2 bool  // 4
}

// change the order of fields on order to reduce the padded memory
type T2 struct {
	I2 int32 // 4
	B1 bool  // 4
	B2 bool  // 0
}
