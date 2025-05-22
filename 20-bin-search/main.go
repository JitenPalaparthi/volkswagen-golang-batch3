package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{10, 12, 13, 45, 67, 89, 91, 100, 120}
	i := BinSearch(slice1, 100)
	fmt.Println("Index:", i)
	slice2 := []int{100, 112, 13, 1, 89, 23, 45, 67, 89, 4, 91, 100, 111, 120}
	sort.Ints(slice2)
	fmt.Println(slice2)
	i = BinSearch(slice2, 89)
	fmt.Println("Index:", i)

	i = BinSearch(nil, 89)
	fmt.Println("Index:", i)

}

// slice must be sorted
func BinSearch(slice []int, find int) int {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == find {
			return mid
		} else if slice[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
