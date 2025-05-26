package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sort"
	"sync"
	"time"
)

var wg *sync.WaitGroup = new(sync.WaitGroup)

func main() {

	wg.Add(1)
	go func() {
		for {
			randV := rand.IntN(1000)
			if randV == 999 {
				wg.Done()
				runtime.Goexit()
			}
			//time.Sleep(time.Millisecond * 100)
			wg.Add(1)
			go func(n int) {
				//println("Go Routines--->", runtime.NumGoroutine())
				if n%2 == 0 {
					println("Even:", n)
				} else {
					println("Odd:", n)
				}
				wg.Done()
			}(randV)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		a, b := 0, 1
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func(n int) {
				println("fib--", n)
				wg.Done()
			}(a)
			a, b = b, a+b
		}
	}()

	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond * 1)
		r := sum(10, 20)
		println("Sum ------->", r)
		wg.Done()
	}()

	//wg.Add(1)
	//go func() {
	arr := [20]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.IntN(1000)
	}
	fmt.Println("Before sort:", arr)

	wg.Add(1)
	go func() {
		sort.Ints(arr[:])
		fmt.Println("After sort:", arr)
		wg.Done()
	}()

	//	wg.Done()
	//}()

	wg.Wait() // wait until the state becomes zero

}

func sum(a, b int) int {
	return a + b
}
