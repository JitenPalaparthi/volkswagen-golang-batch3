## Go 

- To instantiate the module
    - every project in Go should have a module. It creates go.mod file

```sh
go mod init demo
```
- To run 

```sh
go run main.go
```

- To run , if main is in multiple files

```sh
go run main.go greet.go
```
 
or

```sh
go run .
```

- Go build

```sh
go build .
```

- build with output name

```sh
go build -o hello .
```

## Escape Analysis

```sh
go run -gcflags="-m" main.go
```

## debug mode 

- symbol information and also DWARF information for the debugging purpose 

## Build for the release mode

- This build is for release but not for debugging purpose. You cannot do debugging using tools like dlb or gdb or lldb etc..

```sh
go build -ldflags="-s -w" -o hello .
```

- compile and link 

```sh
go tool compile -o main.o main.go 
go tool link -o demo main.o  
```

- Go env 

```sh
go env
```

GOROOT,GOPATH,GOBIN,GOOS,GOARCH,GOPROXY

# Go Environment Variables Overview

| Variable   | Description |
|------------|-------------|
| `GOROOT`   | The location where the Go SDK is installed. It contains the standard library, compiler, and core tools (`go`, `compile`, `link`, etc.). Usually set automatically during Go installation. |
| `GOPATH`   | The root of your workspace where Go code (outside the standard library) resides. It contains `src/`, `pkg/`, and `bin/`. Prior to Go modules, this was essential for project structure. |
| `GOBIN`    | The directory where Go installs compiled binaries when you run `go install`. If not set, defaults to `$GOPATH/bin`. |
| `GOOS`     | Target **operating system** for cross-compilation. For example, `linux`, `windows`, `darwin`. |
| `GOARCH`   | Target **architecture** for cross-compilation. For example, `amd64`, `arm64`, `386`. |
| `GOPROXY`  | URL of the Go module proxy used to fetch dependencies. Default is `https://proxy.golang.org`, but can be set to `"direct"` to disable or `"off"` to completely disable module downloading. |

- cross compilation

- list of valid os/arch options
```sh 
go tool dist list 
```

```sh
GOOS=windows GOARCH=amd64 go build -o win-d
emo.exe .   
GOOS=linux GOARCH=amd64 go build -o linux-demo
```
- For micro controllers

tinyGo is the project, it uses LLVM. Tyny go is good for micro controllers and also wasm.

## Keywords (25 of 25)

break,case,chan,const,continue,default,defer,else,fallthrough,for,func,go,goto,if,import,interface,map,package,range,return,select,struct,switch,type,var


## Builtin functions (18 of 18)

append,cap,complex,clear,close,copy,delete,imag,len,make,max,min,new,panic,print,println,real,recover


## Git repo
git@github.com:JitenPalaparthi/volkswagen-golang-batch3.git


## Package management

```sh
go get github.com/rs/zerolog
go mod tidy
```

- to create a vendor directory (to keep the source code of others). It creates a vendor directory and keeps the whole source code

```sh
go mod vendor
```

## Youtube channel for Golang based stuff..

https://www.youtube.com/playlist?list=PLJE7PIP1qj_Rn9vq4V4jGJbj5KqEIWSUc

## presentation

https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing

