package main

import (
	"fmt"
	"reflect"
)

func main() {

	// func(params)returntype{
	// body of the function
	// }(executor)
	func() {
		println("Hello VolksWagen")
	}()

	var fn1 func(uint) // a type a function can also be a variable type

	// var struct1 struct{ name string }
	// var iface1 interface{}
	fn1 = func(num uint) {
		a, b := 0, 1
		for i := 1; i <= int(num); i++ {
			print(a, " ")
			a, b = b, a+b
		}
	}

	if fn1 == nil {
		println("Yes fn1 is nil")
	}

	fmt.Println("Type of fn1:", reflect.TypeOf(fn1))
	fn1(10)
	println()
	func(num uint) {
		a, b := 0, 1
		for i := 1; i <= int(num); i++ {
			print(a, " ")
			a, b = b, a+b
		}
	}(20)

}
