package main

import "unsafe"

func main() {

	str1 := "Hello World"
	t1 := T{unsafe.Pointer(&str1), [2]string{"hello", "world"}, true, []any{1, "Hello world"}}
	map1 := make(map[string]any)
	t1.ptr = unsafe.Pointer(&map1)
	slice1 := []int{1, 12, 3}
	t1.ptr = unsafe.Pointer(&slice1)
	t1.ptr = unsafe.Pointer(&t1)
	p := DoSomething
	t1.ptr = unsafe.Pointer(&p)

}

func DoSomething() {
	println("do'in nothing")
}

type T struct {
	ptr   unsafe.Pointer // what address it is ?
	array [2]string
	any1  any
	slice []any
}
