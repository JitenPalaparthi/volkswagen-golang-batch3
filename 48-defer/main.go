package main

import "fmt"

func main() {
	defer fmt.Println("End of main") // defers to the end of the call stack, main is the caller
	defer fmt.Println("End of main-1")
	defer fmt.Println("End of main-2")
	fmt.Println("Start of main")

	func() { // func-2
		defer println("end if func-2")
		defer println("end if func-2.1 2")
		defer println("end if func-2.1 3")
		println("start of func-2")

		func() { // func-2.1
			defer println("end if func-2.1")
			println("start of func-2.1")
		}()

		func() { // func-2.2
			defer println("end if func-2.2")
			println("start of func-2.2")
		}()
	}()

	func() { // func-3
		defer println("end if func-3")
		println("start of func-3")

		func() { // func-3.1
			defer println("end if func-3.1")
			println("start of func-3.1")
		}()

		func() { // func-3.2
			defer println("end if func-3.2")
			println("start of func-3.2")
		}()
	}()

}
