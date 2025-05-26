package main

import (
	"math/rand/v2"
	"runtime"
	"time"
)

func main() {

	go func() {
		for {

			randV := rand.IntN(1000)
			if randV == 999 {
				runtime.Goexit()
			}
			//time.Sleep(time.Millisecond * 100)
			go func(n int) {
				//println("Go Routines--->", runtime.NumGoroutine())
				if n%2 == 0 {
					println("Even:", n)
				} else {
					println("Odd:", n)
				}
			}(randV)
		}
	}()

	time.Sleep(time.Second * 1)

}
