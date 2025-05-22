package main

var Global1 int
var Global2 int = 9999

func main() {
	arr1 := [100000000]int{}
	// fmt.Println(arr)

	var slice1 []int // dont give the size/length
	// slice1 --> Ptr:nil Len:0 Cap:0

	if slice1 == nil { // The pointer in the slice structure/header is checked
		println("yes it is nil")
	}

	slice1 = make([]int, 5) // instantiating the slice, allocating memory to the slice
	println(slice1)

	slice1[0] = 100
	slice1[1] = 200
	slice1[2], slice1[3], slice1[4] = 300, 400, 500
	println("Sum of slice1", SumOf(slice1))

	slice2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	println("Sum of slice2", SumOf(slice2))

	ptr := IncrPI(100)
	println(*ptr)

	println("Address of arr1:", &arr1[0])
	println("Address of Global1:", &Global1)
	println("Address of Global2:", &Global2)
	PrintHeader("slice1", slice1)
	PrintHeader("slice2", slice2)

}

func PrintHeader(name string, slice []int) {
	println("Address of ", name, ":", &slice)
	for _, v := range slice {
		print(v, "")
	}
	println()

	if slice != nil && len(slice) > 0 {
		println("Ptr of the ", name, ":", &slice[0])
		println("Len of the ", name, ":", len(slice))
		println("Cap of the ", name, ":", cap(slice))
	}
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// dangling pointer
func IncrPI(num int) *int {
	ptr := new(int) // creating a ptr
	*ptr = num * num
	println("address of ptr:", &ptr)
	return ptr
}

// make is built in function to instantiate the slice
// slice can be nil, the internal ptr is checked
// slice has zero values
/*
Slice header
-----------
Ptr: nil
Len: 0
Cap: 0
*/

/*
Ptr:0x112233
Len:5
Cap:5
&|  |  |  |  |  |

*/
