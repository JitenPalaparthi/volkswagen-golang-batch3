package main

import (
	"time"
)

func main() {
	ch := make(chan int) // unbuffered channel
	sig := make(chan struct{})
	//sig := make(chan Empty)
	go GenFib(50, ch)
	go Receive(ch, sig)

	<-sig // blocking main.. how long until the receiver receives all the values from the ch

	// how can I block the main ?
}

type Empty struct{} //named empty structure

func GenFib(c int, ch chan<- int) {
	a, b := 0, 1
	for i := 1; i <= c; i++ {
		ch <- a
		time.Sleep(time.Millisecond * 50)
		a, b = b, a+b
	}
	close(ch) // only the sender can close the channel

}

func Receive(ch <-chan int, sig chan<- struct{}) {
	//v, ok := <-ch // bool ok tells whether the channel is closed or not
	for v := range ch { // range loop tries to receive value from the channel until it is closed
		println(v)
	}
	sig <- struct{}{}
}
