package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	defer println("\nend of Main, ofcourse it is crashed here still")
	func() {
		defer RecoverThis()
		defer func() {
			println("\nend of func1")
			println("\n2-end of func1")
		}()
		a, b := 0, 1
		for i := 1; i <= 1000; i++ {
			if a > math.MaxInt || a < 0 {
				panic("fib value is beyond the range") // custom panic
			}
			print(a, " ")
			a, b = b, a+b
		}

		defer println("defer after panic in func1")
	}()

	func() {
		sum := 0
		arr := [10]int{}
		for i := 1; i < 10; i++ {
			arr[i] = rand.IntN(1000)
			sum += arr[i]
		}
		fmt.Println("Arr", arr)
		fmt.Println("Sum:", sum)
	}()
}

func RecoverThis() {
	if r := recover(); r != nil {
		fmt.Println(r) // just printing
		// can fetch data and perform any operation on it.may be logging, retry or anything. that to be done
		// r is any so you can type assert and fetch the data from r
		// similar to type assertion done in custom error example
	}

}

// Where is panic? -- line no 10
// Who is the caller of the panic? main
// before panic-? the run time checks are there any deffered calles in the call stack of the caller, if yes execute them
// before going to crash the application
// whether or not there is a panic, deferred calles are executed
// defer calles only before the panic not after panic

// recover--> checkes if there any panic in the callestack , if yes then recover from it, so that the whole call stack is not hampered
