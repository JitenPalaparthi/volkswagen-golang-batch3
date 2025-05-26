package main

func main() {

	var sliceFn1 []func()

	for i := 1; i <= 10; i++ {
		sliceFn1 = append(sliceFn1, func() {
			println(i * i)
		})
	}

	for _, fn := range sliceFn1 {
		fn()
	}

	var sliceFn2 []func()

	i := 1

loop:

	sliceFn2 = append(sliceFn2, func() {
		println(i * i)
	})
	i++
	if i < 10 {
		goto loop
	}

	println("Lazy execution on a non for loop")
	for _, fn := range sliceFn2 {
		fn()
	}

	println("lazy execution solution")

	var sliceFn3 []func(int)

	j := 1

loop1:

	sliceFn3 = append(sliceFn3, func(i int) {
		println(i * i)
	})
	j++
	if j < 10 {
		goto loop1
	}

	println("Lazy execution on a non for loop")
	for i, fn := range sliceFn3 {
		fn(i)
	}

}
