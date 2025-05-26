package main

import (
	"fmt"
	"reflect"
)

func main() {
	map1 := make(map[string]any)
	a, b := 10, 20
	map1["greet"] = func() {
		println("Hello VolksWagen")
	}
	map1["fib"] = func(num uint) {
		a, b := 0, 1
		for i := 1; i <= int(num); i++ {
			print(a, " ")
			a, b = b, a+b
		}
		println()
	}

	map1["sum"] = func(a, b int) int {
		return a + b
	}
	map1["sub"] = sub

	map1["calc"] = Calc(a, b)

	any1 := any("Hello World")
	map1["iface"] = any1

	for k, v := range map1 {
		switch v1 := v.(type) {
		case func():
			fmt.Println(k, "-->", reflect.TypeOf(v))
			v1()
		case func(int, int) int:
			fmt.Println(k, "-->", reflect.TypeOf(v))
			fmt.Println(v1(a, b))
		case func(uint):
			fmt.Println(k, "-->", reflect.TypeOf(v))
			v1(5)
		case func(int, int) func() int:
			fmt.Println(k, "-->", reflect.TypeOf(v))
			fn := v1(a, b)
			r := fn()
			println(r)
		case string:
			fmt.Println(k, "-->", reflect.TypeOf(v))
			fmt.Println(v1)
		default:
			fmt.Println(k, "-->", reflect.TypeOf(v))
		}

	}

}

func sub(a, b int) int {
	return a - b
}

func Calc(a, b int) func() int {
	return func() int {
		return a + b
	}
}
