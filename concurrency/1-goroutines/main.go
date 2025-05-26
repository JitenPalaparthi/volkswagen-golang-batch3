package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() { // Func-2
		println("Hello World!")
	}()

	go GreetVW()

	go func() {
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				time.Sleep(time.Millisecond * 200)
				println(i)
			}
		}
		runtime.Goexit()
	}()

	fmt.Println("End of main")

	//time.Sleep(time.Millisecond * 10) // sleep is not a solution

	runtime.Goexit()

}

func GreetVW() {
	println("Hello VolksWagen!")
}

// go routine is nothing but a green thread.
// go rutines are managed by go runtime

// 1. main is also a goroutine
// 2. no goroutine waits for other goroutines to complete its execution
// 3. there are no parent and child go routines
// 4. The order of execution is not guaranteed

// Totally 12 threads are created in my macbook bcz of 12 core processor
// These threads are wrapped with some additional data and called a process in go (process is not a os process)
// G -> Go routine
// T -> Thread
// P -> Process ( Thread and Local Queue and more information)

// When ever a go routine is poised to execute,
// Example
// Func-2,GreetVW  are called by Main
// since is Main is a goroutine, main is running on a P
// The caller is Main and assume main is running on P2,
// Func-2 and GreetVW are placed to run on P2
// Each P, contains a local Queue
// Every Go routine that is scheduled on a P , goes to Local Queue of the P
// They run picked from P Local Queue

// Logically all three or many may be running only on P2 but that is not what happen
// 1. There is a work stealing algorithm
// 2. All available P(s) steals work from busy P, in this case P2
// 3. What if, there is a IO call, there is netpoller package is there, the IO call goroutine is
// is registered to netpoller, once it is registerd it is unblocked and waiting for information from epoll from the sytem
// once the epoll notifies, it is put back into Global Queue
// P(s) takes the tasks froim Global Queue and put them into its local queue and execute them.
