package main

import "sync"

var Count int = 0 // a shared variable which is not thread safe

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Incr(wg)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Decr(wg)
		}
		wg.Done()
	}()

	wg.Wait()
	println(Count)
}

func Incr(wg *sync.WaitGroup) {
	Count++
	wg.Done()
}
func Decr(wg *sync.WaitGroup) {
	Count--
	wg.Done()
}
