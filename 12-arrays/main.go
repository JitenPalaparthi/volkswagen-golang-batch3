package main

func main() {

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			print(arr2d[i][j], " ")
		}
		println()
	}
	println()

	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				print(v, " ")
			}
			println()
		}
	}
	//println()

	for _, v := range FibRange() {
		print(v, " ")
	}
	println()

}

func FibRange() [10]int {
	a, b := 0, 1
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = a
		a, b = b, a+b
	}
	return arr
}
