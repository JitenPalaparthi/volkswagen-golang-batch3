package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	func() { // func-2
		func() { // func2-1
			//num := 0
			//println(100 / num) // divide by zero , once there is a panic , the func would not execute the next statement
			// arr := [5]int{10, 11, 12, 13, 14}
			// for i := 0; i <= len(arr); i++ { // panic , index out of range
			// 	println(arr[i])
			// }
			var ptr *int
			*ptr = 100 // invalid memory address or nil pointer dereference
			println(*ptr)

		}()
		println("Hello func-2")
	}()
	fmt.Println("Hello VolksWagen")
}

//main.func-2.func2-1.println()
