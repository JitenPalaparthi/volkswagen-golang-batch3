# Golang Practice Tasks (100 Intermediate to Advanced)

## 🔁 Slices

1. Reverse a slice in place.
2. Remove duplicates from a slice.
3. Rotate a slice by N positions.
4. Merge two sorted slices into one sorted slice.
5. Find the longest increasing subsequence in a slice.
6. Remove all occurrences of an element from a slice.
7. Write a function that chunks a slice into smaller slices.
8. Implement a sliding window on a slice of integers.
9. Create a generic slice filter function using `constraints` (Go 1.18+).
10. Benchmark performance of slice vs array access.

## 🧠 Pointers

11. Swap two integers using pointer dereferencing.
12. Implement a linked list using pointers.
13. Use pointers to modify elements of a slice in a function.
14. Use pointer receivers in structs and explain when to use them.
15. Benchmark pointer vs value receiver methods.
16. Demonstrate escaping vs non-escaping pointers with escape analysis.
17. Use `unsafe.Pointer` to convert between different types.
18. Modify private struct fields using reflection and unsafe.
19. Use pointers to allocate memory for custom buffer management.
20. Build a memory pool for objects using pointer reuse.

## 🗺️ Maps

21. Create a map of maps and perform nested access.
22. Write a function that merges two maps.
23. Count word frequencies in a paragraph using a map.
24. Use a map to detect duplicates in a slice.
25. Create a concurrent-safe cache using `sync.Map`.
26. Implement a histogram generator using maps.
27. Benchmark `map[string]int` vs `map[int]int`.
28. Use a struct as a map key and explain hashability.
29. Delete all keys in a map without reallocation.
30. Use a map to group words by anagram.

## 🏗️ Structs

31. Implement a constructor function for a struct.
32. Embed multiple structs and resolve field conflicts.
33. Add tags to a struct and use them in JSON marshaling.
34. Implement custom `MarshalJSON` and `UnmarshalJSON`.
35. Use interface embedding to define service contracts.
36. Implement a tree structure using recursive structs.
37. Create an ORM-like struct with CRUD methods.
38. Implement struct field validation using reflection.
39. Compare shallow vs deep copy for nested structs.
40. Simulate inheritance and method overriding with embedding.

## ⛓️ Goroutines & Channels

41. Build a pipeline using multiple goroutines and channels.
42. Use `select` with `time.After` for timeout logic.
43. Implement a fan-in pattern (merging channels).
44. Implement a fan-out pattern (distributing tasks).
45. Build a worker pool using goroutines and buffered channels.
46. Implement rate limiting using `time.Ticker` and goroutines.
47. Create a broadcast system using multiple channel listeners.
48. Detect deadlocks in a multi-goroutine setup.
49. Create a goroutine leak detector.
50. Gracefully stop goroutines using context cancellation.

## ⚙️ Channels - Advanced Patterns

51. Implement a job queue with backpressure using buffered channels.
52. Use channels for asynchronous logging.
53. Implement a priority queue using channels and goroutines.
54. Build a pub-sub (publish-subscribe) system using channels.
55. Use channel-of-channels for dynamic task routing.
56. Monitor channel blocking behavior and add metrics.
57. Create a throttling mechanism using channels.
58. Implement a reusable barrier using channels and goroutines.
59. Build a concurrent rate-limited downloader.
60. Create an event aggregator that buffers bursts using channels.

## 🧪 Concurrency with `sync` and `context`

61. Use `sync.Once` to initialize a singleton.
62. Use `sync.WaitGroup` to wait for multiple goroutines to finish.
63. Use `sync.Mutex` to protect shared resources.
64. Implement `sync.Pool` for object reuse.
65. Use `sync.Cond` for signaling between goroutines.
66. Implement a context-aware timeout operation.
67. Pass `context.Context` across service layers.
68. Build a recursive function with context cancellation support.
69. Create a goroutine-safe counter using atomic operations.
70. Use `sync/atomic` to build a lock-free queue (simple version).

## 🧩 Interfaces and Composition

71. Implement a plugin architecture using interfaces.
72. Create multiple implementations of a service interface.
73. Use type assertion to check dynamic interface types.
74. Implement an adapter pattern using interfaces.
75. Chain decorators using interfaces and embedding.
76. Use interfaces for mocking in tests.
77. Compose behaviors from multiple interfaces.
78. Create a dynamic command handler using `map[string]func()`.
79. Dynamically register types and methods using reflection.
80. Use reflection to serialize struct to custom format.

## 🔥 High Performance Patterns

81. Benchmark performance of channel vs mutex.
82. Build a lock-free circular buffer.
83. Optimize memory usage in slice-heavy applications.
84. Minimize heap allocation using pre-allocated buffers.
85. Analyze memory and CPU usage with `pprof`.
86. Use goroutines and channels to parallelize map-reduce.
87. Use `sync.Map` for high-throughput read-heavy workloads.
88. Profile and reduce GC pressure in a concurrent app.
89. Implement your own slab allocator for object pooling.
90. Use goroutines for parallel I/O (e.g., file reading/writing).

## 🧭 Design Challenges & Patterns

91. Build a task scheduler that executes jobs at specific times.
92. Implement a retry mechanism with exponential backoff.
93. Build a concurrent-safe message queue.
94. Design a concurrent-safe configuration reloader.
95. Create a microservice handler using goroutines and channels.
96. Write a concurrent TCP server that uses goroutines per client.
97. Design a key-value store with in-memory maps and disk persistence.
98. Build a real-time analytics dashboard using goroutines.
99. Create a file watcher that uses goroutines and channels.
100. Implement event-driven architecture using select-based event loops.