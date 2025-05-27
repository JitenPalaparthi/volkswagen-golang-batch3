package main

import (
	"fmt"
	"time"
)

func main() {
	sig := make(chan struct{})

	ch1 := GenFib(10, "gen-1")
	ch2 := GenFib(10, "gen-2")

	go Receive(ch1, sig)
	go Receive(ch2, sig)

	<-sig // blocking main.. how long until the receiver receives all the values from the ch
	<-sig

	// how can I block the main ?
}

func GenFib(c int, name string) chan string {
	ch := make(chan string) // unbuffered channel
	go func() {
		a, b := 0, 1
		for i := 1; i <= c; i++ {
			ch <- fmt.Sprint(name, "-->", a)
			time.Sleep(time.Millisecond * 50)
			a, b = b, a+b
		}
		close(ch) // only the sender can close the channel
	}()
	return ch
}

func Receive(ch <-chan string, sig chan<- struct{}) {
	//v, ok := <-ch // bool ok tells whether the channel is closed or not
	for v := range ch { // range loop tries to receive value from the channel until it is closed
		println(v)
	}
	sig <- struct{}{}
}
