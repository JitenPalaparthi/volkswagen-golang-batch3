package main

import (
	"math/rand/v2"
)

func main() {
	println("\nclassical loop")

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println("\nwhile loop kind of")

	i := 1
	for i <= 10 { // like while loop
		if i%2 == 0 {
			print(i, " ")
		}
		i++
	}
	println("\nloop with outside init")

	i = 1

	for ; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}

	println("\nloop with inside condition")

	for j := 1; ; j++ {
		if j > 10 {
			break // two areas break can be used 1.loop 2.switch
		}
		print(j, " ")
	}

	println("\nloop with multiple init and post conditions")

	for i, j := 1, 10; i < j; i, j = i+1, j-1 {
		println(i, "-->", j)
	}

	println("\n normal nested loops with inner loop break")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				println("break")
				break
			}
			println(i, "-->", j)
		}
	}

	// println()
	// done := false
	// for i := 1; i <= 5; i++ {
	// 	if done {
	// 		break
	// 	}
	// 	for j := 1; j <= 5; j++ {
	// 		if i == 3 {
	// 			done = true
	// 			println("break")
	// 			break
	// 		}
	// 		println(i, "-->", j)
	// 	}
	// }
	println("\nnested loops break lable")

outer:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				println("break")
				break outer
			}
			println(i, "-->", j)
		}
	}

	println("\nusing goto lable")

	c := 0
loop:
	c++
	if c%3 == 0 {
		print(c, " ")
	}
	if c <= 10 {
		goto loop
	}
	//else {
	// 	goto done
	// }

	// done:
	//
	//	println("Done with the loop")
	println("\n loop and switch break-label")
end:
	for {
		num := rand.IntN(100)
		switch {
		case num%2 == 0:
			println("div by 2", num)
			fallthrough
		default:
			if num%5 == 0 {
				println("div by 5", num)
				break end
			}
		}
	}
}

// loops
