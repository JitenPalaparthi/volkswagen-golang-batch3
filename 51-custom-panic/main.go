package main

import "math"

func main() {
	a, b := 0, 1
	for i := 1; i <= 1000; i++ {
		if a > math.MaxInt || a < 0 {
			panic("\nfib value is beyond the range") // custom panic
		}
		print(a, " ")
		a, b = b, a+b
	}
}
