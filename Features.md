
## Why Go

- Fast Compilation
- Simple Language (25 keywords)
- Not a OOPS, not a functional, not a procedural but Hybrid
- General purpose and systems programming language
- TinyGo(for embedded)
- Cloud native applications

- Cross compilation
- Static Binaries 
- No additional runtime (no jvm kind of)
- type safe (statically typed)(all types should be known to the compiler)
- fast startup times 
- GC (concurrent GC)
- GC pauses (1-10 ms)
        - 200 milliseconds to process/request --> 10 requests/sec 100-200 ms are taken by the GC
- Binaries are statically linked (not dynamically linked to libc/musl/glibc stc)
- Concurrency(go routines) is the first class citizen , CSP 
- Better performers
- Supports WASM/WAPI (Web Assembly)

## Where Go lacks?

- C Interoparabilty is slow (It is not a go's problem still)
- When compare to Rust/Zig, go binaries are little big
- Some claim that error handling is not that great(I personally prefer It is good)
- Not matured enough for mobile and desktop development
- Cant write Operating System(there are alternatives but not in pure Go)
- No control over memory (Heap, Stack etc..)
- No control over OS threads (from developer perspective)
- Some claim that there is Thread contention, particularly using Mutex
- Unless frequently used go channels are quite complex to use (subjective to the developer)
- Does not follow OOPS principles (Some write Java code in Go), need to know the idiamatic approach of Go
- Go lacks AI and ML

## Install

- installl Go from go.dev
- install VSCode 
- install extentions (after installing go and vscode)

