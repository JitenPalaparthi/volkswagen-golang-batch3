package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr := [5]int{10, 11, 12, 13, 14}

	// ptr1 := &arr[0]
	// //ptr1 += 8 // this is what pointer arithmetic, this cannot be done in Go
	// fmt.Printf("0x%x\n", &ptr1)

	// var uintptr1 uintptr = 1374390640688 //0x1400012e000
	// uintptr1 += 8
	// println(uintptr1)
	// fmt.Printf("0x%x\n", uintptr1)
	uintptr1 := uintptr(unsafe.Pointer(&arr[0]))
	for i := 0; i < len(arr); i++ {
		// 1 and 4
		val := *(*int)(unsafe.Pointer(uintptr1)) // 3 and 2
		fmt.Println(val)
		uintptr1 += unsafe.Sizeof(arr[i]) // 8 bytes farword for the original address
	}

	slice1 := []int{}
	uintptr1 = uintptr(unsafe.Pointer(&slice1))
	val := (*[3]int)(unsafe.Pointer(uintptr1)) // 3 and 2
	fmt.Println((*val)[0])
	fmt.Println((*val)[1])
	fmt.Println((*val)[2])
	(*val)[0] = int(uintptr(unsafe.Pointer(&arr[0])))
	(*val)[1] = 5
	(*val)[2] = 5

	fmt.Println(slice1)

	str1 := "Hello Wrold" // len : 11
	//str1 = ""
	//var str1 string //:= "Hello Wrold" // len : 11

	//ptr1 := uintptr(unsafe.Pointer(&str1))
	strPtr := (*[2]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&str1))))
	fmt.Println((*strPtr)[0])
	fmt.Println((*strPtr)[1])
	(*strPtr)[1] = 100

	fmt.Println(len(str1))
}

// 1.  - A pointer value of any type can be converted to a Pointer.
// 2.  - A Pointer can be converted to a pointer value of any type.
// 3.  - A uintptr can be converted to a Pointer.
// 4.  - A Pointer can be converted to a uintptr.
