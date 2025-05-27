package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string) // unbuffered channel
	sig := make(chan struct{})
	//sig := make(chan Empty)
	wg := new(sync.WaitGroup)

	wg.Add(2)
	go GenFib(10, "gen-1", ch, wg)
	go GenFib(10, "gen-2", ch, wg)

	go Receive(ch, sig)
	go func() {
		wg.Wait() // How long it has to wait ? until the state/count becomes zero
		close(ch)
	}()

	<-sig // blocking main.. how long until the receiver receives all the values from the ch

	// how can I block the main ?
}

func GenFib(c int, name string, ch chan<- string, wg *sync.WaitGroup) {
	a, b := 0, 1
	for i := 1; i <= c; i++ {
		ch <- fmt.Sprint(name, "-->", a)
		time.Sleep(time.Millisecond * 50)
		a, b = b, a+b
	}
	wg.Done()
	//close(ch) // only the sender can close the channel
}

func Receive(ch <-chan string, sig chan<- struct{}) {
	//v, ok := <-ch // bool ok tells whether the channel is closed or not
	for v := range ch { // range loop tries to receive value from the channel until it is closed
		println(v)
	}
	sig <- struct{}{}
}
