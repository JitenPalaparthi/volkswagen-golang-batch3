package main

func main() {

	str1 := "Hello World!"
	str2 := "💻📦→ Hello, World! ✅🚀"

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]))
	}
	println()

	for i := 0; i < len(str2); i++ {
		println(i, "-->", str2[i], "-->", string(str2[i]))
	}
	println()

	for i, v := range str2 {
		println(i, "-->", string(v))
	}

	for _, v := range str2 {
		print(string(v))
	}

	println()
	// for i, _ := range str2 {
	for i := range str2 {
		print(i, "  ")
	}

	// range loop
	// on a array, string, slice, index and value
	// on a map -> key and value
	// on a channel --> receive value from channel

	println()
	// from 1.22 version
	for i := range 10 {
		print(i, " ")
	}
}

// range loop

//
