package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := slice[:]   // or slice
	slice2 := slice[:5]  // 0-5 th element but not 5 // len 5 cap 10
	slice3 := slice[3:8] // 3rd index to 8th but not 8 // len 5 cap;7
	slice4 := slice[5:]  // len : 5 cap :5

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	slice4[0] = 9999
	fmt.Println(slice)

	slice5 := append(slice[:5], slice[6:]...) // delete an elemnt from a slice and create a new slice
	fmt.Println(slice5)

	slice6 := append([]int{}, 1, 2, 3, 4, 5) // create a new slice
	fmt.Println(slice6)

	slice7 := []int{} // not nil // Ptr: Dummp Ptr, Len:0 Cap:0
	var slice8 []int  // this is nil

	slice9 := new([]int)
	if slice9 == nil {
		println("slice9 is nil")
	}
	*slice9 = append(*slice9, 10, 11, 12, 13, 14)
	fmt.Println(*slice9)
	// is slice7 nil or not

	if slice7 == nil {
		println("slice7 is nil")
	}

	if slice8 == nil {
		println("slice8 is nil")
	}
}
