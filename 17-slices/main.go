package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum slice1:", SumOf(slice1))
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr1[:] // convert an array to the slice
	// a new slice header is created with pointer of the array and len and cap are given
	fmt.Println("Sum slice1:", SumOf(arr1[:]))
	slice2[0] = 9999
	fmt.Println(arr1)
	slice2 = append(slice2, 1111)
	slice2[1] = 8888
	fmt.Println(arr1)
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
