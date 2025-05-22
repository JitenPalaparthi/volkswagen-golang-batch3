package main

import "math/rand/v2"

func main() {

	var num int = rand.IntN(100)
	var ptr *int // does not have any address . What is the size of the pointer--?
	// *ptr = 9999
	if ptr == nil {
		println("ptr is nil")
	}
	ptr = &num // address of num to the pointer
	*ptr = 9999
	//*(&num) = 8888
	println(num)

	var ptr2 *bool // what is the size of the pointer , 8 bytes
	//ptr2 = &true   // valid in rust
	ptr2 = new(bool) // new allocates memory of 1 byte and assignes the address to ptr2
	println(*ptr2)   // zero value false
	*ptr2 = true
	println(*ptr2)

	var ptr3 **int = &ptr
	ptr4 := &ptr3
	***ptr4 = 1111 // dereference the num
	println(num)

	arr1 := [3]int{10, 20, 30}
	ptrArr1 := &arr1
	for _, v := range ptrArr1 { // safe to use range loop even ptrArr is nil
		println(v)
	}

	num1 := 100   //
	Incr(num1)    // call by or pass by value
	println(num1) // ?

	num1 = IncrR(num1)
	println(num1) // ?

	IncrP(&num1)  // pass/call by refernce
	println(num1) // ?

	//str1 := "Hello World"
	// Ptr --> What is this pointer
	// Len
}

// Node
// Data
// ptr *Node

// Pointer holds the address
// what is the zero value of a pointer

// null pointer dereferencing / nil pointer dereference --> should be taken care even in Go
// double free --> ? frees the heap allocation , no need to worry in Go, bcz of GC
// dangling pointer --> ?
// memory leak --> ?

func Incr(num int) {
	num++
}
func IncrR(num int) int {
	num++
	return num
}

func IncrP(num *int) {
	if num != nil {
		*num += 1
	}
}
