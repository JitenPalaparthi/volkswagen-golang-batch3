package main

import "fmt"

func main() {
	var slice1 []int                                         // Ptr:nil,Len:0,Cap:0
	slice2 := []int{}                                        // Ptr:SomePtr, Len:0, Cap:0
	slice3 := make([]int, 5)                                 // Ptr:allocationPtr, Len:5:Cap:5
	slice4 := append([]int{}, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // Ptr:AllocationPtr, Len:10,Cap:10
	fmt.Println("Len:", len(slice4), "Cap:", cap(slice4))
	slice5 := []int{1, 2, 3} // Ptr:SomePTr, Len:3, Cap:3
	arr := [5]int{1, 2, 3, 4, 5}
	slice6 := arr[:]             // Ptr: PtrOfArr,Len:5,Cap:5
	slice7 := make([]int, 5, 10) // Ptr:allocationPtr, Len:5:Cap:10
	fmt.Println(slice1, slice2, slice3, slice4, slice5, slice6, slice7)
	slice := append([]int{}, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	copy(slice3, slice) // deep copy the destination should not be nil.
	// copy happens only based on  number of elements to be copied
	fmt.Println(slice3)
	clear(slice3) // it clears all the values to zero values
	fmt.Println(slice3)
}
