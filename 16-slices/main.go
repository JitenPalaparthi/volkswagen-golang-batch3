package main

import "fmt"

func main() {

	var slice1 []int // dont give the size/length
	// slice1 --> Ptr:nil Len:0 Cap:0

	if slice1 == nil { // The pointer in the slice structure/header is checked
		println("yes it is nil")
	}

	slice1 = make([]int, 5, 10) // instantiating the slice, allocating memory to the slice
	println(slice1)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2], slice1[3], slice1[4] = 3, 4, 5

	// Ptr: 0x1122 Len:5 Cap:10

	DoubleSlice(slice1) //  Ptr: 0x1122 Len:5 Cap:10
	fmt.Println(slice1) // 1,4,9,16,25 Ptr: 0x1122 Len:5 Cap:10
	//slice1 = AppendAndDoubleSlice(6, slice1) // Ptr: 0x1122 Len:6 Cap:10
	AppendAndDoubleSliceP(6, &slice1)
	fmt.Println(slice1) // Ptr: 0x1122 Len:5 Cap:10
	// 1 16 81 256 625

	slice2 := slice1  // header copy slice2 and slice1 have same headers
	slice3 := &slice1 // just a pointer pointing to slice1
	slice2[0] = 0
	slice2[1] = 1
	slice2[2] = 4
	fmt.Println(slice1)
	// slice2 len 6 cap 10
	slice2 = append(slice2, 10, 11, 12, 13, 14, 15) // len 12 cap 20
	slice2[0] = 99999
	slice2[1] = 88888
	fmt.Println(slice1)
	// 0 1 4 256 625 36
	*slice3 = append(*slice3, 10, 11, 12, 13, 14, 15) // len 12 cap 20
	(*slice3)[0] = 99999
	(*slice3)[1] = 88888
	fmt.Println(slice1)
}

func DoubleSlice(slice []int) {
	for i, v := range slice {
		slice[i] = v * v
	}
}

func AppendAndDoubleSlice(elem int, slice []int) {
	// The header of the slice got chagned but not the orign slice
	slice = append(slice, elem)
	for i, v := range slice {
		slice[i] = v * v
	}
}

func AppendAndDoubleSliceR(elem int, slice []int) []int {
	// The header of the slice got chagned but not the orign slice
	slice = append(slice, elem)
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}

func AppendAndDoubleSliceP(elem int, slice *[]int) {
	// The header of the slice got chagned but not the orign slice
	*slice = append(*slice, elem)
	for i, v := range *slice {
		(*slice)[i] = v * v
	}
}

func PrintHeader(name string, slice []int) {
	println("Address of ", name, ":", &slice)
	for _, v := range slice {
		print(v, " ")
	}
	println()

	if slice != nil && len(slice) > 0 {
		println("Ptr of the ", name, ":", &slice[0])
		println("Len of the ", name, ":", len(slice))
		println("Cap of the ", name, ":", cap(slice))
	}
	println()
}
