package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("End of main")
	fmt.Println("Start of main")

	// f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	// if err != nil {
	// 	println(err.Error())
	// }
	// defer f.Close() //1. I dont forget to close the file

	// n, err := f.Write([]byte("Hello World"))
	// if err != nil {
	// 	println(err.Error())
	// }
	// if n > 0 {
	// 	println("Successfully written")
	// }
	n, err := WriteToFile("data.txt", "Hello VolksWagen!")
	if err != nil {
		println(err.Error())
	}

	if n > 0 {
		println("Successfully written")
	}
}

func WriteToFile(filename string, msg string) (int, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close() //1. I dont forget to close the file
	return f.Write([]byte("Hello World"))
}

// defer
// defer is a keyword
// defer deferes the execution to the end of the call stack
// main is the caller, it is executed to the end of main
