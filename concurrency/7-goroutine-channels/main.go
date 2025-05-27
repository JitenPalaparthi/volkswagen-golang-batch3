package main

func main() {

	var ch chan int // This is nil bcz not instantiated

	if ch == nil {
		println("yes nil channel")
		ch = make(chan int, 2) // This is unbuffered ----------
		//										->	 300 200 100
		//											 ----------

	}

	ch <- 100 // sending data , it is waiting here for the receiver to receive the data
	ch <- 200
	v := <-ch // receiving data
	println(v)
	v = <-ch // receiving data
	println(v)
	ch <- 300
}

// CSP-> Communicating Sequential Processes

// Do not communicate by sharing memory;Share memory by communicating

// there is a separate communicating which is called channel

// channel means flow of something
// flow is always in one direction

// two kinds of channels, buffered and unbuffered
//

// 1. Until the receiver receives the data, the sender is blocked
// 2. until the sender sends the data, the receiver is blocked
// 3. both the sender and receiver should run concurrently, and communicate seamlessly in order not to have the deadlock
// 4. Sender --> Nothing but the goroutine sending data, thru the channel
// 5. Receiver -> Nothing but a goroutine receiving data, thru the channel
// On buffered channel
// 6. On a buffered channel ,  the sender is not blocked until the buffer is full

/*

| |
| | -----------------  20ltr of water
					|
				   | |
					-
Sender --> Over head tank
Receiver --> A tub
Chennel --> Tube
Context --> Flow of water

*/
