package main

import "fmt"

func main() {

	str := "Hello World"
	println("Hello VolksWagen")

	defer println("End of main-1")

	for _, v := range str {
		defer print(string(v), " ")
	}
	println()
	defer println("end of main-2")

	a := new(int)
	*a = 100
	defer fmt.Println("a deffered one-->", *a) // it is call by value
	*a += 1
	fmt.Println("a normal call-->", *a)

	b := 100

	defer func(*int) {
		fmt.Println("b defer-->", b)
	}(&b) // call by reference

	b = b + 1
	println("b before defer -->", b)

}
