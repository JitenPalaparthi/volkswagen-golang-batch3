package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")
	//fmt.Fprintln(nil, "Hello World")
	fmt.Fprintln(os.Stdout, "Hello World")
	fops := NewFileOps("data.txt")
	fmt.Fprintln(fops, "Hello World")

	var fops2 FileOps

	_, err := fmt.Fprintln(&fops2, "Hello VolksWagen")

	switch fo := err.(type) {
	case *FileError:
		fmt.Println("it is FileError object")
		fmt.Println("The error code is:", fo.Code)
		fmt.Println("The error message is:", fo.Msg)
	default:
		fmt.Println(fo.Error())
	}

	//io.Writer
}

type FileOps struct {
	FileName string
}

type FileError struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *FileError {
	return &FileError{code, msg}
}

func (fe *FileError) Error() string {
	return fmt.Sprintf("Code:%d Message:%s", fe.Code, fe.Msg)
}

// in Go no need to explicitly tell that the type is implementing
// any interface. If it implements the method then it is autoimplemented
func (fo *FileOps) Write(p []byte) (n int, err error) {
	if fo == nil {
		//return 0, errors.New("nil object")
		return 0, NewError(101, "nil file object")
	}
	if fo.FileName == "" {
		//return 0, errors.New("invalid or empty file name")
		return 0, NewError(102, "invalid or empty file name")
	}
	f, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close() // will explain defer later

	return f.Write(p)
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

// pure functions --> c++
// what is an interface --> go, java, cshapr, python, type script
// protocol --> closure
// trait --> trait

// interface is a shared behavior, contract, agreement
// multiple parties can adhead to a common implmenataion
// a specification

/*
4 - read
2- Write
1 - execute
---------
7 (RWX)

0777 (O means it is a file,7 current user(RWX), 7(anybody from current Group(RWX)), Others(RWE))
*/
