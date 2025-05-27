package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//sig := make(chan struct{})

	ch1 := GenFib(10, "gen-1")
	ch2 := GenFib(10, "gen-2")
	ch3 := GenFib(10, "gen-3")
	ch4 := GenFib(10, "gen-4")

	inChan := Receive(ch1, ch2, ch3, ch4)

	for c := range inChan {
		println(c)
	}
	//	<-sig
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

// func Receive(ch <-chan string, sig chan<- struct{}) {
// 	//v, ok := <-ch // bool ok tells whether the channel is closed or not
// 	for v := range ch { // range loop tries to receive value from the channel until it is closed
// 		println(v)
// 	}
// 	sig <- struct{}{}
// }

func Receive(chs ...chan string) chan string {
	//sig := make(chan struct{})
	inChan := make(chan string, len(chs)) // not to clock the merge func so gave the buffer

	wg := new(sync.WaitGroup)

	go func() {
		wg.Wait()
		close(inChan)
		//sig <- struct{}{}
	}()

	for _, ch := range chs {
		wg.Add(1)
		go func() { //merge func
			for c := range ch {
				inChan <- c
				//println(c)
			}
			wg.Done()
		}()
	}

	return inChan
}
