package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	pub := GenFib(20, "publisher-1")
	sliceCh := Subscribers(5, pub)
	wg := new(sync.WaitGroup)
	for _, ch := range sliceCh {
		wg.Add(1)
		go func() {
			for c := range ch {
				println(c)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// time.Sleep(time.Second * 5)
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

func Subscribers(subs int, pubCh chan string) []chan string {
	sliceCh := make([]chan string, subs)
	for i := 0; i < len(sliceCh); i++ {
		sliceCh[i] = make(chan string, 10)
	}

	wg := new(sync.WaitGroup)

	for pub := range pubCh { // this is received until closed
		for i, ch := range sliceCh {
			wg.Add(1)
			go func(j int, ch1 chan string, val string) {
				ch1 <- fmt.Sprint("chan--", j+1, "-->", val)
				wg.Done()
			}(i, ch, pub)
		}
	}
	go func() {
		wg.Wait()
		for i := 0; i < len(sliceCh); i++ {
			close(sliceCh[i])
		}
	}()

	return sliceCh
}
