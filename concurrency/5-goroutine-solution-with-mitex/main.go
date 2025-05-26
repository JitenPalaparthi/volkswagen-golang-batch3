package main

import "sync"

var Count int = 0 // a shared variable which is not thread safe

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Incr(wg, mu)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Decr(wg, mu)
		}
		wg.Done()
	}()

	wg.Wait()
	println(Count)
}

func Incr(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	Count++
	mu.Unlock()
	wg.Done()
}
func Decr(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	Count--
	mu.Unlock()
	wg.Done()
}
