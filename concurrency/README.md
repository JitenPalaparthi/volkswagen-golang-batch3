Go runtime

Go Scheduler
 Go routines--> Small green threads
 Go routines are managed by Go runtime rather than Os System

The scheduling of Goroutines are done by Go Runtime rather than depending on OS

Go Routines are small in size, with initial size of ~2kb

 Go runtime uses threads
The numbers of threads those are created by Goruntime is by default equal number of cores(12 threads on my machine)

Multipluxing is between Gorotuines: Threads(OS)

Technically few thousand(I even ran a million)gorouintes are scheduled on few number of OS Threads



