package main

import (
	"sync"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	// go func() { // func-2
	// 	ch <- 100
	// 	println("Sender sends the value")
	// 	wg.Done()
	// }()
	// time.Sleep(time.Second * 5)
	// println(<-ch) // can do this

	go func() {
		println(<-ch) // can do this
		wg.Done()
	}()

	ch <- 200

	wg.Wait()
}
