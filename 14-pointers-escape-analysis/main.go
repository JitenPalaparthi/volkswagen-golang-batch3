package main

import "log"

func main() {
	num := 100
	Incr(num)
	println(num)

	num = 100
	Incr(num)
	println(num)

	num = IncrR(num)
	println(num)

	IncrP(&num)
	println(num)

	ptr := IncrPI(num)
	println(*ptr)

	arr1 := [100]int{}
	//arr1[0] = 111
	log.Println(arr1) // should not be used in production

	arr2 := [1000000]int{}
	arr2[999999] = 111

}

// nil/null pointer dereference --> should be taken care even in Go
// dangling pointer --> go compiler/runtime deals it by moving the variable to Heap memory
// double free --> GC handles it. Memory is freed by GC so no point of double free
// memory leak --> allocating on heap but not deallocating (Go runtime-->GC)

func Incr(num int) {
	num++
}

func IncrR(num int) int {
	num++
	return num
}

func IncrP(num *int) { // call ref
	if num != nil {
		*num += 1
	}
}

// dangling pointer
func IncrPI(num int) *int {
	ptr := new(int) // creating a ptr
	*ptr = num * num
	return ptr
}
