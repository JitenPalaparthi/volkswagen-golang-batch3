package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

const SIZE = 5 // Data Segment ro

var GSIZE = 5 // Data Segment

var BCSIZE int

func main() {
	var arr1 [5]int
	// array legth should be fixed
	// array also has zero values
	// what is the type of an array -> [5]int
	// where does array store -> ? in Heap or Stack
	// cannot cast between two different arrays types
	// unlike string array does not have any header

	fmt.Println(arr1)
	fmt.Println("Type of an array:", reflect.TypeOf(arr1))

	arr1[0], arr1[1], arr1[2], arr1[3], arr1[4] = 101, 102, 103, 104, 105

	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	println("Sum:", sum)

	// < 1.22
	for i := 0; i < len(arr1); i++ {
		println(arr1[i])
	}

	arr2 := [5]int{10, 20, 30, 40, 50}
	arr3 := [...]int{10, 20, 30, 40, 50, 60} // the size is evaluated at comptime
	fmt.Println(arr2, arr3)
	//size := 5
	var arr4 [SIZE]int

	for i := 0; i < len(arr4); i++ {
		arr4[i] = rand.IntN(100)
	}
	fmt.Println(arr4)

	var arr5 [3]int

	//arr5 = ([3]int)(arr4) // cannot type cast

	// can copy but cannot type cast
	for i := 0; i < min(len(arr4), len(arr5)); i++ {
		arr5[i] = arr4[i]
	}
	fmt.Println(arr5)

	fmt.Println("Sum Of:", SumOf(arr1))
	fmt.Println("Sum Of:", SumOf(arr2))
	//fmt.Println("Sum Of:", SumOf(arr5))
}

// we seldom use arrays as the parameters

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
