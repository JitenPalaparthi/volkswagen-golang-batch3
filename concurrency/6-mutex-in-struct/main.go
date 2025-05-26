package main

import "sync"

var Count int = 0 // a shared variable which is not thread safe

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	counter := Counter{count: 0, mu: mu}

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func() {
				counter.Incr()
				wg.Done()
			}()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func() {
				counter.Decr()
				wg.Done()
			}()
		}
		wg.Done()
	}()
	// wg.Add(1)
	// go func() {
	// 	for i := 1; i <= 200; i++ {
	// 		wg.Add(1)
	// 		go func() {
	// 			counter.Println()
	// 			wg.Done()
	// 		}()
	// 	}
	// 	wg.Done()
	// }()

	wg.Wait()
	println(counter.GetCount())
}

type Counter struct {
	count int
	//mu    *sync.RWMutex
	mu *sync.Mutex
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func (c *Counter) Decr() {
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
}

// func (c *Counter) Println() {
// 	c.mu.Lock()
// 	println("current value of Count:", c.count)
// 	c.mu.Unlock()
// }

func (c *Counter) GetCount() int {
	return c.count
}
