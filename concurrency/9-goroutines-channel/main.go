package main

import (
	"sync"
	"time"
)

func main() {

	ch := make(chan int) // unbuffered channel
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go GenFib(10, ch, wg)
	wg.Add(1)
	go Receive(10, ch, wg)
	wg.Wait()
}

func GenFib(c int, ch chan<- int, wg *sync.WaitGroup) {
	a, b := 0, 1
	for i := 1; i <= c; i++ {
		ch <- a
		time.Sleep(time.Millisecond * 200)
		a, b = b, a+b
	}
	wg.Done()
}

// HOw does the receiver know the sender is sending number of values ?
func Receive(c int, ch <-chan int, wg *sync.WaitGroup) {
	for i := 1; i <= c; i++ {
		println(<-ch)
	}
	wg.Done()
}
