package main

import "math"

func main() {
	defer println("\nend of Main, ofcourse it is crashed here still")
	func() {
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
}

// Where is panic? -- line no 10
// Who is the caller of the panic? main
// before panic-? the run time checks are there any deffered calles in the call stack of the caller, if yes execute them
// before going to crash the application
// whether or not there is a panic, deferred calles are executed
// defer calles only before the panic not after panic
