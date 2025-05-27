package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sig := make(chan struct{})

	//ch1 := GenFib(time.Millisecond*5, "gen-1")
	ch1, ch2 := GenValue(time.Second*1, "gen-1")

	go Receive(ch1, sig)
	go Receive(ch2, sig)
	//go Receive(ch2, sig)

	<-sig // blocking main.. how long until the receiver receives all the values from the ch
	<-sig

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
	timeout1 := TimeAfter(duration)
	timeout2 := TimeAfter(duration)

	go func() {
		a, b := 0, 1
		for {
			select { // select is only used for channels
			case <-timeout1: // every case should be a channel
				println("Timeout--1")
				close(ch)
				runtime.Goexit()
			case <-timeout2: // every case should be a channel
				println("Timeout--2")
				close(ch)
				runtime.Goexit()
			// case ch <- fmt.Sprint(name, "-->", a):
			// 	time.Sleep(time.Millisecond * 1)
			// 	a, b = b, a+b
			default:
				println("---------")
				ch <- fmt.Sprint(name, "-->", a)
				time.Sleep(time.Millisecond * 1)
				a, b = b, a+b
			}
		}
		//close(ch) // only the sender can close the channel
	}()
	return ch
}

func GenValue(duration time.Duration, name string) (chan string, chan string) {
	ch1 := make(chan string) // unbuffered channel
	ch2 := make(chan string) // unbuffered channel
	//timeout := time.After(duration)
	timeout1 := TimeAfter(duration)
	timeout2 := TimeAfter(duration)

	go func() {
		a := 1
		for {
			select { // select is only used for channels
			case <-timeout1: // every case should be a channel
				println("Timeout--1")
				close(ch1)
				close(ch2)
				runtime.Goexit()
			case <-timeout2: // every case should be a channel
				println("Timeout--2")
				close(ch1)
				close(ch2)
				runtime.Goexit()
			// case ch1 <- fmt.Sprint(name, "square -->", a*a):
			// 	time.Sleep(time.Millisecond * 100)
			// case ch2 <- fmt.Sprint(name, "cube -->", a*a*a):
			// 	time.Sleep(time.Millisecond * 1)
			default:
				ch1 <- fmt.Sprint(name, " square of ", a, " -->", a*a)
				ch2 <- fmt.Sprint(name, " cube ", a, " -->", a*a*a)
				time.Sleep(time.Millisecond * 100)
			}
			a++
		}
		//close(ch) // only the sender can close the channel
	}()
	return ch1, ch2
}

func Receive(ch <-chan string, sig chan<- struct{}) {
	//v, ok := <-ch // bool ok tells whether the channel is closed or not
	for v := range ch { // range loop tries to receive value from the channel until it is closed
		println(v)
	}
	sig <- struct{}{}
}
