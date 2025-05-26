package main

func main() {
	map1 := make(map[string]any)

	map1["greet"] = func() {
		println("Hello VolksWagen")
	}
	map1["fib"] = func(num uint) {
		a, b := 0, 1
		for i := 1; i <= int(num); i++ {
			print(a, " ")
			a, b = b, a+b
		}
	}

	map1["sum"] = func(a, b int) int {
		return a + b
	}
	map1["sub"] = sub

	any1 := any("Hello World")
	map1["iface"] = any1

}

func sub(a, b int) int {
	return a - b
}
