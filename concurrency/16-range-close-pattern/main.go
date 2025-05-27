package main

import (
	"fmt"
	"time"
)

func main() {
	ch := Gen(10)
	for c := range ch {
		time.Sleep(time.Millisecond * 100)
		println(c)
	}

}

func Gen(num int) chan int {
	ch := make(chan int, 10)
	go func() {
		for i := range num {
			ch <- i * i
		}
		close(ch)
		fmt.Println("Chennel is closed")
	}()
	return ch
}
