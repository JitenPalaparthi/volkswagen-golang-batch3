package main

import (
	"runtime"
	"sync"
	"time"
)

func main() {

	ch := make(chan int) // unbuffered channel
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go GenFib(50, ch, wg)
	wg.Add(1)
	go Receive(ch, wg)
	wg.Wait()
}

func GenFib(c int, ch chan<- int, wg *sync.WaitGroup) {
	a, b := 0, 1
	for i := 1; i <= c; i++ {
		ch <- a
		time.Sleep(time.Millisecond * 50)
		a, b = b, a+b
	}
	close(ch) // only the sender can close the channel
	wg.Done()
}

func Receive(ch <-chan int, wg *sync.WaitGroup) {
	//v, ok := <-ch // bool ok tells whether the channel is closed or not
	for {
		v, ok := <-ch // only the receiver can know whether the channel is closed or not

		if !ok { // the channel is closed
			wg.Done()
			runtime.Goexit()
		}
		println(v)
	}
}
