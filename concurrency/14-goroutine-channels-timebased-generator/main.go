package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := make(chan struct{})

	ch1 := GenFib(time.Millisecond*10, "gen-1")

	go Receive(ch1, sig)
	//go Receive(ch2, sig)

	<-sig // blocking main.. how long until the receiver receives all the values from the ch
	//<-sig

	// how can I block the main ?
}

func TimeAfter(d time.Duration) chan struct{} {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(d)
		timeout <- struct{}{}
		close(timeout)
	}()
	return timeout
}

func GenFib(duration time.Duration, name string) chan string {
	ch := make(chan string) // unbuffered channel
	//timeout := time.After(duration)
	timeout := TimeAfter(duration)
	done := false
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		<-timeout
		wg.Done()
	}()
	go func() {
		wg.Wait()
		done = true
	}()

	go func() {
		a, b := 0, 1
		for {
			if !done {
				ch <- fmt.Sprint(name, "-->", a)
				time.Sleep(time.Millisecond * 1)
				a, b = b, a+b
			} else {
				close(ch)
				break
			}
		}
		//close(ch) // only the sender can close the channel
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
